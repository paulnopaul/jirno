package sqlite_repo

import (
	"database/sql"
	"fmt"
	domain "jirno/internal/pkg/domain/task"

	"github.com/google/uuid"
)

type sqliteTaskRepo struct {
	db *sql.DB
}

func NewSQLiteTaskRepo(sqliteDB *sql.DB) domain.ITaskRepo {
	return sqliteTaskRepo{
		db: sqliteDB,
	}
}

func (s sqliteTaskRepo) GetByID(id uuid.UUID) (*domain.Task, error) {
	row := s.db.QueryRow(getByIDQuery, id.String())
	dRes := SQLiteTask{}
	err := row.Scan(&dRes.ID, &dRes.User, &dRes.Project,
		&dRes.Title, &dRes.Description,
		&dRes.IsCompleted, &dRes.CreatedDate, &dRes.CompletedDate, &dRes.DateTo)
	if err != nil {
		return nil, fmt.Errorf("task get by id failed: %v", err)
	}
	res, err := domainFromSQLite(&dRes)
	if err != nil {
		return nil, fmt.Errorf("task get by id (sqlite to domain) failed %v", err)
	}
	return res, nil
}

func (s sqliteTaskRepo) GetByFilter(filter domain.TaskFilter) ([]domain.Task, error) {
	sqliteFilter := filterFromDomain(filter)
	req, data, err := buildGetByFilterQuery(sqliteFilter)
	if err != nil {
		return nil, fmt.Errorf("task get by filter (build query) failed")
	}
	rows, err := s.db.Query(req, data...)
	if err != nil {
		return nil, fmt.Errorf("task get by filter (query) failed: %v", err)
	}
	res := make([]domain.Task, 0)
	for rows.Next() {
		rowDRes := SQLiteTask{}
		err = rows.Scan(&rowDRes.ID, &rowDRes.User, &rowDRes.Project,
			&rowDRes.Title, &rowDRes.Description,
			&rowDRes.IsCompleted, &rowDRes.CreatedDate,
			&rowDRes.CompletedDate, &rowDRes.DateTo)
		if err != nil {
			return nil, fmt.Errorf("task get by filter failed: %v", err)
		}
		rowRes, err := domainFromSQLite(&rowDRes)
		if err != nil {
			return nil, fmt.Errorf("task get by filter (sqlite to domain) failed %v", err)
		}
		res = append(res, *rowRes)
	}
	return res, nil
}

func (s sqliteTaskRepo) Create(task domain.Task) error {
	dbTask, err := sqliteFromDomain(task)
	if err != nil {
		return fmt.Errorf("task create (domain to sqlite) failed: %v", err)
	}
	_, err = s.db.Exec(createQuery,
		dbTask.ID, dbTask.User, dbTask.Project,
		dbTask.Title, dbTask.Description,
		dbTask.IsCompleted, dbTask.CreatedDate,
		dbTask.CompletedDate, dbTask.DateTo)
	if err != nil {
		return fmt.Errorf("task create failed: %v", err)
	}
	return nil
}

func (s sqliteTaskRepo) Update(task domain.TaskUpdate) error {
	req, data, err := buildUpdateQuery(task)
	if err != nil {
		return fmt.Errorf("task update (build query) failed: %v", err)
	}
	_, err = s.db.Exec(req, data...)
	if err != nil {
		return fmt.Errorf("task update failed: %v", err)
	}
	return nil
}

func (s sqliteTaskRepo) Delete(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM Tasks WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("task delete failed: %v", err)
	}
	return nil
}
