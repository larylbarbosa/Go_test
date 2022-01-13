package models

import (
	"database/sql"
	"errors"
)

type OrganisationsSchema struct {
	ID               int
	Name             string
	Type             string
	Country          string
	PartnerID        int
	IntegratorID     int
	AdminName        string
	AdminMobile      string
	AdminEmail       string
	AdminEditable    bool
	ExtServicesState string
}

func GetAllPartnerOrganisations(tx *sql.Tx, partnerID int) (organisations []OrganisationsSchema, err error) {

	rows, err := tx.Query(`SELECT id, name, "extServicesState" FROM organisations WHERE "partnerId"=$1 ORDER BY id DESC`, partnerID)
	if err != nil {
		err = errors.New("dbGetAllPartnerOrganisations, Error occured while quering " + err.Error())
		return organisations, err
	}
	defer rows.Close()

	for rows.Next() {
		org := OrganisationsSchema{}

		err = rows.Scan(&org.ID, &org.Name, &org.ExtServicesState)
		if err != nil {
			err = errors.New("dbGetAllPartnerOrganisations, Error while scanning " + err.Error())
		} else {
			organisations = append(organisations, org)
		}
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		err = errors.New("dbGetAllPartnerOrganisations, Error occuer during iteration " + err.Error())
	}

	return organisations, err
}
func CreateOrganisation(tx *sql.Tx, org OrganisationsSchema) (orgID int, err error) {

	err = tx.QueryRow(`INSERT INTO organisations (name, type, country, "partnerId", "integratorId", "adminName", "adminMobile", "adminEmail", "adminEditable", "extServicesState") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		org.Name, org.Type, org.Country, org.PartnerID, org.IntegratorID, org.AdminName, org.AdminMobile, org.AdminEmail, org.AdminEditable, org.ExtServicesState).Scan(&orgID)

	if err != nil {
		err = errors.New("CreateOrganisation, Error occured while inserting " + err.Error())
	} else {
		err = nil
	}

	return orgID, err
}

func SetOrganisationExtServisesState(tx *sql.Tx, orgID int, state string) (err error) {

	var result sql.Result

	result, err = tx.Exec(`UPDATE organisations SET "extServicesState"=$1 WHERE id=$2;`, state, orgID)
	if err != nil {
		err = errors.New("dbSetOrganisationExtServisesState, Failed to update organisations table: " + err.Error())
	} else {
		_, err = result.RowsAffected()
		if err != nil {
			err = errors.New("dbSetOrganisationExtServisesState, Failed to verify update in organisations table: " + err.Error())
		} else {
			err = nil
		}
	}

	return err
}
