package services

import (
	"src/db"
	"src/models"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct{}

func (this *userService) Get(id int64) *models.User {
	return models.UserRepository.Get(db.GetDB(), id)
}
