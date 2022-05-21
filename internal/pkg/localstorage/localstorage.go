package localstorage

import (
	"github.com/google/uuid"
	"jirno/internal/pkg/domain"
)

type localStorage struct {
	lastTasks map[int]uuid.UUID
	lastProjects map[int]uuid.UUID
	currentUser int64
}

type LocalStorage interface {
	GetTaskID(number int) (uuid.UUID, error)
	GetProjectID(number int) (uuid.UUID, error)
	GetUserID() (int64, error)

	SetTaskList(tasks []domain.Task) error
	SetProjectList(projects []domain.Project) error
	SetCurrentUser(domain.User) error
}

