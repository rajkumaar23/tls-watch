package api

import (
	"encoding/gob"
	"log"
	"net/url"
	"os"
	"tls-watch/api/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(auth *OIDCAuthenticator) *gin.Engine {
	router := gin.Default()

	webOrigin := os.Getenv("WEB_ORIGIN")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{webOrigin}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	gob.Register(store.User{})

	webOriginURL, err := url.Parse(webOrigin)
	if err != nil {
		log.Fatalf("web origin url could not be parsed: %v", err)
	}

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: webOriginURL.Hostname(), Path: "/", Secure: webOriginURL.Scheme == "https"})
	router.Use(sessions.SessionsMany([]string{"auth-session", "user-session"}, store))

	authRouter := router.Group("/auth")
	{
		authRouter.GET("/login", Login(auth))
		authRouter.GET("/callback", LoginCallback(auth))
		authRouter.GET("/me", IsAuthenticated, Me)
		authRouter.GET("/logout", Logout)
	}

	domainsRouter := router.Group("/domains", IsAuthenticated)
	{
		domainsRouter.GET("/", GetAllDomains)
		domainsRouter.POST("/create", CreateDomain)
		domainsRouter.DELETE("/delete", DeleteDomain)
	}

	notificationsRouter := router.Group("/notifications/settings", IsAuthenticated)
	{
		notificationsRouter.GET("/", GetAllNotificationSettings)
		notificationsRouter.POST("/create", CreateOrUpdateNotificationSetting)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "route not found"})
	})

	return router
}
