package service

import (
	"testing"
	"tour/model"
	"tour/repo"
)

func initEquipmentService() EquipmentService {
	return EquipmentService{EquipmentRepository: &repo.EquipmentRepository{DatabaseConnection: initDB()}}
}

// Testing create equipment
func TestCreateEquipment(t *testing.T) {
	equipmentService := initEquipmentService()
	equipment := model.Equipment{Name: "TestEquipment", Description: "TestEquipmentDescription"}

	error := equipmentService.CreateEquipment(&equipment)

	if error != nil {
		t.Fatal("Failed create equipment")
	}
}

// Testing update equipment
func TestUpdateEquipment(t *testing.T) {
	equipmentService := initEquipmentService()
	equipment := model.Equipment{ID: 1, Name: "TestEquipment", Description: "TestEquipmentDescription"}

	error := equipmentService.UpdateEquipment(&equipment)
	if error != nil {
		t.Fatal("Failed update equipment")
	}

	eq, err := equipmentService.GetEquipmentById("1")
	if err != nil || eq.Name != "TestEquipment" || eq.Description != "TestEquipmentDescription" {
		t.Fatal("Failed update equipment")
	}
}

// Testing get all equipment
func TestGetAllEquipment(t *testing.T) {
	equipmentService := initEquipmentService()
	equipment, error := equipmentService.GetAllEquipment()
	if error != nil || len(*equipment) != 3 {
		t.Fatal("Failed get all equipment")
	}
}

// Testing get equipment by id
func TestGetEquipmentByID(t *testing.T) {
	equipmentService := initEquipmentService()
	equipment, error := equipmentService.GetEquipmentById("1")
	if error != nil || equipment.Name != "TestEq1" || equipment.Description != "TestEqDesc1" {
		t.Fatal("Failed get equipment by id")
	}
}

// Testing get equipment by tour id
func TestGetEquipmentByTourID(t *testing.T) {
	equipmentService := initEquipmentService()
	equipment, error := equipmentService.GetEquipmentByTourId("1")
	if error != nil || len(*equipment) != 2 {
		t.Fatal("Failed get equipment by tour id")
	}
}
