package user

import (
	"fmt"
	"jirno/internal/pkg/utils"
)

type DeliveryUser struct {
	Name     string
	Nickname string
	Email    string
	Password string
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
