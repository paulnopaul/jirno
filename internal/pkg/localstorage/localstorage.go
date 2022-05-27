package localstorage

import (
	"github.com/google/uuid"
	"jirno/internal/pkg/domain/project"
	"jirno/internal/pkg/domain/task"
	"jirno/internal/pkg/domain/user"
)

type localStorage struct {
	lastTasks    map[int]uuid.UUID
	lastProjects map[int]uuid.UUID
	currentUser  int64
}

type LocalStorage interface {
	GetTaskID(number int) (uuid.UUID, error)
	GetProjectID(number int) (uuid.UUID, error)
	GetUserID() (int64, error)

	SetTaskList(tasks []task.Task) error
	SetProjectList(projects []project.Project) error
	SetCurrentUser(user.User) error
}
