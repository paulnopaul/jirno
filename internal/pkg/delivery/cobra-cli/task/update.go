package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addTaskUpdateHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmdUpdate := &cobra.Command{
		Use:   "update",
		Short: "Update Task",
		Run:   handler.Update,
	}
	cmdUpdate.Flags().StringP("id", "i", "", "task id")
	cmdUpdate.Flags().Int64P("uid", "u", 0, "user id")
	cmdUpdate.Flags().StringP("pid", "p", "", "project id")
	cmdUpdate.Flags().StringP("description", "d", "", "task description")
	cmdUpdate.Flags().StringP("datecompleted", "x", "", "date when task was completed")
	cmdUpdate.Flags().StringP("dateto", "t", "", "date task has to be completed")
	cmdUpdate.Flags().BoolSliceP("completed", "c", []bool{}, "is task completed")
	taskRoot.AddCommand(cmdUpdate)

}

func (h taskHandler) Update(cmd *cobra.Command, args []string) {
	dTask, err := parseTask(cmd, args, typeTaskUpdate)
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
