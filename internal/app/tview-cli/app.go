package tview_cli

import (
	"database/sql"
	"fmt"
	"jirno/internal/pkg/delivery/tview-cli/task_list"
	"jirno/internal/pkg/localstorage"
	"jirno/internal/pkg/repository/task/sqlite_repo"
	"jirno/internal/pkg/usecase/task"

	"github.com/rivo/tview"

	_ "github.com/mattn/go-sqlite3"
)

func RunApp() error {
	db, err := sql.Open("sqlite3", "jirno.db")
	if err != nil {
		return fmt.Errorf("db connection error: %v", err)
	}

	taskRepo := sqlite_repo.NewSQLiteTaskRepo(db)
	localStorage := localstorage.NewSQLiteLocalStorage(db)

	taskUsecase := task.NewTaskUsecase(taskRepo)

	app := tview.NewApplication()

	taskListViewHandler := task_list.NewTaskListView(taskUsecase, localStorage)
	taskListView := taskListViewHandler.SetView()

	app.SetRoot(taskListView, true)

	return app.Run()
}
