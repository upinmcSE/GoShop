package repository

type UserRepository interface {
	FindAll() ()
	Create()
	FindByUUID() ()
	Update() error
	Delete() error
	FindByEmail() ()
}
