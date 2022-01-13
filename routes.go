package main

import (
	"github.com/yashdiniz/soa_infrastructure_management/endpoint"
	"github.com/yashdiniz/soa_infrastructure_management/service"
)

func (s *server) routes(svc service.Service) {
	// Key Point: to add extra functionality like metrics and logging
	// Wrap the handlers into Middleware!
	// Do not modify the service/endpoints unless it is involving a core service.
	endpt := endpoint.New( // TODO: should this be here?
		svc,
	)

	// https://stackoverflow.com/questions/59147492/gin-error-panic-wildcard-route-conflicts-with-existing-children
	v2 := s.router.Group("/infrastructureManagement/v2/")

	v2DeleteIntegrationType := s.router.Group("/infrastructureManagement/v2/integrationTypes")

	// Adding the health API
	health := v2.Group("/health")
	health.GET("", endpt.Health())

	//Getting all the organisations under a partner
	organisations := v2.Group("/partners/:partnerId/organisations").Use(endpt.Authentication())
	organisations.GET("", endpt.GetAllPartnerOrganisations())

	integrationType := v2DeleteIntegrationType.Group("").Use(endpt.AuthenticationWithApiKey())
	integrationType.POST("", endpt.CreateIntegrationType())
	integrationType.DELETE("/:integrationTypeId", endpt.DeleteIntegrationType())

	// organisations.POST("", endpt.GetAllPartnerOrganisations())

	// //Accessor Api
	// accessor := v2.Group("/accessor")
	// accessor.DELETE("/:accessorId", endpt.DeleteAccessor())
	// //clients Api
	// clients := v2.Group("/clients")
	// clients.Use(endpt.CheckXApi())
	// clients.POST("", endpt.CreateClient())
	// clients.GET("/:integratorId", endpt.GetClients())
	// clients.DELETE("/:clientId", endpt.DeleteClient())
}
