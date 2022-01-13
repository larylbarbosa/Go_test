package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var result map[string]interface{}

func (endpt *Endpoint) GetAllPartnerOrganisations() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqParams := ctx.Params
		partnerID, _ := reqParams.Get("partnerId")

		// get partnerId from url
		partnerId, _ := strconv.Atoi(partnerID) //convert string to int64

		//Get the organisations under the partner
		results, err := endpt.svc.GetOrganisationsUnderPartner(int(partnerId))

		if err != nil {
			log.Println("GetAllPartnerOrganisations, errorMessage: " + err.Error())
			json.Unmarshal([]byte(ErrRespDatabaseError), &result)
			ctx.AbortWithStatusJSON(http.StatusProxyAuthRequired, result)
			return
		} else {

			ctx.JSON(http.StatusOK, gin.H{
				"type": "success",
				"message": gin.H{
					"organisations": results,
				},
			})
			return
		}
	}
}

//CreateOrganisation creates a new organisation under a partner

// func (endpt *Endpoint) CreatePartnerOrganisations() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		req := ctx.Params
// 		partnerID, _ := req.Get("partnerId")                // get partnerId from url
// 		partnerId, _ := strconv.ParseInt(partnerID, 10, 64) //convert string to int64

// 		var input service.CreateOrganisationUnderPartnerRequest
// 		if err := ctx.ShouldBindJSON(&input); err != nil {
// 			ctx.JSON(http.StatusBadRequest, ErrorOutput{
// 				"error",
// 				400,
// 				err.Error(),
// 			})
// 			return

// 		}

// 		//Create the organisation under the partner
// 		result, err := endpt.svc.GetOrganisationsUnderPartner(int(partnerId))

// 		if err != nil {
// 			log.Println("v2GetOrganisationsUnderPartner, errorMessage: " + err.Error())
// 			ctx.JSON(http.StatusProxyAuthRequired, gin.H{
// 				"type":         "error",
// 				"errorMessage": "Error getting organisations under partner ,errorMessage: " + err.Error(),
// 			})
// 			return
// 		} else {

// 			ctx.JSON(http.StatusOK, gin.H{
// 				"MsgType": "success",
// 				"data":    result,
// 			})
// 			return
// 		}
// 	}
// }
