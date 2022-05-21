package task

import (
	"github.com/spf13/cobra"
	"jirno/internal/pkg/domain"
	"jirno/internal/pkg/localstorage"
)

type taskHandler struct {
	tUsecase domain.ITaskUsecase
	storage localstorage.LocalStorage
}

func NewTaskHandler(root *cobra.Command, taskUsecase domain.ITaskUsecase, localStorage localstorage.LocalStorage) {
	handler := taskHandler{
		tUsecase: taskUsecase,
		storage: localStorage,
	}

	taskRoot := &cobra.Command{Use: "task"}

	addTaskCreateHandler(taskRoot, &handler)

	addTaskUpdateHandler(taskRoot, &handler)

	addTaskFilterHandler(taskRoot, &handler)

	addTaskDeleteHandler(taskRoot, &handler)

	addTaskCompleteHandler(taskRoot, &handler)

	root.AddCommand(taskRoot)
}