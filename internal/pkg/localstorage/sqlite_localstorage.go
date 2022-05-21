package localstorage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"jirno/internal/pkg/domain"
	"strconv"
)

type sqliteLocalStorage struct {
	data localStorage
	db   *sql.DB
}

func NewSQLiteLocalStorage(sqliteConnector *sql.DB) LocalStorage {
	return &sqliteLocalStorage{
		db: sqliteConnector,
	}
}
func (s sqliteLocalStorage) GetTaskID(number int) (uuid.UUID, error) {
	return getUUIDByNumber(s.db, "last_tasks", number)
}

func (s sqliteLocalStorage) GetProjectID(number int) (uuid.UUID, error) {
	return getUUIDByNumber(s.db, "last_projects", number)
}

func (s sqliteLocalStorage) GetUserID() (int64, error) {
	row := s.db.QueryRow("SELECT value from LocalStorage where field='current_user'")
	var resStr string
	err := row.Scan(&resStr)
	if err != nil {
		return 0, fmt.Errorf("localstorage get current User failed: %v", err)
	}
	res, err := strconv.ParseInt(resStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("localstorage get current user (int parsing) failed: %v", err)
	}
	return res, nil
}

func (s sqliteLocalStorage) SetTaskList(tasks []domain.Task) error {
	idMap := map[int]string{}
	for i, value := range tasks {
		idMap[i] = value.ID.String()
	}
	res, err := json.Marshal(idMap)
	if err != nil {
		return fmt.Errorf("localstorage set task list (marshaling) error: %v", err)
	}
	return updateDBValue(s.db, "last_tasks", string(res))
}

func (s sqliteLocalStorage) SetProjectList(projects []domain.Project) error {
	idMap := map[int]string{}
	for i, value := range projects {
		idMap[i] = value.ID.String()
	}
	res, err := json.Marshal(idMap)
	if err != nil {
		return fmt.Errorf("localstorage set project list (marshaling) error: %v", err)
	}
	return updateDBValue(s.db, "last_tasks", string(res))
}

func (s sqliteLocalStorage) SetCurrentUser(user domain.User) error {
	return updateDBValue(s.db, "current_user", fmt.Sprint(user.ID))
}
