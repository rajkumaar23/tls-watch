package api

import (
	"database/sql"
	"log"
	"net/http"
	store "tls-watch/api/store"

	"github.com/gin-gonic/gin"
)

func CreateOrUpdateNotificationSetting(ctx *gin.Context) {
	var newSetting store.NotificationSetting

	if err := ctx.ShouldBindJSON(&newSetting); err != nil {
		log.Printf("adding new notification setting failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "adding new notification setting failed"})
		return
	}

	newSetting.UserID = getUserProfile(ctx).ID
	setting, err := store.GetNotificationSetting(&newSetting)

	if setting != nil {
		err = store.UpdateNotificationSetting(&newSetting)
		if err != nil {
			log.Printf("updating notification setting failed : %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "updating notification setting failed"})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "notification setting updated successfully"})
		return
	}

	if err != sql.ErrNoRows {
		log.Printf("adding new notification setting failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "adding new notification setting failed"})
		return
	}

	err = store.CreateNotificationSetting(&newSetting)

	if err != nil {
		log.Printf("adding new notification setting failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "adding new notification setting failed"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "notification setting added successfully"})
}

func GetAllNotificationSettings(ctx *gin.Context) {
	user := getUserProfile(ctx)
	settings, err := store.GetAllNotificationSettingsByUserID(user.ID)
	if err != nil {
		log.Printf("fetching notification settings failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "fetching notification settings failed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"settings": settings})
}
