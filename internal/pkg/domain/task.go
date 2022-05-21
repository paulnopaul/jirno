package domain

import (
	"time"

	"github.com/google/uuid"
)

type DeliveryTask struct {
	ID            string
	User          *int64
	Project       string
	Title         string
	Description   string
	IsCompleted   *bool
	CompletedDate *time.Time
	DateTo        *time.Time
}

func (d DeliveryTask) ToUpdate() (*TaskUpdate, error) {
	res := &TaskUpdate{}
	if d.User != nil {
		res.User = new(int64)
		*res.User = *d.User
	}
	if d.Project != "" {
		parsedPId, err := uuid.Parse(d.Project)
		if err != nil {
			return nil, err
		}
		res.Project = new(uuid.UUID)
		*res.Project = parsedPId
	}
	res.Title = d.Title
	res.Description = d.Description
	if d.IsCompleted != nil {
		res.IsCompleted = new(bool)
		*res.IsCompleted = *d.IsCompleted
	}
	if d.CompletedDate != nil {
		res.CompletedDate = new(time.Time)
		*res.CompletedDate = *d.CompletedDate
	}
	if d.DateTo != nil {
		res.DateTo = new(time.Time)
		*res.DateTo = *d.DateTo
	}
	return res, nil
}

func (d DeliveryTask) ToDomain() (*Task, error) {
	res := &Task{}
	if d.User != nil {
		res.User = *d.User
	}
	if d.Project != "" {
		parsedPId, err := uuid.Parse(d.Project)
		if err != nil {
			return nil, err
		}
		res.Project = parsedPId
	}
	res.Title = d.Title
	res.Description = d.Description
	if d.IsCompleted != nil {
		res.IsCompleted = *d.IsCompleted
	}
	if d.CompletedDate != nil {
		res.CompletedDate = new(time.Time)
		*res.CompletedDate = *d.CompletedDate
	}
	if d.DateTo != nil {
		res.DateTo = new(time.Time)
		*res.DateTo = *d.DateTo
	}
	return res, nil
}

type TaskUpdate struct {
	ID            uuid.UUID
	User          *int64
	Project       *uuid.UUID
	Title         string
	Description   string
	Additional    map[string]string
	IsCompleted   *bool
	CompletedDate *time.Time
	DateTo        *time.Time
}

type Task struct {
	ID            uuid.UUID
	User          int64
	Project       uuid.UUID
	Title         string
	Description   string
	Additional    map[string]string
	IsCompleted   bool
	CreatedDate   time.Time
	CompletedDate *time.Time
	DateTo        *time.Time
}

type ITaskUsecase interface {
	Create(task Task) (uuid.UUID, error)
	GetByFilter(filter SmartTaskFilter) ([]Task, error)
	GetByID(id uuid.UUID) (*Task, error)
	Complete(id uuid.UUID) error
	Update(update TaskUpdate) error
	Delete(id uuid.UUID) error
}

//go:generate mockgen -destination=../repository/task/mock/mock_repo.go -package=mock jirno/internal/pkg/domain ITaskRepo
type ITaskRepo interface {
	GetByID(id uuid.UUID) (*Task, error)
	GetByFilter(filter TaskFilter) ([]Task, error)
	Create(task Task) error
	Update(update TaskUpdate) error
	Delete(id uuid.UUID) error
}
