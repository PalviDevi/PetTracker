package routes

import (
	ctrlPetTracker "PetTracker/app/http/controller"
	"github.com/gin-gonic/gin"
)

// Pet routes

func PetsTrackerRoutes(rgMain *gin.RouterGroup) {
	rgTracker := rgMain.Group("/v1/pet/events")
	{
		// All routes for Basepath => /api/v1/pet/events
		rgTracker.GET("/list", ctrlPetTracker.GetPetList)
		rgTracker.GET("/get", ctrlPetTracker.GetPetDataById)
		rgTracker.POST("/add", ctrlPetTracker.AddPetData)
		rgTracker.PUT("/update", ctrlPetTracker.UpdatePetData)
		rgTracker.DELETE("/delete", ctrlPetTracker.DeletePetData)
		rgTracker.GET("/add/event", ctrlPetTracker.AddEventData)
	}
}
