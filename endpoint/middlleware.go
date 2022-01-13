package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (endpt *Endpoint) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqParams := ctx.Params
		partnerID, _ := reqParams.Get("partnerId")

		// get partnerId from url
		partnerId, err := strconv.Atoi(partnerID) //convert string to int64
		if partnerId == 0 || err != nil {
			log.Println(partnerId)
			json.Unmarshal([]byte(ErrRequiredMethod), &result)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
			return
		}

		//TODO abstract into middleware

		//Athenticate the user
		scopes, err := endpt.svc.AuthenticateUser(ctx)
		if err != nil {
			log.Println("GetAllPartnerOrganisations, Error authenticating user: " + err.Error())
			json.Unmarshal([]byte(ErrRespAPIAuthenticationFailure), &result)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, result)
			return
		}
		//Check if the user has the right scope
		partnerUserScope, partnerExist := endpt.svc.CheckIfPartnerExistsInScope(scopes, partnerId)
		if !partnerExist {
			json.Unmarshal([]byte(rrRespAuthorisationFailure), &result)
			ctx.AbortWithStatusJSON(http.StatusProxyAuthRequired, result)
			return
		}
		if ctx.Request.Method == "POST" {

			if partnerUserScope.UserRole != "admin" && partnerUserScope.UserRole != "super_admin" {
				json.Unmarshal([]byte(ErrRespAuthorisationFailure), &result)
				ctx.AbortWithStatusJSON(http.StatusProxyAuthRequired, result)
				return
			}
		}

	}
}
func (endpt *Endpoint) AuthenticationWithApiKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Athenticate the user
		err := endpt.svc.AuthenticateWithApiKey(ctx)
		if err != nil {
			log.Println("GetAllPartnerOrganisations, Error authenticating user: " + err.Error())
			json.Unmarshal([]byte(ErrRespAPIAuthenticationFailure), &result)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, result)
			return
		}

	}
}
