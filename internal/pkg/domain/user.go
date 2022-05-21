package domain

import (
	"fmt"
	"jirno/internal/pkg/utils"
)

type Users []int64

type DeliveryUser struct {
	Name     string
	Nickname string
	Email    string
	Password string
}

type DeliveryUserUpdate struct {
	DeliveryUser
	NewNickname string
}

func (u DeliveryUser) ToDomain() (*User, error) {
	res := &User{}
	res.Name = u.Name
	res.Nickname = u.Nickname
	res.Email = u.Email
	if u.Password != "" {
		pwd, err := utils.HashPassword(u.Password)
		if err != nil {
			return nil, fmt.Errorf("user to delivery casting failed: %v", err)
		}
		res.Password = pwd
	}
	return res, nil
}

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
