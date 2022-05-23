package project

import (
	"github.com/google/uuid"
	"jirno/internal/pkg/utils"
	"time"
)

type ProjectFilter struct {
	User          *int64
	StartDate     *time.Time
	EndDate       *time.Time
	ParentProject *uuid.UUID
}

type SmartProjectFilter struct {
	Smart         string
	User          *int64
	StartDate     *time.Time
	EndDate       *time.Time
	ParentProject string
}

func (f SmartProjectFilter) ToDomain() (*ProjectFilter, error) {
	res := &ProjectFilter{}
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
	if f.ParentProject != "" {
		parsedID, err := uuid.Parse(f.ParentProject)
		if err != nil {
			return nil, err
		}
		res.ParentProject = &parsedID
	}
	res.handleSmart(f.Smart)
	return res, nil
}

func (filter *ProjectFilter) handleSmart(smart string) {
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
