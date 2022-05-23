package project

import (
	"github.com/google/uuid"
	"time"
)

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
