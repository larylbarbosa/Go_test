package models

import (
	"database/sql"
	"errors"
)

type IntegrationTypeSchema struct {
	ID   int
	Name string
}

func CreateIntegrationType(tx *sql.Tx, integration IntegrationTypeSchema) (typeID int, err error) {

	err = tx.QueryRow(`INSERT INTO integration_types ("id","type") VALUES ($1,$2) RETURNING id`,
		integration.ID, integration.Name).Scan(&typeID)

	if err != nil {
		err = errors.New("dbCreateIntegrationType, Error occured while inserting " + err.Error())
		tx.Rollback()
	}

	return typeID, err
}
func CheckIfOrganisationIntegrationExist(tx *sql.Tx, integrationTypeID int) (orgIntegrationExists bool, orgID int, err error) {

	row := tx.QueryRow(`SELECT "orgId" FROM organisations_integrations WHERE "integrationType"=$1 LIMIT 1`, integrationTypeID)

	err = row.Scan(&orgID)
	switch err {
	case sql.ErrNoRows:
		orgIntegrationExists, err = false, nil
	case nil:
		orgIntegrationExists, err = true, nil
	default:
		orgIntegrationExists, err = false, errors.New("dbCheckIfOrganisationIntegrationExist, Unknown error, probably query field empty: "+err.Error())
	}

	return orgIntegrationExists, orgID, err
}
func DeleteIntegrationType(tx *sql.Tx, integrationTypeID int) (err error) {

	var result sql.Result
	result, err = tx.Exec(`DELETE FROM integration_types WHERE id=$1;`, integrationTypeID)
	if err != nil {

		return errors.New("dbDeleteIntegrationType  Error occured while deleting row" + err.Error())
	}
	_, err = result.RowsAffected()
	if err != nil {

		err = errors.New("dbDeleteIntegrationType, Failed to verify delete" + err.Error())
	}

	return err
}
