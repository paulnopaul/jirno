package sqlite_repo

import (
	"fmt"
	"jirno/internal/pkg/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSqliteUserRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}
	dbUser := &domain.User{
		ID:       1,
		Name:     "Name1",
		Nickname: "Nickname1",
		Email:    "email@mail.ru",
		Password: []byte("sdfojisdfoijsdf"),
	}
	query := fmt.Sprintf("SELECT id, name, nickname, email, password")
	userRow := sqlmock.NewRows([]string{"id", "name", "nickname", "email", "password"}).
		AddRow(dbUser.ID, dbUser.Name, dbUser.Nickname, dbUser.Email, dbUser.Password)
	mock.ExpectQuery(query).WithArgs(dbUser.ID).WillReturnRows(userRow)

	repo := NewSqliteUserRepo(db)
	res, err := repo.GetByID(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, dbUser, res)
}

func TestSqliteUserRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	dbUser := &domain.User{
		ID:       1,
		Name:     "Name1",
		Nickname: "Nickname1",
		Email:    "email@mail.ru",
		Password: []byte("sdfojisdfoijsdf"),
	}

	mock.ExpectExec("INSERT INTO Users(id, name, nickname, email, password)*").WithArgs(dbUser.ID, dbUser.Name, dbUser.Nickname, dbUser.Email, dbUser.Password).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewSqliteUserRepo(db)
	err = repo.Create(*dbUser)
	assert.Equal(t, nil, err)
}

func TestSqliteUserRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	var deleteID int64 = 10
	mock.ExpectExec("DELETE FROM Users where id = ?").WithArgs(deleteID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewSqliteUserRepo(db)
	err = repo.Delete(deleteID)
	assert.Equal(t, nil, err)
}

func TestSqliteUserRepo_GetByNickname(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}
	dbUser := &domain.User{
		ID:       1,
		Name:     "Name1",
		Nickname: "Nickname1",
		Email:    "email@mail.ru",
		Password: []byte("sdfojisdfoijsdf"),
	}
	query := "SELECT id, name, nickname, email, password"
	userRow := sqlmock.NewRows([]string{"id", "name", "nickname", "email", "password"}).
		AddRow(dbUser.ID, dbUser.Name, dbUser.Nickname, dbUser.Email, dbUser.Password)
	mock.ExpectQuery(query).WithArgs(dbUser.Nickname).WillReturnRows(userRow)

	repo := NewSqliteUserRepo(db)
	res, err := repo.GetByNickname(dbUser.Nickname)
	assert.Equal(t, nil, err)
	assert.Equal(t, dbUser, res)
}

func TestSqliteUserRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock fail: %v", err)
	}

	userUpdate := domain.User{
		ID:       1,
		Email:    "123",
		Nickname: "abc",
		Name:     "newname",
		Password: []byte("newpass"),
	}

	mock.ExpectExec("UPDATE Users").
		WithArgs(userUpdate.Email, userUpdate.Nickname, userUpdate.Name, userUpdate.Password, userUpdate.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewSqliteUserRepo(db)
	err = repo.Update(userUpdate)
	assert.Equal(t, nil, err)
}
