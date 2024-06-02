package service

import (
	"log"
	"os"
	"stakeholder/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtGenerator struct {
	Audience string
	Issuer   string
	Key      string
}

type AccessTokenService struct {
	UserService   *UserService
	PersonService *PersonService
}

func NewJwtGenerator() *JwtGenerator {
	return &JwtGenerator{
		Audience: getEnv("JWT_AUDIENCE", "explorer-front.com"),
		Issuer:   getEnv("JWT_ISSUER", "explorer"),
		Key:      getEnv("JWT_KEY", "explorer_secret_key"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func (service *AccessTokenService) RegisterUser(user *model.Register) error {
	a, error := service.UserService.GetUserByUsername(user.Username)
	if a.Username == user.Username {
		log.Printf("[DB] - postoji korisnik vec")
		return error
	}

	newUser := model.User{
		ID:        0,
		Username:  user.Username,
		Password:  user.Password,
		Role:      1,
		IsActive:  true,
		IsBlocked: false,
		IsEnabled: true,
	}

	error = service.UserService.CreateUser(&newUser)
	if error != nil {
		return error
	}

	usr, err := service.UserService.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}

	newPerson := model.Person{
		ID:           0,
		UserId:       usr.ID,
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
		Biography:    user.Biography,
		Quote:        user.Quote,
	}

	error = service.PersonService.CreatePerson(&newPerson)
	if error != nil {
		return error
	}
	return nil
}

func (service *AccessTokenService) LoginUser(user *model.Login) (*model.AccessToken, error) {

	usr, err := service.UserService.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	if usr.Password != user.Password || user.Password == "" {
		return nil, nil
	}

	claims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       usr.ID,
		"username": usr.Username,
		"role":     usr.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	gen := NewJwtGenerator()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(gen.Key))
	if err != nil {
		return nil, err
	}

	return &model.AccessToken{
		ID:          usr.ID,
		AccessToken: tokenString,
	}, nil
}
