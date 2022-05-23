package user

import (
	"fmt"
	"github.com/spf13/cobra"
	domain "jirno/internal/pkg/domain/user"
)

func addUserLoginHandler(userRoot *cobra.Command, handler *userHandler) {
	cmdSignUp := &cobra.Command{
		Use:   "login",
		Short: "Log in",
		Run:   handler.Login,
	}
	cmdSignUp.Flags().StringP("nickname", "n", "", "user nickname")
	cmdSignUp.Flags().StringP("password", "p", "", "user password")
	userRoot.AddCommand(cmdSignUp)
}

func (h userHandler) Login(cmd *cobra.Command, args []string) {
	parsedUser, err := parseUser(cmd, args)
	if err != nil {
		fmt.Printf("user signup (parse user) failed: %v", err)
		return
	}
	isCorrect, id, err := h.uUsecase.Check(*parsedUser)
	if err != nil {
		fmt.Printf("user login failed: %v", err)
		return
	}
	if !isCorrect {
		fmt.Printf("Failed login: wrong credentials")
		return
	}
	err = h.storage.SetCurrentUser(domain.User{ID: id})
	if err != nil {
		fmt.Printf("User updating local data error: %v", err)
		return
	}
	fmt.Printf("Successfull login with id %v", id)
}
