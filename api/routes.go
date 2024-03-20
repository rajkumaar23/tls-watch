package api

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(auth *OIDCAuthenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	authRouter := router.Group("/auth")
	{
		authRouter.GET("/login", Login(auth))
		authRouter.GET("/callback", LoginCallback(auth))
		authRouter.GET("/me", IsAuthenticated, Me)
		authRouter.GET("/logout", Logout)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "route not found"})
	})

	return router
}
