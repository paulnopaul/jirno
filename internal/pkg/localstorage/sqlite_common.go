package localstorage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

func getUUIDByNumber(db *sql.DB, table string, number int) (uuid.UUID, error) {
	row := db.QueryRow("SELECT value from LocalStorage where field=?", table)
	var resBytes []byte
	err := row.Scan(&resBytes)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("localstorage %s failed: %v", table, err)
	}
	dbTask := map[int]string{}
	err = json.Unmarshal(resBytes, &dbTask)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("localstorage %s (unmarshalling) failed: %v", table, err)
	}
	value, found := dbTask[number]
	if !found {
		return uuid.UUID{}, fmt.Errorf("localstorage %s failed: id not found", table)
	}

	res, err := uuid.Parse(value)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("localstorage %s (uuid parsing) failed: %v", table, err)
	}
	return res, nil
}

func updateDBValue(db *sql.DB, table string, value string) error {
	_, err := db.Exec("DELETE FROM LocalStorage WHERE field=?", table)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO LocalStorage(field, value) values(?, ?)", table, value)
	if err != nil {
		return err
	}
	return nil
}
