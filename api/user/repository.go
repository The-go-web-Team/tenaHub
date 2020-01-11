package user

import "github.com/TenaHub/api/entity"

type UserRepository interface {
	Users()([]entity.User, []error)
	User(user *entity.User)(*entity.User, []error)
	UserByID(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User)(*entity.User, []error)
	DeleteUser(id uint)(*entity.User, []error)
	StoreUser(user *entity.User)(*entity.User, []error)
	User(id uint) (*entity.User, []error)
}