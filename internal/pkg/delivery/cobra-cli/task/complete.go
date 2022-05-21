package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

func addTaskCompleteHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmd := &cobra.Command{
		Use:   "complete",
		Short: "Complete Task",
		Run:   handler.Complete,
	}
	cmd.Flags().StringP("uuid", "i", "", "task id (full)")
	taskRoot.AddCommand(cmd)
}

func (h taskHandler) Complete(cmd *cobra.Command, args []string) {
	id, err := h.getTaskID(cmd, args)
	if err != nil {
		fmt.Printf("Can't get task id: %v\n", err)
		return
	}
	err = h.tUsecase.Complete(id)
	if err != nil {
		fmt.Printf("Task completion error: %v\n", err)
		return
	}
	fmt.Printf("Task %v successfully completed!\n", id)
}
