package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addTaskCreateHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmd := &cobra.Command{
		Use:   "create TITLE",
		Short: "Create new task",
		Run:   handler.Create,
	}
	cmd.Flags().StringP("pid", "p", "", "project id")
	cmd.Flags().Int64P("uid", "u", 0, "user id")
	cmd.Flags().StringP("title", "n", "", "project id")
	cmd.Flags().StringP("description", "d", "", "task description")
	cmd.Flags().StringP("dateto", "t", "", "date task has to be completed")
	cmd.Flags().BoolP("completed", "c", false, "is task completed")
	taskRoot.AddCommand(cmd)

}

func (h taskHandler) Create(cmd *cobra.Command, args []string) {
	dTask, err := h.parseTask(cmd, args, typeTask)
	if err != nil {
		fmt.Printf("Task parsing error: %v\n", err)
		return
	}

	newTask, err := dTask.ToDomain()
	if err != nil {
		fmt.Printf("Task casting error: %v\n", err)
		return
	}

	id, err := h.tUsecase.Create(*newTask)
	if err != nil {
		fmt.Printf("Task creating error: %v\n", err)
		return
	}
	fmt.Printf("Created new task '%s' with id %v'\n", newTask.Title, id)
}
