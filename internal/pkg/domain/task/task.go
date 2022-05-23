package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID            uuid.UUID
	User          int64
	Project       uuid.UUID
	Title         string
	Description   string
	IsCompleted   bool
	CreatedDate   time.Time
	CompletedDate *time.Time
	DateTo        *time.Time
}

type ITaskUsecase interface {
	Create(task Task) (uuid.UUID, error)
	GetByFilter(filter DeliveryTaskFilter) ([]Task, error)
	GetByID(id uuid.UUID) (*Task, error)
	Complete(id uuid.UUID) error
	Update(update TaskUpdate) error
	Delete(id uuid.UUID) error
}

//go:generate mockgen -destination=../../repository/task/mock/mock_repo.go -package=mock jirno/internal/pkg/domain/task ITaskRepo
type ITaskRepo interface {
	GetByID(id uuid.UUID) (*Task, error)
	GetByFilter(filter TaskFilter) ([]Task, error)
	Create(task Task) error
	Update(update TaskUpdate) error
	Delete(id uuid.UUID) error
}
