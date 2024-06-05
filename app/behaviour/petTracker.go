package bvrPetTracker

import (
	"PetTracker/app/repository"
	"log"
)

func GetListOfPet(species string) map[string]interface{} {
	return repository.GetListOfAllPets(species)
}
func GetPetById(id int) map[string]interface{} {
	if id == 0 {
		log.Fatalf("bvrPetTracker.GetPetById: Pet Id not found ")
	}
	return repository.GetPetEventById(id)
}

func AddPetData(data map[string]interface{}) map[string]interface{} {
	if data["pet"] == nil {
		log.Fatalf("bvrPetTracker.AddPet: Pet data not found ")
	}
	return repository.AddPetTrackerData()
}
func UpdatePetData(id int) bool {
	if id == 0 {
		log.Fatalf("bvrPetTracker.UpdatePetData: Pet data not found ")
	}
	return repository.UpdatePetById(id)
}
func AddEventData(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		log.Fatalf("bvrPetTracker.AddEventData: Event data not found ")
	}
	return repository.AddEventData()
}
func DeletePetData(id int) map[string]interface{} {
	if id == 0 {
		return nil
	}
	return repository.DeletePetById(id)
}
