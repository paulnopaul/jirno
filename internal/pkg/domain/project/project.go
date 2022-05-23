package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID            uuid.UUID
	ParentProject uuid.UUID
	Users         []int64
	Tasks         []uuid.UUID
	Title         string
	Description   string
	Additional    map[string]string
	IsCompleted   bool
	CreatedDate   time.Time
	CompletedDate *time.Time
	DateTo        *time.Time
}

type IProjectUsecase interface {
	GetByID(id uuid.UUID) (*Project, error)
	GetByFilter(filter SmartProjectFilter) ([]Project, error)
	Create(project Project) (uuid.UUID, error)
	Update(project ProjectUpdate) error
	Complete(id uuid.UUID) error
	Delete(id uuid.UUID) error
}

//go:generate mockgen -destination=../repository/project/mock/mock_repo.go -package=mock jirno/internal/pkg/domain IProjectRepo
type IProjectRepo interface {
	GetByID(id uuid.UUID) (*Project, error)
	GetByFilter(filter ProjectFilter) ([]Project, error)
	Create(project Project) error
	Update(project ProjectUpdate) error
	Delete(id uuid.UUID) error
}
