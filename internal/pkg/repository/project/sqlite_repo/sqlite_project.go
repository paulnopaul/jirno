package sqlite_repo

import (
	"encoding/json"
	"jirno/internal/pkg/domain"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type SQLiteProject struct {
	ID            string
	Title         string
	Description   string
	ParentProject string
	Additional    string
	IsCompleted   bool
	CreatedDate   int64
	CompletedDate int64
	DateTo        int64
}

func sqliteProjectToDomainProject(src *SQLiteProject) (*domain.Project, error) {
	resAdditional := map[string]string{}
	err := json.Unmarshal([]byte(src.Additional), &resAdditional)
	if err != nil {
		return nil, err
	}

	parsedID, err := uuid.Parse(src.ID)
	if err != nil {
		return nil, err
	}
	parsedPID, err := uuid.Parse(src.ParentProject)
	if err != nil {
		return nil, err
	}

	resCreatedDate := time.Unix(src.CreatedDate, 0)
	resCompletedDate := time.Unix(src.CompletedDate, 0)
	resDateTo := time.Unix(src.CompletedDate, 0)

	return &domain.Project{
		ID:            parsedID,
		ParentProject: parsedPID,
		Title:         src.Title,
		Description:   src.Description,
		Users:         nil,
		Tasks:         nil,
		IsCompleted:   src.IsCompleted,
		Additional:    resAdditional,
		CreatedDate:   resCreatedDate,
		CompletedDate: &resCompletedDate,
		DateTo:        &resDateTo,
	}, nil
}

func domainProjectToSqliteProject(src domain.Project) (*SQLiteProject, error) {
	resAdditional, err := json.Marshal(src.Additional)
	if err != nil {
		return nil, err
	}
	res := &SQLiteProject{
		ID:            src.ID.String(),
		Title:         src.Title,
		Description:   src.Description,
		ParentProject: src.ParentProject.String(),
		IsCompleted:   src.IsCompleted,
		Additional:    string(resAdditional),
		CreatedDate:   src.CreatedDate.Unix(),
	}
	if src.CompletedDate != nil {
		res.CompletedDate = src.CompletedDate.Unix()
	}
	if src.CompletedDate != nil {
		res.DateTo = src.DateTo.Unix()
	}
	return res, nil
}

type SQLiteProjectFilter struct {
	User          *int64
	StartDate     *int64
	EndDate       *int64
	ParentProject string
}

func buildGetByFilterQuery(filter SQLiteProjectFilter) (string, []interface{}, error) {
	req := sq.Select("id", "title", "description", "additional", "is_completed", "parent_pid", "created_date", "completed_date", "date_to").
		From("Projects")
	if filter.User != nil {
		req = req.Where(sq.Eq{"uid": *filter.User})
	}
	if filter.ParentProject != "" {
		req = req.Where(sq.Eq{"parent_pid": filter.ParentProject})
	}
	if filter.StartDate != nil {
		req = req.Where(sq.GtOrEq{"date_to": *filter.StartDate})
	}
	if filter.EndDate != nil {
		req = req.Where(sq.LtOrEq{"date_to": *filter.EndDate})
	}
	return req.ToSql()
}

func sqliteFilterFromDomain(filter domain.ProjectFilter) SQLiteProjectFilter {
	res := SQLiteProjectFilter{}
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
	if filter.ParentProject != nil {
		res.ParentProject = filter.ParentProject.String()
	}
	return res
}
