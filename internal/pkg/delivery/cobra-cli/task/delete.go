package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

func addTaskDeleteHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmdDelete := &cobra.Command{
		Use:   "delete",
		Short: "Delete Task",
		Run:   handler.Delete,
	}
	taskRoot.AddCommand(cmdDelete)
}

func (h taskHandler) Delete(cmd *cobra.Command, args []string) {
	id, err := h.getTaskID(cmd, args)
	if err != nil {
		fmt.Printf("Can't get task id: %v\n", err)
		return
	}
	err = h.tUsecase.Delete(id)
	if err != nil {
		fmt.Printf("Task deletion error: %v\n", err)
	}
}
