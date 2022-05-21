package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProjectID uuid.UUID

type DeliveryProject struct {
	ID            string
	Users         []int64
	ParentProject string
	Title         string
	Description   string
	IsCompleted   *bool
	CompletedDate *time.Time
	DateTo        *time.Time
}

func (d DeliveryProject) ToUpdate() (*ProjectUpdate, error) {
	res := &ProjectUpdate{}
	if d.Users != nil {
		res.Users = d.Users
	}

	if d.ParentProject != "" {
		parsedPId, err := uuid.Parse(d.ParentProject)
		if err != nil {
			return nil, err
		}
		res.ParentProject = new(uuid.UUID)
		*res.ParentProject = parsedPId
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

func (d DeliveryProject) ToDomain() (*Project, error) {
	res := &Project{}
	if d.Users != nil {
		res.Users = d.Users
	}
	if d.ParentProject != "" {
		parsedPId, err := uuid.Parse(d.ParentProject)
		if err != nil {
			return nil, err
		}
		res.ParentProject = parsedPId
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

type ProjectUpdate struct {
	ID            uuid.UUID
	Title         string
	Description   string
	ParentProject *uuid.UUID
	Additional    map[string]string
	Users         []int64
	IsCompleted   *bool
	CompletedDate *time.Time
	DateTo        *time.Time
}

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

// TODO project create returns id
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
