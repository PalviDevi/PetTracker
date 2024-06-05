package repository

import (
	"PetTracker/app/globals/dbconnections"
	"log"
)

func AddPetTrackerData() map[string]interface{} {
	var dbRows map[string]interface{}
	sql := "INSERT INTO pets_data (name, owner, species, birth, death) VALUES (?, ?, ?, ?, ?)"
	queryValues := []interface{}{dbRows["name"], dbRows["owner"], dbRows["species"], dbRows["birth"], dbRows["death"]}

	if err := dbconnections.MainDB.Raw(sql, queryValues...).Scan(&dbRows).Error; err != nil {
		log.Fatal("repository::AddPetTrackerData - Failure during SQL execution", err, sql)
		return nil
	}
	return dbRows
}

func GetListOfAllPets(species string) map[string]interface{} {
	var dbRows map[string]interface{}
	sql := "SELECT * FROM pets_data WHERE species = ? OR ? IS NULL "
	if err := dbconnections.MainDB.Raw(sql, species).Scan(&dbRows).Error; err != nil {
		log.Fatal("repository::GetListOfAllPets - Failure during SQL execution", err, sql, species)
		return nil
	}
	return dbRows
}

func GetPetEventById(id int) map[string]interface{} {
	var dbRows map[string]interface{}
	sql := "SELECT  p.id AS pet_id," +
		" p.name AS pet_name," +
		" p.owner AS pet_owner, " +
		" p.species AS pet_species, " +
		" p.birth AS pet_birth, " +
		" p.death AS pet_death, " +
		" e.date AS event_date, " +
		" e.type AS event_type, " +
		" e.remark AS event_remark " +
		"FROM " +
		"pets_data p " +
		"LEFT JOIN " +
		"pet_events e ON e.id = p.pet_id " +
		"WHERE  p.id = ? " +
		"ORDER BY e.date DESC "
	if err := dbconnections.MainDB.Raw(sql, id).Scan(&dbRows).Error; err != nil {
		log.Fatal("repository::GetPetById - Failure during SQL execution", err, sql, id)
		return nil
	}
	return dbRows
}

func UpdatePetById(id int) bool {
	var dbRows map[string]interface{}
	sql := "UPDATE pets_data SET name = ?, owner = ?, species = ?, birth = ?, death = ? WHERE id = ?"

	params := []interface{}{
		dbRows["name"],
		dbRows["owner"],
		dbRows["species"],
		dbRows["birth"],
		id,
	}
	if err := dbconnections.MainDB.Exec(sql, params...).Error; err != nil {
		log.Fatal("repoCorrection::UpdatePetById - Failure during SQL execution", err, sql, params)
		return false
	}
	return true
}

func AddEventData() map[string]interface{} {
	var dbRows map[string]interface{}
	sql := "INSERT INTO events_data (pet_id, date, type, remark)  VALUES (?, ?, ?, ?)"
	queryValues := []interface{}{dbRows["pet_id"], dbRows["date"], dbRows["type"], dbRows["remark"]}

	if err := dbconnections.MainDB.Raw(sql, queryValues...).Scan(&dbRows).Error; err != nil {
		log.Fatal("repository::AddEventData - Failure during SQL execution", err, sql, queryValues)
		return nil
	}
	return dbRows
}
func DeletePetById(id int) map[string]interface{} {
	if id == 0 {
		return nil
	}
	var dbRows map[string]interface{}
	sql := "DELETE FROM pets_data WHERE id = ?"

	if err := dbconnections.MainDB.Raw(sql, id).Scan(&dbRows).Error; err != nil {
		log.Fatal("repository::AddEventData - Failure during SQL execution", err, sql, id)
		return nil
	}
	return dbRows
}
