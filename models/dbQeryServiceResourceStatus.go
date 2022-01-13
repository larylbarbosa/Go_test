package models

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type ServiceResourceStatusSchema struct {
	ID                int
	ResourceName      string
	ResourceOperation string
	ResourceID        int
	ServiceName       string
	ServiceStatus     string
	ErrorLogID        int
	TransactionID     uuid.UUID
}

func CreateServiceResourceStatusLog(tx *sql.Tx, statusLog ServiceResourceStatusSchema) (err error) {

	err = tx.QueryRow(`INSERT INTO service_resource_status ("resourceName", "resourceOperation", "resourceId", "serviceName", "serviceStatus", "errorLogId", "transactionID") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		statusLog.ResourceName, statusLog.ResourceOperation, statusLog.ResourceID, statusLog.ServiceName, statusLog.ServiceStatus, statusLog.ErrorLogID, statusLog.TransactionID).Scan(&statusLog.ID)
	if err != nil {
		err = errors.New("dbCreateServiceResourceStatusLog, Error occured while inserting " + err.Error())
	} else {
		err = nil
	}

	return err
}
