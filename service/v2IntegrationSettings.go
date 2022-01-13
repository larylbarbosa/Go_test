package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yashdiniz/soa_infrastructure_management/models"
)

func (svc service) CreateIntegrationType(reqBody CreateIntegrationTypesRequest) error {
	var integrator models.IntegrationTypeSchema
	integrator.Name = reqBody.Name
	integrator.ID = reqBody.ID
	tx, err := svc.db.BeginTx(&gin.Context{}, nil)
	if err != nil {
		return fail("v2CreateIntegrationType", err)
	}
	defer tx.Rollback()

	_, err = models.CreateIntegrationType(tx, integrator)
	if err != nil {
		return fail("v2CreateIntegrationType", err)
	}

	tx.Commit()
	return nil

}

func (svc service) DeleteIntegrationType(integrationTypeID int) error {
	tx, err := svc.db.BeginTx(&gin.Context{}, nil)
	if err != nil {
		return fail("v2DeleteIntegrationType", err)
	}
	defer tx.Rollback()

	orgIntegrationExists, _, err := models.CheckIfOrganisationIntegrationExist(tx, integrationTypeID)
	if err != nil {
		return fail("v2DeleteIntegrationType", err)

	}

	if orgIntegrationExists {
		// TODO: Create a separate `errorCodes.go` in service package itself.
		return errors.New(`{"type":"error", "message":{ "errorCode":1236, "errorMessage":"Failed to delete integration type."}}`)
	}

	err = models.DeleteIntegrationType(tx, integrationTypeID)
	if err != nil {
		return fail("v2DeleteIntegrationType", err)
	}
	if err = tx.Commit(); err != nil {
		return fail("v2DeleteIntegrationType", err)
	}
	return nil
}
