package domain

import (
	"github.com/google/uuid"
	"jirno/internal/pkg/utils"
	"time"
)

type TaskFilter struct {
	User      *int64
	StartDate *time.Time
	EndDate   *time.Time
	Project   *uuid.UUID
}

type SmartTaskFilter struct {
	Smart     string
	User      *int64
	StartDate *time.Time
	EndDate   *time.Time
	Project   string
}

func (f SmartTaskFilter) ToDomain() (*TaskFilter, error) {
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
	res.handleSmart(f.Smart)
	return res, nil
}

func (filter *TaskFilter) handleSmart(smart string) {
	if smart == "" {
		return
	}

	filter.StartDate = new(time.Time)
	filter.EndDate = new(time.Time)

	now := time.Now()
	switch smart {
	case "today":
		*filter.StartDate, *filter.EndDate = utils.GetDayRange(now)
	case "tomorrow":
		*filter.StartDate, *filter.EndDate = utils.GetDayRange(now.AddDate(0, 0, -1))
	}
}
