package sqlite_repo

import (
	"jirno/internal/pkg/domain/project"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type SQLiteProject struct {
	ID            string
	Title         string
	Description   string
	ParentProject string
	IsCompleted   bool
	CreatedDate   int64
	CompletedDate *int64
	DateTo        *int64
}

func sqliteProjectToDomainProject(src *SQLiteProject) (*project.Project, error) {
	parsedID, err := uuid.Parse(src.ID)
	if err != nil {
		return nil, err
	}
	parsedPID, err := uuid.Parse(src.ParentProject)
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
		*resDateTo = time.Unix(*src.DateTo, 0)
	}

	return &project.Project{
		ID:            parsedID,
		ParentProject: parsedPID,
		Title:         src.Title,
		Description:   src.Description,
		Users:         nil,
		Tasks:         nil,
		IsCompleted:   src.IsCompleted,
		CreatedDate:   resCreatedDate,
		CompletedDate: resCompletedDate,
		DateTo:        resDateTo,
	}, nil
}

func sqliteFromDomain(src project.Project) (*SQLiteProject, error) {
	res := &SQLiteProject{
		ID:            src.ID.String(),
		Title:         src.Title,
		Description:   src.Description,
		ParentProject: src.ParentProject.String(),
		IsCompleted:   src.IsCompleted,
		CreatedDate:   src.CreatedDate.Unix(),
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

type SQLiteProjectFilter struct {
	User          *int64
	StartDate     *int64
	EndDate       *int64
	ParentProject string
}

func buildGetByFilterQuery(filter SQLiteProjectFilter) (string, []interface{}, error) {
	req := sq.Select("id", "title", "description", "is_completed", "parent_pid", "created_date", "completed_date", "date_to").
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

func sqliteFilterFromDomain(filter project.ProjectFilter) SQLiteProjectFilter {
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
