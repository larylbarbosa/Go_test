package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type OrganisationModule struct {
	ID         int
	OrgId      int
	ModuleType int
}

func CreateOrganisationModuleType(tx *sql.Tx, orgID int, modulesToAdd []int) (err error) {

	log.Println("createOrganisationModuleType : ", modulesToAdd)

	sqlStatement := `INSERT INTO organisations_modules (id, "orgId", "moduleType") VALUES (DEFAULT, $1, $2)`
	for i := 0; i < len(modulesToAdd); i++ {

		_, err = tx.Exec(sqlStatement, orgID, modulesToAdd[i])
		if err != nil {
			err = errors.New("createOrganisationModuleType , Error occured while inserting " + err.Error())
		} else {
			err = nil
		}
	}

	return err
}

func GetModuleNameByIds(tx *sql.Tx, moduleIDs []int) (moduleNames []string, err error) {

	var modString []string

	if len(moduleIDs) == 0 {
		return modString, nil
	}
	for _, ID := range moduleIDs {
		modString = append(modString, strconv.Itoa(ID))
	}
	modIdString := strings.Join(modString, "','")

	sqlRaw := fmt.Sprintf(`SELECT name FROM modules WHERE id IN ('%s')`, modIdString)

	log.Println("dbGetModuleNameByIds, sqlRaw: ", sqlRaw)

	log.Println(sqlRaw)

	rows, err := tx.Query(sqlRaw)
	if err != nil {
		err = errors.New("dbGetModuleNameByIds, Error occured while quering " + err.Error())
		return moduleNames, err
	}
	defer rows.Close()

	for rows.Next() {
		var moduleName string

		err = rows.Scan(&moduleName)
		if err != nil {
			err = errors.New("dbGetModuleNameByIds, Error while scanning " + err.Error())
		} else {
			log.Println("dbGetModuleNameByIds,", "Name:", moduleName)
			moduleNames = append(moduleNames, moduleName)
		}
	}
	log.Println("dbGetModuleNameByIds, moduleTypeArray: ", moduleNames)

	return moduleNames, err
}
