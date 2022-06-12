package psql_repo

import (
	"encoding/json"
	"jirno/internal/pkg/domain"
	"time"

	"github.com/google/uuid"
)

type SQLiteTask struct {
	ID            string
	Project       string
	User          int64
	Title         string
	Description   string
	Additional    string
	IsCompleted   bool
	CreatedDate   int64
	CompletedDate *int64
	DateTo        *int64
}

type SQLiteTaskFilter struct {
	User      *int64
	StartDate *int64
	EndDate   *int64
	Project   string
}

func filterFromDomain(filter domain.TaskFilter) SQLiteTaskFilter {
	res := SQLiteTaskFilter{}
	if filter.User != nil {
		res.User = new(int64)
		*res.User = *filter.User
	}
	if filter.StartDate != nil {
		res.StartDate = new(int64)
		*res.StartDate = filter.StartDate.Unix()
	}
	if filter.EndDate != nil {
		res.EndDate = new(int64)
		*res.EndDate = filter.EndDate.Unix()
	}
	if filter.Project != nil {
		res.Project = filter.Project.String()
	}
	return res
}

func domainFromSQLite(src *SQLiteTask) (*domain.Task, error) {
	resAdditional := map[string]string{}
	err := json.Unmarshal([]byte(src.Additional), &resAdditional)
	if err != nil {
		return nil, err
	}

	parsedID, err := uuid.Parse(src.ID)
	if err != nil {
		return nil, err
	}

	parsedPID, err := uuid.Parse(src.Project)
	if err != nil {
		return nil, err
	}

	resCreatedDate := time.Unix(src.CreatedDate, 0)

	var resCompletedDate *time.Time
	if src.CompletedDate != nil {
		resCompletedDate = &time.Time{}
		*resCompletedDate = time.Unix(*src.CompletedDate, 0)
	}

	var resDateTo *time.Time
	if src.DateTo != nil {
		resDateTo = &time.Time{}
		*resDateTo = time.Unix(*src.CompletedDate, 0)
	}

	return &domain.Task{
		ID:            parsedID,
		Project:       parsedPID,
		User:          src.User,
		Title:         src.Title,
		Description:   src.Description,
		IsCompleted:   src.IsCompleted,
		Additional:    resAdditional,
		CreatedDate:   resCreatedDate,
		CompletedDate: resCompletedDate,
		DateTo:        resDateTo,
	}, nil
}

func sqliteFromDomain(src domain.Task) (*SQLiteTask, error) {
	resAdditional, err := json.Marshal(src.Additional)
	if err != nil {
		return nil, err
	}

	res := &SQLiteTask{
		ID:          src.ID.String(),
		Project:     src.Project.String(),
		User:        src.User,
		Title:       src.Title,
		Description: src.Description,
		IsCompleted: src.IsCompleted,
		Additional:  string(resAdditional),
		CreatedDate: src.CreatedDate.Unix(),
	}

	if src.CompletedDate != nil {
		res.CompletedDate = new(int64)
		*res.CompletedDate = src.CompletedDate.Unix()
	}

	if src.DateTo != nil {
		res.DateTo = new(int64)
		*res.DateTo = src.DateTo.Unix()
	}

	return res, nil
}
