package userService

import "FinalProject/models"

type UserService interface {
	AddUser(user *models.AddUser) (*models.User, error)
	GetUser() (*[]models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) (int, error)
	DeleteUser(id int) (int, error)
}