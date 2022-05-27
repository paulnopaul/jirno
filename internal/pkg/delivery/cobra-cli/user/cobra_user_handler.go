package user

import (
	"github.com/spf13/cobra"
	domain "jirno/internal/pkg/domain/user"
	"jirno/internal/pkg/localstorage"
)

type userHandler struct {
	uUsecase domain.IUserUsecase
	storage  localstorage.LocalStorage
}

func NewUserHandler(root *cobra.Command, userUsecase domain.IUserUsecase, localStorage localstorage.LocalStorage) {
	handler := userHandler{
		uUsecase: userUsecase,
		storage:  localStorage,
	}

	userRoot := &cobra.Command{Use: "user"}

	addUserSignUpHandler(userRoot, &handler)

	addUserLoginHandler(userRoot, &handler)

	addUserDeleteHandler(userRoot, &handler)

	addUserUpdateHandler(userRoot, &handler)

	root.AddCommand(userRoot)
}
