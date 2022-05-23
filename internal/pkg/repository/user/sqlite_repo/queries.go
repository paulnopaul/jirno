package sqlite_repo

import (
	sq "github.com/Masterminds/squirrel"
	domain "jirno/internal/pkg/domain/user"
)

func buildUpdateQuery(userUpdate domain.User) (string, []interface{}, error) {
	req := sq.Update("Users").Where(sq.Eq{"id": userUpdate.ID})
	if userUpdate.Email != "" {
		req = req.Set("email", userUpdate.Email)
	}
	if userUpdate.Nickname != "" {
		req = req.Set("nickname", userUpdate.Nickname)
	}
	if userUpdate.Name != "" {
		req = req.Set("name", userUpdate.Name)
	}
	if userUpdate.Password != nil {
		req = req.Set("password", userUpdate.Password)
	}
	res, data, err := req.ToSql()
	return res, data, err
}
