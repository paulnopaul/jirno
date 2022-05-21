package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

func addProjectDeleteHandler(projectRoot *cobra.Command, handler *projectHandler) {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete Project",
		Run:   handler.Delete,
	}
	cmd.Flags().StringP("id", "i", "", "project id")
	projectRoot.AddCommand(cmd)
}

func (h projectHandler) Delete(cmd *cobra.Command, args []string) {
	id, err := parseProjectID(cmd, args)
	if err != nil {
		fmt.Printf("Id parsing failed: %v\n", err)
		return
	}
	err = h.pUsecase.Delete(id)
	if err != nil {
		fmt.Printf("Project deletion error: %v\n", err)
	}
	fmt.Printf("Project %v deleted successfully\n", id)
}
