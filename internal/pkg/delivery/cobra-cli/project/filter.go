package project

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addProjectFilterHandler(projectRoot *cobra.Command, handler *projectHandler) {
	cmd := &cobra.Command{
		Use:   "filter",
		Short: "Get project by filter",
		Run:   handler.Filter,
	}
	cmd.Flags().StringP("ppid", "i", "", "parent project id")
	cmd.Flags().Int64P("uid", "u", 0, "user id")
	cmd.Flags().StringP("datestart", "s", "", "date from")
	cmd.Flags().StringP("dateend", "e", "", "date to")

	projectRoot.AddCommand(cmd)
}

func (h projectHandler) Filter(cmd *cobra.Command, args []string) {
	filter, err := parseFilter(cmd, args)
	if err != nil {
		fmt.Printf("Task filter parsing error %v\n", err)
		return
	}

	res, err := h.pUsecase.GetByFilter(*filter)
	if err != nil {
		fmt.Printf("Task get by filter error: %v", err)
		return
	}

	for i, value := range res {
		fmt.Println(i, value.ID, value.Title, value.Description, value.IsCompleted)
	}
}
