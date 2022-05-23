package sqlite_repo

import (
	"database/sql"
	"fmt"
	domain "jirno/internal/pkg/domain/project"

	"github.com/google/uuid"
)

type sqliteProjectRepo struct {
	db *sql.DB
}

func NewSqliteProjectRepo(sqliteDB *sql.DB) domain.IProjectRepo {
	return sqliteProjectRepo{
		db: sqliteDB,
	}
}
func (s sqliteProjectRepo) GetByID(id uuid.UUID) (*domain.Project, error) {
	res := &SQLiteProject{}
	row := s.db.QueryRow(getByIDQuery, id)
	err := row.Scan(&res.ID, &res.Title, &res.Description,
		&res.IsCompleted, &res.ParentProject,
		&res.CreatedDate, &res.CompletedDate, &res.DateTo)
	if err != nil {
		return nil, fmt.Errorf("project get by id failed: %v", err)
	}
	dRes, err := sqliteProjectToDomainProject(res)
	if err != nil {
		return nil, fmt.Errorf("project get by id (sqlite to domain) failed: %v", err)
	}
	return dRes, nil
}

func (s sqliteProjectRepo) Create(project domain.Project) error {
	dbProject, err := sqliteFromDomain(project)
	if err != nil {
		return fmt.Errorf("project create (domain to sqlite) failed: %v", err)
	}
	_, err = s.db.Exec(createQuery,
		dbProject.ID, dbProject.Title, dbProject.Description, dbProject.IsCompleted,
		dbProject.ParentProject, dbProject.CreatedDate, dbProject.CompletedDate, dbProject.DateTo)
	if err != nil {
		return fmt.Errorf("project create failed: %v", err)
	}
	return nil
}

func (s sqliteProjectRepo) Update(project domain.ProjectUpdate) error {
	// TODO handle users
	req, data, err := buildUpdateQuery(project)
	if err != nil {
		return fmt.Errorf("project update (build query) failed: %v", err)
	}
	_, err = s.db.Exec(req, data...)
	if err != nil {
		return fmt.Errorf("project update failed: %v", err)
	}
	return nil
}

func (s sqliteProjectRepo) Delete(id uuid.UUID) error {
	// TODO handle users
	// TODO handle tasks
	_, err := s.db.Exec("DELETE FROM Projects WHERE id = ?", id.String())
	if err != nil {
		return fmt.Errorf("project delete failed: %v", err)
	}
	return nil
}

func (s sqliteProjectRepo) GetByFilter(filter domain.ProjectFilter) ([]domain.Project, error) {
	sqliteFilter := sqliteFilterFromDomain(filter)
	req, data, err := buildGetByFilterQuery(sqliteFilter)
	if err != nil {
		return nil, fmt.Errorf("task get by filter (build query) failed")
	}
	rows, err := s.db.Query(req, data...)
	if err != nil {
		return nil, fmt.Errorf("task get by filter (query) failed: %v", err)
	}
	res := make([]domain.Project, 0)
	for rows.Next() {
		rowDRes := SQLiteProject{}
		err := rows.Scan(&rowDRes.ID, &rowDRes.Title, &rowDRes.Description,
			&rowDRes.IsCompleted, &rowDRes.ParentProject, &rowDRes.CreatedDate,
			&rowDRes.CompletedDate, &rowDRes.DateTo)
		if err != nil {
			return nil, fmt.Errorf("task get by filter failed: %v", err)
		}
		rowRes, err := sqliteProjectToDomainProject(&rowDRes)
		if err != nil {
			return nil, fmt.Errorf("task get by filter (sqlite to domain) failed %v", err)
		}
		res = append(res, *rowRes)
	}
	return res, nil
}
