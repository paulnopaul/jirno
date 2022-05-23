package sqlite_repo

import (
	"database/sql"
	"fmt"
	domain "jirno/internal/pkg/domain/user"
)

type sqliteUserRepo struct {
	db *sql.DB
}

func NewSqliteUserRepo(sqliteDB *sql.DB) domain.IUserRepo {
	return sqliteUserRepo{
		db: sqliteDB,
	}
}

func (s sqliteUserRepo) GetByID(id int64) (*domain.User, error) {
	res := &domain.User{}
	row := s.db.QueryRow("SELECT id, name, nickname, email, password from Users WHERE id = ?", id)
	err := row.Scan(&res.ID, &res.Name, &res.Nickname, &res.Email, &res.Password)
	if err != nil {
		return nil, fmt.Errorf("user get by id failed: %v", err)
	}
	return res, nil
}

func (s sqliteUserRepo) GetByNickname(nickname string) (*domain.User, error) {
	res := &domain.User{}
	row := s.db.QueryRow("SELECT id, name, nickname, email, password from Users WHERE nickname = ?", nickname)
	err := row.Scan(&res.ID, &res.Name, &res.Nickname, &res.Email, &res.Password)
	if err != nil {
		return nil, fmt.Errorf("user get by nickname failed: %v", err)
	}
	return res, nil
}

func (s sqliteUserRepo) Create(user domain.User) error {
	_, err := s.db.Exec("INSERT INTO Users(id, name, nickname, email, password) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("user create failed: %v", err)
	}
	return nil
}

func (s sqliteUserRepo) Update(user domain.User) error {
	req, data, err := buildUpdateQuery(user)
	if err != nil {
		return fmt.Errorf("user update (query build) failed: %v", err)
	}
	_, err = s.db.Exec(req, data...)
	if err != nil {
		return fmt.Errorf("user update failed: %v", err)
	}
	return nil
}

func (s sqliteUserRepo) Delete(id int64) error {
	_, err := s.db.Exec("DELETE FROM Users where id = ?", id)
	if err != nil {
		return fmt.Errorf("user delete failed: %v", err)
	}
	return nil
}

func (s sqliteUserRepo) GetMaxUserID() (int64, error) {
	var res *int64
	row := s.db.QueryRow("SELECT max(id) from Users")
	err := row.Scan(&res)
	if err != nil {
		return 0, fmt.Errorf("user get by id failed: %v", err)
	}
	if res == nil {
		res = new(int64)
		*res = 0
	}
	return *res, nil
}
