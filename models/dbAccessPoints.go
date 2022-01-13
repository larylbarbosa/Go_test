package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type AccessPointSchema struct {
	ID               int
	Name             string
	OrgID            int
	SiteID           int
	AccessPointType  int
	InstallationType string
	ExtServicesState string
}

func dbCreateAccessPoint(tx *sql.Tx, ap AccessPointSchema) (accessPointID int, err error) {

	err = tx.QueryRow(`INSERT INTO access_points (name, "orgId", "siteId", "accessPointType", "installationType", "extServicesState") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		ap.Name, ap.OrgID, ap.SiteID, ap.AccessPointType, ap.InstallationType, ap.ExtServicesState).Scan(&accessPointID)
	if err != nil {
		err = errors.New("dbCreateAccessPoint, Error occured while inserting " + err.Error())
	} else {
		err = nil
	}

	return accessPointID, err
}

func dbGetAllOrganisationAccessPoints(tx *sql.Tx, orgID int) (accessPoints []AccessPointSchema, err error) {

	rows, err := tx.Query(`SELECT id, name, "siteId", "accessPointType", "installationType", "extServicesState" FROM access_points WHERE "orgId"=$1`, orgID)
	if err != nil {
		err = errors.New("dbGetAllOrganisationAccessPoints, Error occured while quering " + err.Error())
		return accessPoints, err
	}
	defer rows.Close()

	for rows.Next() {
		accessPoint := AccessPointSchema{}

		err = rows.Scan(&accessPoint.ID, &accessPoint.Name, &accessPoint.SiteID, &accessPoint.AccessPointType,
			&accessPoint.InstallationType, &accessPoint.ExtServicesState)
		if err != nil {
			err = errors.New("dbGetAllOrganisationAccessPoints, Error while scanning " + err.Error())
		} else {
			log.Println("dbGetAllOrganisationAccessPoints,", "ID:", accessPoint.ID, "Name:", accessPoint.Name, "SiteID:", accessPoint.SiteID)
			accessPoints = append(accessPoints, accessPoint)
		}
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		err = errors.New("dbGetAllOrganisationAccessPoints, Error occuer during iteration " + err.Error())
	}

	return accessPoints, err
}
func dbGetAccessPoint(tx *sql.Tx, accessPointID int) (accessPointExist bool, accessPoint AccessPointSchema, err error) {

	row := tx.QueryRow(`SELECT name, "orgId", "siteId", "accessPointType", "installationType", "extServicesState" FROM access_points WHERE id=$1`, accessPointID)

	err = row.Scan(&accessPoint.Name, &accessPoint.OrgID, &accessPoint.SiteID, &accessPoint.AccessPointType, &accessPoint.InstallationType, &accessPoint.ExtServicesState)
	switch err {
	case sql.ErrNoRows:
		accessPointExist, err = false, nil
	case nil:
		accessPointExist, err = true, nil
	default:
		accessPointExist, err = false, errors.New("dbGetAccessPoint, Unknown error, probably query field empty: "+err.Error())
	}

	return accessPointExist, accessPoint, err
}
func dbDeleteAccessPoint(tx *sql.Tx, accessPointID int) (err error) {

	var result sql.Result
	result, err = tx.Exec("DELETE FROM access_points WHERE id=$1;", accessPointID)
	if err != nil {
		err = errors.New("deleteAccessPoint, Error occured while deleting row: " + err.Error())
	} else {
		_, err = result.RowsAffected()
		if err != nil {
			err = errors.New("dbDeleteAccessPoint, Failed to verify delete: " + err.Error())
		} else {
			err = nil
		}
	}

	return err
}

func dbUpdateAccessPointDetail(tx *sql.Tx, accessPointID int, ap AccessPointSchema) (err error) {

	var result sql.Result

	result, err = tx.Exec(`UPDATE access_points SET name=$1, "accessPointType"=$2, "installationType"=$3 WHERE id=$4;`,
		ap.Name, ap.AccessPointType, ap.InstallationType, accessPointID)
	if err != nil {
		err = errors.New("dbUpdateAccessPointDetail, Error occured during update: " + err.Error())
	} else {
		_, err = result.RowsAffected()
		if err != nil {
			err = errors.New("dbUpdateAccessPointDetail, Failed to verify update: " + err.Error())
		} else {
			err = nil
		}
	}

	return err
}

func dbGetAccessPointCountInSite(tx *sql.Tx, siteID int) (accessPointCount int, err error) {

	err = tx.QueryRow(`SELECT COUNT(*) FROM access_points WHERE "siteId"=$1`, siteID).Scan(&accessPointCount)
	if err != nil {
		err = errors.New("dbGetAccessPointCountInSite, Error occured during query: " + err.Error())
	} else {
		fmt.Printf("dbGetAccessPointCountInSite, Number of access points in site are: %d\n", accessPointCount)
	}
	return accessPointCount, err
}

func dbSetAccessPointExtServisesState(tx *sql.Tx, accessPointID int, state string) (err error) {

	var result sql.Result

	result, err = tx.Exec(`UPDATE access_points SET "extServicesState"=$1 WHERE id=$2;`, state, accessPointID)
	if err != nil {
		err = errors.New("dbSetAccessPointExtServisesState, Failed to update access points table: " + err.Error())
	} else {
		_, err = result.RowsAffected()
		if err != nil {
			err = errors.New("dbSetAccessPointExtServisesState, Failed to verify update in access points table: " + err.Error())
		} else {
			err = nil
		}
	}

	return err
}
