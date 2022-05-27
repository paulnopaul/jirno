package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addTaskFilterHandler(taskRoot *cobra.Command, handler *taskHandler) {
	cmd := &cobra.Command{
		Use:   "filter",
		Short: "Get tasks by filter",
		Run:   handler.Filter,
	}
	cmd.Flags().StringP("pid", "i", "", "project id")
	cmd.Flags().Int64P("uid", "u", 0, "user id")
	cmd.Flags().StringP("datestart", "s", "", "date from")
	cmd.Flags().StringP("dateend", "e", "", "date to")

	taskRoot.AddCommand(cmd)
}

func (h taskHandler) Filter(cmd *cobra.Command, args []string) {
	filter, err := parseFilter(cmd, args)
	if err != nil {
		fmt.Printf("Task filter parsing error %v\n", err)
		return
	}
	res, err := h.tUsecase.GetByFilter(*filter)
	if err != nil {
		fmt.Printf("Task get by filter error: %v", err)
		return
	}

	err = h.storage.SetTaskList(res)
	if err != nil {
		fmt.Printf("Task updating local data error: %v", err)
		return
	}

	for i, value := range res {
		fmt.Println(taskToStr(value, i))
	}
}
