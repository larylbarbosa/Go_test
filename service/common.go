package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yashdiniz/soa_infrastructure_management/models"
)

type CustomUserScopes []PartnerUserScope

type PartnerUserScope struct {
	PartnerID          int    `json:"pi"`
	UserRole           string `json:"rn"`
	PartnerIntegrators []int  `json:"ii"`
}
type serviceResourceStatusInfo struct {
	RequestInfo               string    `json:"requestInfo"`
	ResourceName              string    `json:"resourceName"`
	ResourceOperation         string    `json:"resourceOperation"`
	ResourceID                int       `json:"resourceID"`
	ServiceName               string    `json:"serviceName"`
	ServiceStatus             string    `json:"serviceStatus"`
	ServiceResponseStatusCode int       `json:"serviceResponseStatusCode"`
	ServiceResponseBody       string    `json:"serviceResponceBody"`
	TransactionID             uuid.UUID `json:"transactionID"`
	Error                     error     `json:"-"`
	ErrorString               string    `json:"error"`
}

var region string
var userPoolID string
var infrastructureMngmtAPIKey string

func (svc service) GetCognitoConfigs() {

	region = os.Getenv("COGNITO_REGION")
	userPoolID = os.Getenv("COGNITO_USER_POOL_ID")

}

func (svc service) GetInfrastructureMngmtAPIKey() {

	infrastructureMngmtAPIKey = os.Getenv("INFRASTRUCTURE_MANAGEMENT_API_KEY")

	log.Println("AuthenticateAPI, INFRASTRUCTURE_MANAGEMENT_API_KEY: ", infrastructureMngmtAPIKey)

}

func (svc service) AuthenticateUser(r *gin.Context) (scopes CustomUserScopes, err error) {

	if len(r.Request.Header["Authorization"]) == 0 {
		return scopes, errors.New("AuthenticateUser, Authorization header not provided")
	}

	token, err := auth.ParseJWT(r.Request.Header["Authorization"][0])

	if err != nil {
		return scopes, errors.New("AuthenticateUser, Error: " + err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return scopes, errors.New("AuthenticateUser, Error: " + err.Error())
	}

	if !token.Valid {
		return scopes, errors.New("AuthenticateUser, Error: " + err.Error())
	}

	userScopeString := fmt.Sprintf("%s", claims["custom:userScopes"])
	fmt.Printf("\nAuthenticateUser, User Scopes String : \n%s\n\n", userScopeString)

	err = json.Unmarshal([]byte(userScopeString), &scopes)
	if err != nil {
		return scopes, errors.New("AuthenticateUser, Error: " + err.Error())
	}

	for i, r := range scopes {
		fmt.Printf("%d, PartnerID: %02d, UserRole: %s, PartnerIntegrators: %v\n", i, r.PartnerID, r.UserRole, r.PartnerIntegrators)
	}

	return scopes, err
}

func (svc service) CheckIfPartnerExistsInScope(scopes CustomUserScopes, partnerID int) (partnerUserScope PartnerUserScope, partnerExist bool) {

	partnerExist = false

	for _, s := range scopes {

		if s.PartnerID == partnerID {
			partnerUserScope = s
			partnerExist = true
			break
		}
	}

	return partnerUserScope, partnerExist
}

func (svc service) StartPartnerManagementServer() {

	var err error
	auth, err = NewAuth(&Config{
		CognitoRegion:     region,
		CognitoUserPoolID: userPoolID,
	})

	if err != nil {
		log.Println("Failed to create authorizer")
	}

}
func checkIfPartnerCanServeIntegrator(partnerUserScope PartnerUserScope, integaratorID int) (integaratorServable bool) {

	integaratorServable = false

	for _, i := range partnerUserScope.PartnerIntegrators {

		if i == integaratorID {
			integaratorServable = true
			break
		}
	}

	return integaratorServable
}

func getRequestInfo(r *http.Request) string {

	rd, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("getRequestInfo, " + err.Error())
	}

	return string(rd)
}

func logServiceResourceStatus(tx *sql.Tx, errorInfo serviceResourceStatusInfo) {

	var logID int

	if errorInfo.ServiceStatus != "success" {

		if errorInfo.Error != nil {

			errorInfo.ErrorString = errorInfo.Error.Error()
		}

		jbs, err := json.Marshal(errorInfo)
		if err != nil {
			log.Println("logServiceResourceStatus, Error occurred while marshalling: " + err.Error())
			return
		}

		errorLog := models.ErrorLogsSchema{
			Message: string(jbs),
		}

		log.Printf("\n\n logServiceResourceStatus, errorLog Message: %s \n\n", string(jbs))

		logID, err = models.CreateErrorLog(tx, errorLog)
		if err != nil {
			log.Println("logServiceResourceStatus, Error occurred while creating error log: " + err.Error())
			return
		}
	} else {
		logID = 0
	}

	statusLog := models.ServiceResourceStatusSchema{
		ResourceName:      errorInfo.ResourceName,
		ResourceOperation: errorInfo.ResourceOperation,
		ResourceID:        errorInfo.ResourceID,
		ServiceName:       errorInfo.ServiceName,
		ServiceStatus:     errorInfo.ServiceStatus,
		ErrorLogID:        logID,
		TransactionID:     errorInfo.TransactionID,
	}

	err := models.CreateServiceResourceStatusLog(tx, statusLog)
	if err != nil {
		log.Println("logServiceResourceStatus, Error occurred while creating service resource status log: " + err.Error())
		return
	}
}
func (svc service) AuthenticateWithApiKey(ctx *gin.Context) (err error) {

	if len(ctx.Request.Header["Authorization"]) == 0 {
		return errors.New("AuthenticateAPI, Authorization header not provided")
	}

	log.Println(ctx.Request.Header["Authorization"])

	apiKeyProvided := ctx.Request.Header["Authorization"][0]

	log.Println("AuthenticateAPI, apiKeyProvided: ", apiKeyProvided)

	if len(infrastructureMngmtAPIKey) == 0 {

		return errors.New("AuthenticateAPI API key not added as env variable")
	}

	if apiKeyProvided != infrastructureMngmtAPIKey {

		return errors.New("AuthenticateAPI, API key provided doesn't match")

	} else {
		err = nil
	}

	return err
}
