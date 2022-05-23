package tview_cli

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
	"jirno/internal/pkg/domain"
	"jirno/internal/pkg/utils"
	"time"
)

func createTaskMap() map[uuid.UUID]*domain.Task {
	res := make(map[uuid.UUID]*domain.Task)
	for _, task := range exampleTasks {
		res[task.ID] = &task
	}
	return res
}

var (
	date1        = time.Now()
	date2        = time.Date(2022, 11, 5, 0, 0, 0, 0, date1.Location())
	exampleTasks = []domain.Task{
		{
			ID:          uuid.New(),
			Title:       "Clean room",
			DateTo:      &date1,
			IsCompleted: false,
		},
		{
			ID:          uuid.New(),
			Title:       "Meow",
			DateTo:      &date1,
			IsCompleted: true,
		},
		{
			ID:          uuid.New(),
			Title:       "Meow meow",
			DateTo:      &date1,
			IsCompleted: false,
		},
		{
			ID:          uuid.New(),
			Title:       "Call lawyer",
			DateTo:      &date2,
			IsCompleted: false,
		},
		{
			ID:          uuid.New(),
			Title:       "Meow meow",
			DateTo:      nil,
			IsCompleted: true,
		},
	}
	taskMap      = createTaskMap()
)

func taskToString(t domain.Task) string {
	completed := ' '
	if t.IsCompleted == true {
		completed = 'x'
	}
	dateStr := ""
	ts, te := utils.GetDayRange(time.Now())
	if t.DateTo != nil {
		if t.DateTo.Unix() > ts.Unix() && t.DateTo.Unix() < te.Unix() {
			dateStr = "today"
		} else {
			dateStr = t.DateTo.Format("2006-01-03")
		}
	}
	dateStr = fmt.Sprintf("(%v)", dateStr)
	return fmt.Sprintf("[ %c ] %v ", completed, t.Title) + dateStr
}

func setVimArrows(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'j' {
		return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
	} else if event.Rune() == 'k' {
		return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
	}
	return event
}

func SetTaskView(showCompleted bool) tview.Primitive {
	list := tview.NewList()
	list.ShowSecondaryText(false)
	list.SetBorder(true)
	if showCompleted {
		list.SetTitle("Completed tasks")
	} else {
		list.SetTitle("New tasks")
	}
	list.SetInputCapture(setVimArrows)

	for _, task := range exampleTasks {
		if task.IsCompleted == showCompleted {
			list.AddItem(taskToString(task), "", 0, nil)
		}
	}
	return list
}

func RunApp() error{
	app := tview.NewApplication()
	app.EnableMouse(true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(SetTaskView(false), 0, 1, true).
		AddItem(SetTaskView(true), 0, 1, false)

	app.SetRoot(flex, true)
	err := app.Run()
	if err != nil {
		return err
	}
	return nil
}
