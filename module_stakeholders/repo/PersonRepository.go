package repo

import (
	"stakeholder/model"
	"strconv"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *PersonRepository) GetPersonById(id string) (model.Person, error) {
	person := model.Person{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return person, err
	}

	/*result := repo.DatabaseConnection.Preload("Tags").
	Preload("Keypoints").
	First(&user, "id = ?", pk)*/
	result := repo.DatabaseConnection.First(&person, "id = ?", pk)

	if result != nil {
		return person, result.Error
	}
	return person, nil
}

func (repo *PersonRepository) GetAllPeople() ([]model.Person, error) {
	persons := []model.Person{}

	result := repo.DatabaseConnection.Find(&persons)

	if result != nil {
		return persons, result.Error
	}
	return persons, nil
}

func (repo *PersonRepository) CreatePerson(person *model.Person) error {
	result := repo.DatabaseConnection.Create(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Stvarno nemam snage sad
/*
func (repo *UserRepository) Update(id int, user *model.User) error {
	result := repo.DatabaseConnection.Model(user).
		Updates(map[string]interface{}{
			"Name":          user.Name,
			"Description":   user.Description,
			"Price":         user.Price,
			"Difficulty":    user.Difficulty,
			"TransportType": user.TransportType,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}*/
