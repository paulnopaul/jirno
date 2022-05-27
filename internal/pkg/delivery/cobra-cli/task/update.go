package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addTaskUpdateHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update Task",
		Run:   handler.Update,
	}
	cmd.Flags().StringP("id", "i", "", "task id")
	cmd.Flags().Int64P("uid", "u", 0, "user id")
	cmd.Flags().StringP("pid", "p", "", "project id")
	cmd.Flags().StringP("title", "n", "", "project id")
	cmd.Flags().StringP("description", "d", "", "task description")
	cmd.Flags().StringP("datecompleted", "x", "", "date when task was completed")
	cmd.Flags().StringP("dateto", "t", "", "date task has to be completed")
	cmd.Flags().BoolP("completed", "c", false, "is task completed")
	taskRoot.AddCommand(cmd)
}

func (h taskHandler) Update(cmd *cobra.Command, args []string) {
	dTask, err := h.parseTask(cmd, args, typeTaskUpdate)
	if err != nil {
		fmt.Printf("Task parsing error: %v\n", err)
		return
	}
	update, err := (*dTask).ToUpdate()
	if err != nil {
		fmt.Printf("Task casting error: %v\n", err)
		return
	}
	err = h.tUsecase.Update(*update)
	if err != nil {
		fmt.Printf("Task updating error: %v\n", err)
	}
}
