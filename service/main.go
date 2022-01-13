// TODO: create independent service_test package for all mock requirements!
// (no need to make a different folder?? <- check)
package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var auth *Auth

type Service interface {
	AuthenticateUser(r *gin.Context) (scopes CustomUserScopes, err error)
	CheckIfPartnerExistsInScope(scopes CustomUserScopes, partnerID int) (partnerUserScope PartnerUserScope, partnerExist bool)
	GetOrganisationsUnderPartner(partnerId int) ([]Organisation, error)
	GetCognitoConfigs()
	StartPartnerManagementServer()
	CreateIntegrationType(reqBody CreateIntegrationTypesRequest) error
	AuthenticateWithApiKey(ctx *gin.Context) error
	DeleteIntegrationType(integrationTypeID int) error

	// NewOrganisationsUnderPartner(reqBody CreateOrganisationUnderPartnerRequest, partnerUserScope PartnerUserScope, ctx *gin.Context) (models.GatewaysSchema, error)
}

type service struct {
	db *sql.DB
}

/*
	Creates a New Service
*/
func New(db *sql.DB) *service {
	return &service{
		db,
	}
}
