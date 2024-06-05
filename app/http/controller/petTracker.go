package ctrlPetTracker

import (
	bvrPetTracker "PetTracker/app/behaviour"
	"PetTracker/app/entity"
	"PetTracker/app/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPetList(c *gin.Context) {
	var reqData entity.Pet
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responses.SuccessApiResponse(c, "success.data_fetch", nil,
		bvrPetTracker.GetListOfPet(reqData.Species))
}

func GetPetDataById(c *gin.Context) {
	var petData entity.Pet
	if err := c.ShouldBindJSON(&petData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responses.SuccessApiResponse(c, "success.data_fetch", nil,
		bvrPetTracker.GetPetById(petData.ID))
}

func AddPetData(c *gin.Context) {
	var petData map[string]interface{}
	if err := c.ShouldBindJSON(&petData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responses.SuccessApiResponse(c, "success.data_created", nil,
		bvrPetTracker.AddPetData(petData))
}

func UpdatePetData(c *gin.Context) {
	var petData entity.Pet
	if err := c.ShouldBindJSON(&petData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responses.SuccessApiResponse(c, "success.data_updated", nil,
		bvrPetTracker.UpdatePetData(petData.ID))
}
func DeletePetData(c *gin.Context) {
	var petData entity.Pet
	if err := c.ShouldBindJSON(&petData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responses.SuccessApiResponse(c, "success.data_deleted", nil,
		bvrPetTracker.DeletePetData(petData.ID))
}
func AddEventData(c *gin.Context) {
	var petData map[string]interface{}
	if err := c.ShouldBindJSON(&petData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responses.SuccessApiResponse(c, "success.data_deleted", nil,
		bvrPetTracker.AddEventData(petData))
}
