package task

import (
	"fmt"

	"github.com/google/uuid"
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
	if len(args) == 0 {
		fmt.Printf("id not provided")
		return
	}
	id := args[0]
	parsedId, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("id parsing error: %v\n", err)
		return
	}
	err = h.tUsecase.Delete(parsedId)
	if err != nil {
		fmt.Printf("Task deletion error: %v\n", err)
	}
}
