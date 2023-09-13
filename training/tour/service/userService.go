package service

// create struct
type UserService struct {
	Nama string
}

// function provider
func NewUserService(nama string) *UserService {
	return &UserService{
		Nama: nama,
	}
}

// method
func (u *UserService) PrintName() string {
	return u.Nama
}
