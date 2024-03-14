package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"tls-watch/api/authenticator"
)

func NewRouter(auth *authenticator.OIDCAuthenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	authRouter := router.Group("/auth")
	{
		authRouter.GET("/login", authenticator.Login(auth))
		authRouter.GET("/callback", authenticator.LoginCallback(auth))
		authRouter.GET("/me", authenticator.IsAuthenticated, authenticator.Me)
		authRouter.GET("/logout", authenticator.Logout)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "route not found"})
	})

	return router
}
