package user

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"jirno/internal/pkg/domain"
	"jirno/internal/pkg/repository/user/mock"
	"testing"
)

type signupTest struct {
	deliveryUser  domain.DeliveryUser
	expectedError error
	id            int64
	user          domain.User
}

var (
	resUser = domain.DeliveryUser{
		Name:     "name",
		Nickname: "nickname",
		Password: "password",
	}
	signupData = []signupTest{
		{
			deliveryUser: resUser,
		},
	}
)

func TestUserUsecase_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	taskRepo := mock.NewMockIUserRepo(ctrl)
	usecase := NewUserUsecase(taskRepo)
	for index, testData := range signupData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().GetMaxUserID().Return(testData.id-1, nil).Times(1)
			taskRepo.EXPECT().Create(gomock.Any()).Return(nil).Times(1)
			res, err := usecase.Signup(testData.deliveryUser)
			assert.Nil(t, err)
			assert.Equal(t, res, testData.id)
		})
	}
}

type updateTest struct {
	deliveryUserUpdate domain.DeliveryUserUpdate
	user               domain.User
	userToUpdate       domain.User
	expectedError      error
}

var (
	updateDeliveryUser = domain.DeliveryUserUpdate{
		NewNickname: "new-nickname",
		DeliveryUser: domain.DeliveryUser{
			Nickname: "nickname",
			Email:    "email",
		},
	}
	updateUser = domain.User{
		ID:       2,
		Nickname: "nickname",
		Email:    "email",
	}
	updateUserToUpdate = domain.User{
		ID:       2,
		Nickname: "new-nickname",
		Email:    "email",
	}
	updateData = []updateTest{
		{
			expectedError:      nil,
			deliveryUserUpdate: updateDeliveryUser,
			user:               updateUser,
			userToUpdate:       updateUserToUpdate,
		},
	}
)

func TestUserUsecase_Update(t *testing.T) {
	taskRepo := mock.NewMockIUserRepo(gomock.NewController(t))
	usecase := NewUserUsecase(taskRepo)
	for index, testData := range updateData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().GetByNickname(testData.deliveryUserUpdate.Nickname).Return(&testData.user, nil).Times(1)
			taskRepo.EXPECT().Update(testData.userToUpdate).Return(nil).Times(1)
			err := usecase.Update(testData.deliveryUserUpdate)
			assert.Nil(t, err)
		})
	}
}

type deleteTest struct {
	expectedErr error
	nickname    string
	user        domain.User
}

var (
	deleteData = []deleteTest{
		{
			expectedErr: nil,
			nickname:    "nickname",
			user:        domain.User{ID: 2, Nickname: "nickname"},
		},
	}
)

func TestUserUsecase_Delete(t *testing.T) {
	taskRepo := mock.NewMockIUserRepo(gomock.NewController(t))
	usecase := NewUserUsecase(taskRepo)
	for index, testData := range deleteData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().GetByNickname(testData.nickname).Return(&testData.user, nil).Times(1)
			taskRepo.EXPECT().Delete(testData.user.ID).Return(nil).Times(1)
			err := usecase.Delete(testData.nickname)
			assert.Nil(t, err)
		})
	}
}
