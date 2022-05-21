package cobra_cli

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	project2 "jirno/internal/pkg/delivery/cobra-cli/project"
	task2 "jirno/internal/pkg/delivery/cobra-cli/task"
	user2 "jirno/internal/pkg/delivery/cobra-cli/user"
	"jirno/internal/pkg/localstorage"
	sqlite_repo2 "jirno/internal/pkg/repository/project/sqlite_repo"
	"jirno/internal/pkg/repository/task/sqlite_repo"
	sqlite_repo3 "jirno/internal/pkg/repository/user/sqlite_repo"
	"jirno/internal/pkg/usecase/project"
	"jirno/internal/pkg/usecase/task"
	"jirno/internal/pkg/usecase/user"


	_ "github.com/mattn/go-sqlite3"
)

func RunApp() error {
	db, err := sql.Open("sqlite3", "jirno.db")
	if err != nil {
		return fmt.Errorf("db connection error: %v", err)
	}

	taskRepo := sqlite_repo.NewSQLiteTaskRepo(db)
	projectRepo := sqlite_repo2.NewSqliteProjectRepo(db)
	userRepo := sqlite_repo3.NewSqliteUserRepo(db)
	localStorage := localstorage.NewSQLiteLocalStorage(db)

	taskUsecase := task.NewTaskUsecase(taskRepo)
	projectUsecase := project.NewProjectUsecase(projectRepo)
	userUsercase := user.NewUserUsecase(userRepo)

	var rootCmd = &cobra.Command{Use: "jirno"}

	task2.NewTaskHandler(rootCmd, taskUsecase, localStorage)
	project2.NewProjectHandler(rootCmd, projectUsecase)
	user2.NewUserHandler(rootCmd, userUsercase, localStorage)

	return rootCmd.Execute()
}

