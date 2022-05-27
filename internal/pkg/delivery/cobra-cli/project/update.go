package project

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addProjectUpdateHandler(projectRoot *cobra.Command, handler *projectHandler) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update project",
		Run:   handler.Update,
	}
	cmd.Flags().StringP("id", "i", "", "project id")
	cmd.Flags().Int64SliceP("uids", "u", []int64{}, "user ids")
	cmd.Flags().StringP("ppid", "p", "", "parent project id")
	cmd.Flags().StringP("description", "d", "", "project description")
	cmd.Flags().StringP("datecompleted", "x", "", "date when project was completed")
	cmd.Flags().StringP("dateto", "t", "", "date project has to be completed")
	cmd.Flags().BoolSliceP("completed", "c", []bool{}, "is project completed")
	projectRoot.AddCommand(cmd)
}

func (h projectHandler) Update(cmd *cobra.Command, args []string) {
	dProject, err := parseProject(cmd, args, typeProjectUpdate)
	if err != nil {
		fmt.Printf("Project parsing error: %v\n", err)
		return
	}

	projectUpdate, err := dProject.ToUpdate()
	if err != nil {
		fmt.Printf("Project casting error: %v\n", err)
		return
	}

	err = h.pUsecase.Update(*projectUpdate)
	if err != nil {
		fmt.Printf("Project update error: %v\n", err)
		return
	}
	fmt.Printf("Updated project successfully\n")
}
