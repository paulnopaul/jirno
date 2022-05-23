package task

import (
	"github.com/google/uuid"
	"time"
)

type TaskFilter struct {
	User      *int64
	StartDate *time.Time
	EndDate   *time.Time
	Project   *uuid.UUID
}

type DeliveryTaskFilter struct {
	Smart     string
	User      *int64
	StartDate *time.Time
	EndDate   *time.Time
	Project   string
}

func (f DeliveryTaskFilter) ToDomain() (*TaskFilter, error) {
	res := &TaskFilter{}
	if f.User != nil {
		res.User = new(int64)
		*res.User = *f.User
	}
	if f.StartDate != nil {
		res.StartDate = &time.Time{}
		*res.StartDate = *f.StartDate
	}
	if f.EndDate != nil {
		res.EndDate = &time.Time{}
		*res.EndDate = *f.EndDate
	}
	if f.Project != "" {
		parsedID, err := uuid.Parse(f.Project)
		if err != nil {
			return nil, err
		}
		res.Project = &parsedID
	}
	return res, nil
}
