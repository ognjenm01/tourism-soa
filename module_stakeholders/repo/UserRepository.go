package repo

import (
	"stakeholder/model"
	"strconv"

	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepository) GetUserById(id string) (model.User, error) {
	user := model.User{}
	pk, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return user, err
	}

	/*result := repo.DatabaseConnection.Preload("Tags").
	Preload("Keypoints").
	First(&user, "id = ?", pk)*/
	result := repo.DatabaseConnection.First(&user, "id = ?", pk)

	if result != nil {
		return user, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetAllUsers() ([]model.User, error) {
	users := []model.User{}

	result := repo.DatabaseConnection.Find(&users)

	if result != nil {
		return users, result.Error
	}
	return users, nil
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	result := repo.DatabaseConnection.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	result := repo.DatabaseConnection.First(&user, "username = ?", username)
	if result.Error != nil {
		return user, nil
	}
	return user, result.Error
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
