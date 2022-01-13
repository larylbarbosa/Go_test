package service

type CreateOrganisationUnderPartnerRequest struct {
	IntegratorID int           `json:"integratorId" binding:"required"`
	Name         string        `json:"name" binding:"required"`
	Type         string        `json:"type" binding:"required"`
	Country      string        `json:"country" binding:"required"`
	Admin        adminDetails  `json:"adminDetails" binding:"required"`
	Site         []siteDetails `json:"site" binding:"required"`
	Integration  integration   `json:"integration" binding:"required"`
	Modules      modules       `json:"modules" binding:"required"`
}
type adminDetails struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Editable    bool   `json:"editable" binding:"required"`
}
type siteDetails struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}
type integration struct {
	Enabled             bool  `json:"enabled" binding:"required"`
	IntegrationToAdd    []int `json:"integrationToAdd"`
	IntegrationToRemove []int `json:"integrationToRemove" `
}
type modules struct {
	ModulesToAdd    []int `json:"modulesToAdd"`
	ModulesToRemove []int `json:"modulesToRemove" `
}
type CreateIntegrationTypesRequest struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
