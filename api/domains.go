package api

import (
	"database/sql"
	"log"
	"net/http"
	store "tls-watch/api/store"

	"github.com/gin-gonic/gin"
)

func CreateDomain(ctx *gin.Context) {
	var newDomain store.Domain

	if err := ctx.ShouldBindJSON(&newDomain); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "adding new domain failed"})
		return
	}

	//TODO: add DNS lookup validation
	
	user_id := getUserProfile(ctx).ID
	domain, err := store.GetDomainByUserAndValue(user_id, newDomain.Domain)
	if domain != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "domain already exists"})
		return
	}
	if err != sql.ErrNoRows {
		log.Printf("adding new domain failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "adding new domain failed"})
		return
	}

	newDomain.UserID = user_id
	err = store.CreateDomain(&newDomain)

	if err != nil {
		log.Printf("adding new domain failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "adding new domain failed"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "domain added successfully"})
}

func DeleteDomain(ctx *gin.Context) {
	var domainToDelete store.Domain
	if err := ctx.ShouldBindJSON(&domainToDelete); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "deleting domain failed"})
		return
	}
	
	user_id := getUserProfile(ctx).ID
	_, err := store.GetDomainByUserAndValue(user_id, domainToDelete.Domain)
	if err != nil {
		log.Printf("deleting domain failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "deleting domain failed"})
		return
	}

	err = store.DeleteDomainByValue(domainToDelete)
	if err != nil {
		log.Printf("deleting domain failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "deleting domain failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "domain deleted successfuly"})

}

func GetAllDomains(ctx *gin.Context) {
	user := getUserProfile(ctx)
	domains, err := store.GetAllDomainsByUserID(user.ID)
	if err != nil {
		log.Printf("fetching domains failed : %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "fetching domains failed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"domains": domains})
}
