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

	web_origin := os.Getenv("WEB_ORIGIN")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{web_origin}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	gob.Register(store.User{})

	web_origin_url, err := url.Parse(web_origin)
	if err != nil {
		log.Fatalf("web origin url could not be parsed: %v", err)
	}

	store := cookie.NewStore([]byte("secret"))
	if gin.Mode() == gin.ReleaseMode {
		store.Options(sessions.Options{Domain: web_origin_url.Hostname(), Path: "/", Secure: true})
	}
	router.Use(sessions.Sessions("auth-session", store))

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
	}

	notificationsRouter := router.Group("/notifications/settings", IsAuthenticated)
	{
		notificationsRouter.GET("/", GetAllNotificationSettings)
		notificationsRouter.POST("/create", CreateNotificationSetting)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "route not found"})
	})

	return router
}
