package psql_repo

import (
	"jirno/internal/pkg/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func testTask() domain.Task {
	complDate := time.Date(2010, time.October, 11, 1, 1, 1, 0, time.Local)
	testTask := domain.Task{
		ID:            uuid.New(),
		User:          1,
		Project:       uuid.New(),
		Title:         "TaskTitle",
		Description:   "Task description",
		Additional:    map[string]string{},
		IsCompleted:   true,
		CreatedDate:   time.Date(2010, time.October, 10, 1, 1, 1, 0, time.Local),
		CompletedDate: &complDate,
		DateTo:        nil,
	}
	return testTask
}

func TestSqliteTaskRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	testRes := testTask()
	testTask, err := sqliteFromDomain(testRes)
	if err != nil {
		t.Fatalf("sqlite task copy fail: %v", err)
	}

	query := "SELECT id, uid, pid, title, description, additional, is_completed, created_date, completed_date, date_to"
	taskRow := sqlmock.NewRows([]string{"id", "uid", "pid",
		"title", "description", "additional",
		"is_completed", "created_date", "completed_date", "date_to"}).
		AddRow(testTask.ID, testTask.User, testTask.Project,
			testTask.Title, testTask.Description, testTask.Additional,
			testTask.IsCompleted, testTask.CreatedDate, testTask.CompletedDate, testTask.DateTo)

	mock.ExpectQuery(query).WithArgs(testTask.ID).WillReturnRows(taskRow)

	repo := NewSQLiteTaskRepo(db)
	res, err := repo.GetByID(testRes.ID)
	assert.Equal(t, nil, err)
	assert.Equal(t, testRes, *res)
}

func TestSqliteTaskRepo_GetByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	data := []domain.Task{
		testTask(), testTask()}
	data[1].Project = uuid.New()

	sqliteData := make([]*SQLiteTask, 2)
	for i, task := range data {
		sqliteData[i], err = sqliteFromDomain(task)
		if err != nil {
			t.Fatalf("sqlite task copy fail: %v", err)
		}
	}
	if err != nil {
		t.Fatalf("sqlite task copy fail: %v", err)
	}

	query := "SELECT id, uid, pid, title, description, additional, is_completed, created_date, completed_date, date_to"
	taskRows := sqlmock.NewRows([]string{"id", "uid", "pid",
		"title", "description", "additional",
		"is_completed", "created_date", "completed_date", "date_to"})

	for _, task := range sqliteData {
		taskRows.AddRow(task.ID, task.User, task.Project,
			task.Title, task.Description, task.Additional,
			task.IsCompleted, task.CreatedDate, task.CompletedDate, task.DateTo)
	}

	testFilter := domain.TaskFilter{
		User:    new(int64),
		EndDate: new(time.Time),
	}
	*testFilter.User = 1
	*testFilter.EndDate = time.Date(2010, time.November, 10, 1, 1, 1, 0, time.Local)

	sqliteTestFilter := filterFromDomain(testFilter)

	mock.ExpectQuery(query).
		WithArgs(*sqliteTestFilter.User, *sqliteTestFilter.EndDate).
		WillReturnRows(taskRows)

	repo := NewSQLiteTaskRepo(db)
	res, err := repo.GetByFilter(testFilter)

	assert.Nil(t, err)
	assert.Equal(t, data[0], res[0])
	assert.Equal(t, data[1], res[1])
}

func TestSqliteTaskRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	testDomainTask := testTask()
	dbTask, err := sqliteFromDomain(testDomainTask)

	mock.ExpectExec("INSERT INTO Tasks").
		WithArgs(
			dbTask.ID, dbTask.User, dbTask.Project,
			dbTask.Title, dbTask.Description, dbTask.Additional,
			dbTask.IsCompleted, dbTask.CreatedDate, dbTask.CompletedDate, dbTask.DateTo,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewSQLiteTaskRepo(db)
	err = repo.Create(testDomainTask)
	assert.Nil(t, err)
}

func TestSqliteTaskRepo_Update(t *testing.T) {

}

func TestSqliteTaskRepo_Delete(t *testing.T) {

}
