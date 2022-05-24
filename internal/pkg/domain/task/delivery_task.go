package task

import (
	"github.com/google/uuid"
	"time"
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
