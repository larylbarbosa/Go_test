package models

import (
	"database/sql"
	"errors"
)

type ErrorLogsSchema struct {
	ID      int
	Message string
}

func CreateErrorLog(tx *sql.Tx, log ErrorLogsSchema) (logID int, err error) {

	err = tx.QueryRow(`INSERT INTO error_logs ("message") VALUES ($1) RETURNING id`, log.Message).Scan(&logID)
	if err != nil {
		err = errors.New("dbCreateErrorLog, Error occured while inserting " + err.Error())
	} else {
		err = nil
	}

	return logID, err
}
