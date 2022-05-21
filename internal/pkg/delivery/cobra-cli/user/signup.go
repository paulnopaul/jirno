package user

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addUserSignUpHandler(userRoot *cobra.Command, handler *userHandler) {

	cmdSignUp := &cobra.Command{
		Use:   "signup",
		Short: "Create new user",
		Run:   handler.SignUp,
	}
	cmdSignUp.Flags().StringP("name", "n", "", "user name")
	cmdSignUp.Flags().StringP("nickname", "s", "", "user nickname")
	cmdSignUp.Flags().StringP("email", "e", "", "user email")
	cmdSignUp.Flags().StringP("password", "p", "", "user password")
	userRoot.AddCommand(cmdSignUp)
}

func (h userHandler) SignUp(cmd *cobra.Command, args []string) {
	parsedUser, err := parseUser(cmd, args)
	if err != nil {
		fmt.Printf("User parsing error: %v\n", err)
		return
	}
	newID, err := h.uUsecase.Signup(*parsedUser)
	if err != nil {
		fmt.Printf("User signup failed: %v\n", err)
		return
	}
	fmt.Printf("Created user with id %v\n", newID)
}
