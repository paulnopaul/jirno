package user

import (
	"fmt"
	"jirno/internal/pkg/domain/user"
	"jirno/internal/pkg/utils"
)

type userUsecase struct {
	repo user.IUserRepo
}

func NewUserUsecase(userRepo user.IUserRepo) user.IUserUsecase {
	return &userUsecase{
		repo: userRepo,
	}
}

func (u userUsecase) GetByID(id int64) (*user.User, error) {
	return u.repo.GetByID(id)
}

func (u userUsecase) GetByNickname(nickname string) (*user.User, error) {
	return u.repo.GetByNickname(nickname)
}

func (u userUsecase) Signup(user user.DeliveryUser) (int64, error) {
	userToSignup, err := user.ToDomain()
	if err != nil {
		return 0, fmt.Errorf("user signup (to domain) failed %v", err)
	}
	maxID, err := u.repo.GetMaxUserID()
	if err != nil {
		return 0, fmt.Errorf("user signup (get id) failed %v", err)
	}
	userToSignup.ID = maxID + 1
	err = u.repo.Create(*userToSignup)
	if err != nil {
		return 0, fmt.Errorf("user signup failed %v", err)
	}
	return maxID + 1, nil
}

func (u userUsecase) Update(user user.DeliveryUserUpdate) error {
	userByID, err := u.repo.GetByNickname(user.Nickname)
	if err != nil {
		return fmt.Errorf("user update (get id) failed %v", err)
	}
	userToUpdate, err := user.ToDomain()
	if err != nil {
		return fmt.Errorf("user update (to domain) failed %v", err)
	}
	userToUpdate.ID = userByID.ID
	userToUpdate.Nickname = user.NewNickname

	return u.repo.Update(*userToUpdate)
}

func (u userUsecase) Delete(nickname string) error {
	user, err := u.repo.GetByNickname(nickname)
	if err != nil {
		return fmt.Errorf("user delete (get id) failed: %v", err)
	}
	return u.repo.Delete(user.ID)
}

func (u userUsecase) Check(user user.DeliveryUser) (bool, int64, error) {
	resUser, err := u.GetByNickname(user.Nickname)
	if err != nil {
		return false, 0, fmt.Errorf("user check (get user) failed %v", err)
	}
	fmt.Println(resUser.Name, resUser.ID)

	if !utils.CmpPassword(user.Password, resUser.Password) {
		return false, 0, nil
	}

	return true, resUser.ID, nil
}
