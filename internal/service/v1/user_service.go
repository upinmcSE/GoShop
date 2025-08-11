package v1service

import "github.com/upinmcSE/goshop/internal/repository"

type userService struct {
	repo repository.UserRepository
}

// CreateUser implements UserService.
func (u *userService) CreateUser() {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(uuid string) error {
	panic("unimplemented")
}

// GetAllUsers implements UserService.
func (u *userService) GetAllUsers(search string, page int, limit int) {
	panic("unimplemented")
}

// GetUserByUUID implements UserService.
func (u *userService) GetUserByUUID(uuid string) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(uuid string) {
	panic("unimplemented")
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
