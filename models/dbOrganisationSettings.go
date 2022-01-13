package models

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type OrganisationSettingsSchema struct {
	ID              int
	PartnerId       int
	OrgId           int
	Integration     bool
	CardType        string
	CardReadBaseKey uuid.UUID
}

func CreateOrganisationSettings(tx *sql.Tx, orgSetting OrganisationSettingsSchema) (err error) {

	sqlStatement := `INSERT INTO organisation_settings (id, "partnerId", "orgId", "cardType", "integration", "cardReadBaseKey") VALUES (DEFAULT, $1, $2, $3, $4, $5)`
	_, err = tx.Exec(sqlStatement, orgSetting.PartnerId, orgSetting.OrgId, orgSetting.CardType, orgSetting.Integration, orgSetting.CardReadBaseKey)
	if err != nil {
		err = errors.New("dbCreateOrganisationSettings, Error occured while inserting " + err.Error())
	} else {
		err = nil
	}
	return err
}
