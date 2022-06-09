package task_list

import (
	"jirno/internal/pkg/domain"
	"jirno/internal/pkg/localstorage"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type taskListHandler struct {
	tUsecase domain.ITaskUsecase
	storage  localstorage.LocalStorage
	tasks    []domain.Task
}

func NewTaskListView(taskUsecase domain.ITaskUsecase, localStorage localstorage.LocalStorage) taskListHandler {
	handler := taskListHandler{
		tUsecase: taskUsecase,
		storage:  localStorage,
	}
	handler.getTasks()
	return handler
}

func (h *taskListHandler) getTasks() {
	tasks, err := h.tUsecase.GetByFilter(domain.SmartTaskFilter{})
	if err != nil {
		return
	}
	h.tasks = tasks
}

func setVimArrows(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'j' {
		return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
	} else if event.Rune() == 'k' {
		return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
	}
	return event
}

func (h *taskListHandler) SetView() tview.Primitive {
	list := tview.NewList()
	list.ShowSecondaryText(false)
	list.SetBorder(true)

	list.SetTitle("Tasks")
	list.SetInputCapture(setVimArrows)
	for i, task := range h.tasks {
		list.AddItem(taskToString(task), "", 0, func() {
			h.tUsecase.Complete(task.ID)
			h.tasks[i].IsCompleted = true
		})
	}
	return list
}
