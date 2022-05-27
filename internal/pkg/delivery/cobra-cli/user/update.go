package user

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addUserUpdateHandler(userRoot *cobra.Command, handler *userHandler) {

	cmdSignUp := &cobra.Command{
		Use:   "update",
		Short: "Update user",
		Long:  "creates new user with login <nickname> and password <password>",
		Args:  cobra.ExactArgs(2),
		Run:   handler.SignUp,
	}
	cmdSignUp.Flags().StringP("name", "n", "", "user name")
	cmdSignUp.Flags().StringP("nickname", "s", "", "user nickname")
	cmdSignUp.Flags().StringP("newnickname", "u", "", "user nickname")
	cmdSignUp.Flags().StringP("email", "e", "", "user email")
	cmdSignUp.Flags().StringP("password", "p", "", "user password")
	userRoot.AddCommand(cmdSignUp)
}

func (h userHandler) Update(cmd *cobra.Command, args []string) {
	parsedUserUpdate, err := parseUserUpdate(cmd, args)
	if err != nil {
		fmt.Printf("User parsing error: %v\n", err)
		return
	}

	err = h.uUsecase.Update(*parsedUserUpdate)
	if err != nil {
		fmt.Printf("Project update failed: %v\n", err)
		return
	}
	fmt.Printf("User successfully updated\n")
}