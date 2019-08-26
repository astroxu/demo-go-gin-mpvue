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

func (this *userService) GetUserByOpt(opt map[string]interface{}) *models.User {
	return models.UserRepository.GetUserByOpt(db.GetDB(), opt)
}

func (this *userService) Post(user *models.User) int64 {
	return models.UserRepository.Post(db.GetDB(), user)
}

func (this *userService) Get(id int64) *models.User {
	return models.UserRepository.Get(db.GetDB(), id)
}

func (this *userService) Put(user *models.User) bool {
	return models.UserRepository.Update(db.GetDB(), user)
}

func (this *userService) Delete(id int64) bool {
	return models.UserRepository.Delete(db.GetDB(), id)
}

func (this *userService) GetUsers() *[]models.User {
	return models.UserRepository.GetUsers(db.GetDB())
}

func (this *userService) GetUsersPaged(offset, perPage int64) *[]models.User {
	return models.UserRepository.GetUsersPaged(db.GetDB(), offset, perPage)
}
