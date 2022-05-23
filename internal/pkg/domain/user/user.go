package user

type User struct {
	ID       int64
	Name     string
	Nickname string
	Email    string
	Password []byte
}

type IUserUsecase interface {
	GetByID(id int64) (*User, error)
	GetByNickname(nickname string) (*User, error)
	Check(user DeliveryUser) (bool, int64, error)
	Signup(user DeliveryUser) (int64, error)
	Update(user DeliveryUserUpdate) error
	Delete(nickname string) error
}

//go:generate mockgen -destination=../repository/user/mock/mock_repo.go -package=mock jirno/internal/pkg/domain IUserRepo
type IUserRepo interface {
	GetByID(id int64) (*User, error)
	GetByNickname(nickname string) (*User, error)
	Create(user User) error
	Update(user User) error
	Delete(id int64) error
	GetMaxUserID() (int64, error)
}
