package user

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addUserDeleteHandler(userRoot *cobra.Command, handler *userHandler) {
	cmdSignUp := &cobra.Command{
		Use:   "delete",
		Short: "Delete user",
		Long:  "creates new user with login <nickname> and password <password>",
		Args:  cobra.ExactArgs(2),
		Run:   handler.Delete,
	}
	userRoot.AddCommand(cmdSignUp)
}

func (h userHandler) Delete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Printf("User nickname not provided\n")
		return
	}
	nickToDelete := args[0]
	err := h.uUsecase.Delete(nickToDelete)
	if err != nil {
		fmt.Printf("User delete failed: %v\n", err)
	}
	fmt.Printf("User %v deleted successfully\n", nickToDelete)
}
