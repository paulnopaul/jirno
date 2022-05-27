package sqlite_repo

import (
	sq "github.com/Masterminds/squirrel"
	"jirno/internal/pkg/domain/project"
)

const (
	getByIDQuery = "SELECT id, title, description, is_completed, parent_pid, created_date, completed_date, date_to FROM Projects WHERE id = ?"
	createQuery  = "INSERT INTO Projects(id, title, description, is_completed, parent_pid, created_date, completed_date, date_to) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

func buildUpdateQuery(update project.ProjectUpdate) (string, []interface{}, error) {
	req := sq.Update("Projects").Where(sq.Eq{"id": update.ID})
	if update.Title != "" {
		req = req.Set("title", update.Title)
	}
	if update.Description != "" {
		req = req.Set("description", update.Description)
	}
	if update.ParentProject != nil {
		req = req.Set("parent_pid", *update.ParentProject)
	}
	if update.IsCompleted != nil {
		req = req.Set("is_completed", *update.IsCompleted)
	}
	if update.CompletedDate != nil {
		req = req.Set("completed_date", (*update.CompletedDate).Unix())
	}
	if update.DateTo != nil {
		req = req.Set("date_to", (*update.DateTo).Unix())
	}
	return req.ToSql()
}
