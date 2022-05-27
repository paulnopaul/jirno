package project

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addProjectCreateHandler(projectRoot *cobra.Command, handler *projectHandler) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new project",
		Run:   handler.Create,
	}
	cmd.Flags().Int64SliceP("uids", "u", []int64{}, "user ids")
	cmd.Flags().StringP("ppid", "p", "", "parent project id")
	cmd.Flags().StringP("description", "d", "", "task description")
	cmd.Flags().StringP("datecompleted", "x", "", "date when task was completed")
	cmd.Flags().StringP("dateto", "t", "", "date task has to be completed")
	cmd.Flags().BoolSliceP("completed", "c", []bool{}, "is task completed")
	projectRoot.AddCommand(cmd)
}

func (h projectHandler) Create(cmd *cobra.Command, args []string) {
	dProject, err := parseProject(cmd, args, typeProject)
	if err != nil {
		fmt.Printf("Project parsing error: %v\n", err)
		return
	}

	newProject, err := dProject.ToDomain()
	if err != nil {
		fmt.Printf("Project casting error: %v\n", err)
		return
	}

	createdID, err := h.pUsecase.Create(*newProject)
	if err != nil {
		fmt.Printf("Project creating error: %v\n", err)
		return
	}
	fmt.Printf("Created new project '%s' with id %v\n", newProject.Title, createdID)
}
