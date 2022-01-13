package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yashdiniz/soa_infrastructure_management/service"
)

func (endpt *Endpoint) CreateIntegrationType() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var input service.CreateIntegrationTypesRequest
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorOutput{
				"error",
				400,
				err.Error(),
			})
			return

		}

		//Create the organisation under the partner
		err := endpt.svc.CreateIntegrationType(input)

		if err != nil {
			log.Println("v2GetOrganisationsUnderPartner, errorMessage: " + err.Error())

			ctx.AbortWithStatusJSON(http.StatusProxyAuthRequired, gin.H{
				"type":         "error",
				"errorMessage": err.Error(),
			})
			return
		} else {

			ctx.JSON(http.StatusOK, gin.H{
				"type":    "success",
				"message": "success",
			})
			return
		}
	}
}
func (endpt *Endpoint) DeleteIntegrationType() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req := ctx.Params
		integrationTypeID, _ := req.Get("integrationTypeId")
		integrationTypeId, _ := strconv.Atoi(integrationTypeID)
		//Create the organisation under the partner
		err := endpt.svc.DeleteIntegrationType(integrationTypeId)

		if err != nil {
			log.Println("DeleteIntegrationType, errorMessage: " + err.Error())
			json.Unmarshal([]byte(err.Error()), &result)
			ctx.AbortWithStatusJSON(http.StatusProxyAuthRequired, gin.H{
				"type":         "error",
				"errorMessage": result,
			})
			return
		} else {

			ctx.JSON(http.StatusOK, gin.H{
				"type":    "success",
				"message": "Integration type deleted succesfully!!!",
			})
			return
		}
	}
}
