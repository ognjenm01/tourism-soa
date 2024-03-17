package service

import (
	"testing"
	"tour/model"
	"tour/repo"
)

func initTourEquipmentService() TourEquipmentService {
	return TourEquipmentService{TourEquipmentRepository: &repo.TourEquipmentRepository{DatabaseConnection: initDB()}}
}

// Testing create tour equipment
func TestCreateTourEquipment(t *testing.T) {
	tequipmentService := initTourEquipmentService()
	tequipment := model.TourEquipment{TourID: 2, EquipmentID: 2}

	error := tequipmentService.CreateTourEquipment(&tequipment)

	if error != nil {
		t.Fatal("Failed create equipment")
	}
}

// Testing get all tour equipment
func TestGetAllTourEquipment(t *testing.T) {
	tequipmentService := initTourEquipmentService()
	tequipment, error := tequipmentService.GetAllTourEquipment()
	if error != nil || len(*tequipment) != 4 {
		t.Fatal("Failed get all tour equipment")
	}
}

// Testing delete tour equipment
func TestDeleteTourEquipment(t *testing.T) {
	tequipmentService := initTourEquipmentService()
	tequipment := model.TourEquipment{TourID: 2, EquipmentID: 1}

	error := tequipmentService.DeleteTourEquipment(&tequipment)

	if error != nil {
		t.Fatal("Failed delete equipment")
	}
}
