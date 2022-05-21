package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addTaskCreateHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmdNew := &cobra.Command{
		Use:   "create",
		Short: "Create new task",
		Run:   handler.Create,
	}
	cmdNew.Flags().Int64P("uid", "u", 0, "user id")
	cmdNew.Flags().StringP("pid", "p", "", "project id")
	cmdNew.Flags().StringP("description", "d", "", "task description")
	cmdNew.Flags().StringP("datecompleted", "x", "", "date when task was completed")
	cmdNew.Flags().StringP("dateto", "t", "", "date task has to be completed")
	cmdNew.Flags().BoolSliceP("completed", "c", []bool{}, "is task completed")
	taskRoot.AddCommand(cmdNew)

}

func (h taskHandler) Create(cmd *cobra.Command, args []string) {
	dTask, err := parseTask(cmd, args, typeTask)
	if err != nil {
		fmt.Printf("Task parsing error: %v\n", err)
		return
	}

	newTask, err := (*dTask).ToDomain()
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
