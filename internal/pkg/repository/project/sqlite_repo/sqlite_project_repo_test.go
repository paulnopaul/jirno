package sqlite_repo

import (
	domain "jirno/internal/pkg/domain/project"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func testProject() domain.Project {
	complDate := time.Date(2010, time.October, 11, 1, 1, 1, 0, time.Local)
	testProject := domain.Project{
		ID:            uuid.New(),
		ParentProject: uuid.New(),
		Title:         "ProjectTitle",
		Description:   "Project description",
		IsCompleted:   true,
		CreatedDate:   time.Date(2010, time.October, 10, 1, 1, 1, 0, time.Local),
		CompletedDate: &complDate,
		DateTo:        nil,
	}
	return testProject
}

func TestSqliteProjectRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	testRes := testProject()
	testProject, err := sqliteFromDomain(testRes)
	if err != nil {
		t.Fatalf("sqlite task copy fail: %v", err)
	}

	query := "SELECT id, title, description, is_completed, parent_pid, created_date, completed_date, date_to"
	taskRow := sqlmock.NewRows([]string{"id", "title", "description",
		"is_completed", "parent_pid", "created_date", "completed_date", "date_to"}).
		AddRow(testProject.ID, testProject.Title, testProject.Description,
			testProject.IsCompleted, testProject.ParentProject,
			testProject.CreatedDate, testProject.CompletedDate, testProject.DateTo)

	mock.ExpectQuery(query).WithArgs(testProject.ID).WillReturnRows(taskRow)

	repo := NewSqliteProjectRepo(db)
	res, err := repo.GetByID(testRes.ID)
	assert.Equal(t, nil, err)
	assert.Equal(t, testRes, *res)
}

func TestSqliteProjectRepo_GetByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	data := []domain.Project{
		testProject(), testProject()}
	data[1].ParentProject = uuid.New()

	sqliteData := make([]*SQLiteProject, 2)
	for i, task := range data {
		sqliteData[i], err = sqliteFromDomain(task)
		if err != nil {
			t.Fatalf("sqlite task copy fail: %v", err)
		}
	}
	if err != nil {
		t.Fatalf("sqlite task copy fail: %v", err)
	}

	query := "SELECT id, title, description,  is_completed, parent_pid, created_date, completed_date, date_to"
	taskRows := sqlmock.NewRows([]string{"id", "title", "description",
		"parent_pid", "is_completed",
		"created_date", "completed_date", "date_to"})

	for _, task := range sqliteData {
		taskRows.AddRow(task.ID, task.Title, task.Description,
			task.IsCompleted, task.ParentProject,
			task.CreatedDate, task.CompletedDate, task.DateTo)
	}

	testFilter := domain.ProjectFilter{
		User:    new(int64),
		EndDate: new(time.Time),
	}
	*testFilter.User = 1
	*testFilter.EndDate = time.Date(2010, time.November, 10, 1, 1, 1, 0, time.Local)

	sqliteTestFilter := sqliteFilterFromDomain(testFilter)

	mock.ExpectQuery(query).
		WithArgs(*sqliteTestFilter.User, *sqliteTestFilter.EndDate).
		WillReturnRows(taskRows)

	repo := NewSqliteProjectRepo(db)
	res, err := repo.GetByFilter(testFilter)

	assert.Nil(t, err)
	assert.Equal(t, data[0], res[0])
	assert.Equal(t, data[1], res[1])
}

func TestSqliteProjectRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	testDomainProject := testProject()
	dbProject, err := sqliteFromDomain(testDomainProject)

	mock.ExpectExec("INSERT INTO Projects").
		WithArgs(
			dbProject.ID, dbProject.Title, dbProject.Description,
			dbProject.IsCompleted, dbProject.ParentProject,
			dbProject.CreatedDate, dbProject.CompletedDate, dbProject.DateTo,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewSqliteProjectRepo(db)
	err = repo.Create(testDomainProject)
	assert.Nil(t, err)
}
