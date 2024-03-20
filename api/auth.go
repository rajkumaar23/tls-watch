package api

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	store "tls-watch/api/store"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OIDCAuthenticator struct {
	*oidc.Provider
	oauth2.Config
}

func NewOIDCAuthenticator() (*OIDCAuthenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+os.Getenv("AUTH0_DOMAIN")+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &OIDCAuthenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

func (a *OIDCAuthenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func Login(auth *OIDCAuthenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := GenerateRandomState()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		session := sessions.Default(ctx)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

func LoginCallback(auth *OIDCAuthenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid state parameter"})
			return
		}

		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "failed to exchange an authorization code for a token"})
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to verify ID token"})
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		oidc_subject := profile["sub"].(string)
		user, err := store.GetUserByOIDCSubject(oidc_subject)
		if err == sql.ErrNoRows {
			err = store.CreateUser(&store.User{
				Name:        profile["name"].(string),
				Picture:     profile["picture"].(string),
				OIDCSubject: profile["sub"].(string),
			})
		}

		if err != nil {
			log.Printf("creating user failed : %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "creating user failed"})
			return
		}

		if user == nil {
			user, err = store.GetUserByOIDCSubject(oidc_subject)
			if err != nil {
				log.Printf("fetching new user failed : %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "fetching new user failed"})
				return
			}
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", user)
		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, "/auth/me")
	}
}

func getUserProfile(ctx *gin.Context) store.User {
	session := sessions.Default(ctx)
	return session.Get("profile").(store.User)
}

func Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"profile": getUserProfile(ctx)})
}

func IsAuthenticated(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	} else {
		ctx.Next()
	}
}

func Logout(ctx *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
