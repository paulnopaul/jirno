package sqlite_repo

import (
	sq "github.com/Masterminds/squirrel"
	domain "jirno/internal/pkg/domain/task"
)

const (
	getByIDQuery = "SELECT id, uid, pid, title, description,  is_completed, created_date, completed_date, date_to FROM Tasks WHERE id = ?"
	createQuery  = "INSERT INTO Tasks(id, uid, pid, title, description,  is_completed, created_date, completed_date, date_to) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

func buildGetByFilterQuery(filter SQLiteTaskFilter) (string, []interface{}, error) {
	req := sq.Select("id", "uid", "pid",
		"title", "description",
		"is_completed", "created_date", "completed_date", "date_to").
		From("Tasks")
	if filter.User != nil {
		req = req.Where(sq.Eq{"uid": *filter.User})
	}
	if filter.Project != "" {
		req = req.Where(sq.Eq{"pid": filter.Project})
	}
	if filter.StartDate != nil {
		req = req.Where(sq.GtOrEq{"date_to": *filter.StartDate})
	}
	if filter.EndDate != nil {
		req = req.Where(sq.LtOrEq{"date_to": *filter.EndDate})
	}
	return req.ToSql()
}

func buildUpdateQuery(update domain.TaskUpdate) (string, []interface{}, error) {
	req := sq.Update("Tasks").Where(sq.Eq{"id": update.ID.String()})
	if update.Project != nil {
		req = req.Set("pid", *update.Project)
	}
	if update.User != nil {
		req = req.Set("uid", *update.User)
	}
	if update.Title != "" {
		req = req.Set("title", update.Title)
	}
	if update.Description != "" {
		req = req.Set("description", update.Description)
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
