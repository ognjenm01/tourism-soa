package service

import (
	"testing"
	"tour/model"
	"tour/repo"
)

func initKeypointService() KeypointService {
	return KeypointService{KeypointRepository: &repo.KeypointRepository{DatabaseConnection: initDB()}}
}

// Testing insert for keypoint
func TestCreateKeypoint(t *testing.T) {
	keypointService := initKeypointService()
	keypoint := model.Keypoint{ID: 0, TourID: 3, Name: "TestKeypoint", Latitude: -42.0, Longitude: -42.0, Description: "TestKeypointDesc", Position: 0, Secret: "TestSecret", Image: "test.png"}

	error := keypointService.CreateKeypoint(&keypoint)

	if error != nil {
		t.Fatal("Keypoint create failed")
	}
}

// Testing delete for keypoint, valid id
func TestDeleteKeypoint(t *testing.T) {
	keypointService := initKeypointService()

	error := keypointService.DeleteKeypoint("1")

	if error != nil {
		t.Fatal("Keypoint delete failed")
	}
}

// Testing delete for keypoint, invalid id
func TestDeleteKeypointInvalid(t *testing.T) {
	keypointService := initKeypointService()

	error := keypointService.DeleteKeypoint("-100")

	if error == nil {
		t.Fatal("Keypoint delete with invalid id failed")
	}
}

// Testing get all keypoints
func TestGetAllKeypoints(t *testing.T) {
	keypointService := initKeypointService()

	keypoints, error := keypointService.GetAll()

	if error != nil {
		t.Fatal("Get all keypoints failed")
	} else if len(*keypoints) != 3 {
		t.Fatal("Invalid keypoint count")
	}
}

// Testing get all keypoints by tour id
func TestGetAllKeypointsByTourId(t *testing.T) {
	keypointService := initKeypointService()

	keypoints, error := keypointService.GetKeypointsByTourId("2")

	if error != nil {
		t.Fatal("Get all keypoints by tour failed")
	} else if len(*keypoints) != 2 {
		t.Fatal("Invalid keypoint count")
	}
}

// Testing get all keypoints by nonexistent tour id
func TestGetAllKeypointsByInvalidTourId(t *testing.T) {
	keypointService := initKeypointService()

	keypoints, error := keypointService.GetKeypointsByTourId("-100")

	if error != nil {
		t.Fatal("Get all keypoints by tour failed")
	} else if len(*keypoints) != 0 {
		t.Fatal("Invalid keypoint count")
	}
}
