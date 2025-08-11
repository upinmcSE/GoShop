package repository

type SqlUserRepository struct {
}

// Create implements UserRepository.
func (i *SqlUserRepository) Create() {
	panic("unimplemented")
}

// Delete implements UserRepository.
func (i *SqlUserRepository) Delete() error {
	panic("unimplemented")
}

// FindAll implements UserRepository.
func (i *SqlUserRepository) FindAll() {
	panic("unimplemented")
}

// FindByEmail implements UserRepository.
func (i *SqlUserRepository) FindByEmail() {
	panic("unimplemented")
}

// FindByUUID implements UserRepository.
func (i *SqlUserRepository) FindByUUID() {
	panic("unimplemented")
}

// Update implements UserRepository.
func (i *SqlUserRepository) Update() error {
	panic("unimplemented")
}

func NewSqlUserRepository() UserRepository {
	return &SqlUserRepository{}
}
