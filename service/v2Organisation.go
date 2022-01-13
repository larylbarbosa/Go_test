package service

import (
	"github.com/gin-gonic/gin"
	"github.com/larylbarbosa/Go_test/models"
)

type Organisation struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ResourceState string `json:"resourceState"`
}

//GetOrganisationsUnderPartner returns all organisations under a partner
func (svc service) GetOrganisationsUnderPartner(partnerId int) ([]Organisation, error) {
	var result []Organisation
	tx, err := svc.db.BeginTx(&gin.Context{}, nil) // creating transaction, but no need to commit, since no DB updates...
	if err != nil {
		return result, fail("svc.GetOrganisationsUnderPartner", err)
	}
	defer tx.Rollback()

	orgs, err := models.GetAllPartnerOrganisations(tx, partnerId)
	if err != nil {
		return result, err
	}
	for _, org := range orgs {
		result = append(result, Organisation{
			ID:            org.ID,
			Name:          org.Name,
			ResourceState: org.ExtServicesState,
		})
	}
	return result, nil

}

// func (svc service) NewOrganisationsUnderPartner(reqBody CreateOrganisationUnderPartnerRequest, partnerUserScope PartnerUserScope, ctx *gin.Context) (models.GatewaysSchema, error) {
// 	var result models.GatewaysSchema

// 	// check if user has access to create organisation under partner
// 	if partnerUserScope.UserRole != "admin" && partnerUserScope.UserRole != "super_admin" {

// 		return result, errors.New("user is not authorized to create organisation under this partner")
// 	}
// 	// check if partner can create organisation
// 	integaratorServable := checkIfPartnerCanServeIntegrator(partnerUserScope, reqBody.IntegratorID)
// 	if !integaratorServable {

// 		return result, errors.New("partner is not authorized to serve integrator")
// 	}
// 	// Create organisation in infrastructure managment
// 	org := models.OrganisationsSchema{}
// 	org.Name = reqBody.Name
// 	org.Type = reqBody.Type
// 	org.Country = reqBody.Country
// 	org.PartnerID = partnerUserScope.PartnerID
// 	org.IntegratorID = reqBody.IntegratorID
// 	org.AdminName = reqBody.Admin.Name
// 	org.AdminMobile = reqBody.Admin.PhoneNumber
// 	org.AdminEmail = reqBody.Admin.Email
// 	org.AdminEditable = true
// 	org.ExtServicesState = "create_pending"
// 	orgID, err := models.CreateOrganisation(svc.tx, org)
// 	if err != nil {

// 		return result, errors.New("error occured while creating organisation " + err.Error())
// 	}
// 	result.OrgID = orgID
// 	orgSettings := models.OrganisationSettingsSchema{}
// 	orgSettings.OrgId = orgID
// 	orgSettings.PartnerId = partnerUserScope.PartnerID
// 	orgSettings.Integration = reqBody.Integration.Enabled
// 	orgSettings.CardType = "secure"
// 	orgSettings.CardReadBaseKey = uuid.New() // creates a new cardReadBaseKey. This key is used to encrypt card data in production.

// 	log.Println("v2CreateOrganisationsUnderPartner:  orgSettings: ", orgSettings)

// 	err = models.CreateOrganisationSettings(svc.tx, orgSettings)
// 	if err != nil {

// 		log.Println("v2CreateOrganisationsUnderPartner: ", err.Error())

// 		return result, errors.New("error occured while creating organisation settings " + err.Error())
// 	}
// 	err = models.CreateOrganisationModuleType(svc.tx, orgID, reqBody.Modules.ModulesToAdd)
// 	if err != nil {
// 		log.Println("v2CreateOrganisationsUnderPartner: ", err.Error())

// 		return result, errors.New("error occured while creating organisation module type " + err.Error())
// 	}
// 	modulesToAdd, err := models.GetModuleNameByIds(svc.tx, reqBody.Modules.ModulesToAdd)
// 	if err != nil {

// 		log.Println("v2CreateOrganisationsUnderPartner: ", err.Error())

// 		return result, errors.New("error occured while getting module name by ids " + err.Error())
// 	}
// 	modulesToRemove := make([]string, 0)
// 	// Create organisation in external services
// 	statusInfo := serviceResourceStatusInfo{
// 		RequestInfo:       getRequestInfo(ctx.Request),
// 		ResourceName:      "organisation",
// 		ResourceOperation: "create",
// 		ResourceID:        orgID,
// 	}

// 	CreateOrg := webhooks.CreateOrganisationData{}
// 	CreateOrg.Name = reqBody.Name
// 	CreateOrg.Type = reqBody.Type
// 	CreateOrg.Country = reqBody.Country
// 	CreateOrg.OrgID = orgID
// 	CreateOrg.PartnerID = partnerUserScope.PartnerID
// 	CreateOrg.IntegratorID = reqBody.IntegratorID
// 	CreateOrg.Admin.Email = reqBody.Admin.Email
// 	CreateOrg.Admin.Name = reqBody.Admin.Name
// 	CreateOrg.Admin.Phone = reqBody.Admin.PhoneNumber
// 	//Create organisation in acaas services
// 	acaas := webhooks.NewACaaS(os.Getenv("ACAAS_WEBHOOK_URL"), os.Getenv("ACAAS_WEBHOOK_API_KEY"))
// 	statusCode, responseBody, err := acaas.CreateOrganisation(CreateOrg)
// 	if err != nil {

// 		// log error status to database
// 		statusInfo.ServiceName = "acaas"
// 		statusInfo.ServiceStatus = "failed"
// 		statusInfo.ServiceResponseStatusCode = statusCode
// 		statusInfo.ServiceResponseBody = string(responseBody)
// 		statusInfo.Error = err
// 		logServiceResourceStatus(svc.tx, statusInfo)

// 		log.Println("v2CreateOrganisationsUnderPartner: ", err.Error())
// 		err = models.SetOrganisationExtServisesState(svc.tx, orgID, "create_failed")
// 		if err != nil {
// 			log.Println("v2CreateOrganisationsUnderPartner: ", err.Error())

// 			return result, errors.New("error occured while setting organisation ext services state " + err.Error())
// 		}

// 		return result, errors.New("error occured while creating organisation in acaas " + err.Error())
// 	} else {

// 		statusInfo.ServiceName = "acaas"
// 		statusInfo.ServiceStatus = "success"
// 		statusInfo.Error = err
// 		logServiceResourceStatus(svc.tx, statusInfo)
// 	}

// 	//
// 	return result, nil

// }
