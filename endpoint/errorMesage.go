package endpoint

import "fmt"

const (
	ErrRequiredMethod               = `{"type":"error", "message":{"errorCode": 1292, "errorMessage":"partnerId is required."}}`
	ErrIsNumberMethod               = `{"type":"error", "message":{"errorCode": 1293, "errorMessage":"partnerId must be a number"}}`
	ErrRespAPIAuthenticationFailure = `{"type":"error", "message":{"errorCode": 1008, "errorMessage":"API authentication failed."}}`
	rrRespAuthorisationFailure      = `{"type":"error", "message":{"errorCode": 1007, "errorMessage":"User not authorised to perform operation."}}`
	ErrRespDatabaseError            = `{"type":"error", "message":{"errorCode": 1005, "errorMessage":"Server database error."}}`
	ErrRespAuthorisationFailure     = `{"type":"error", "message":{"errorCode": 1007, "errorMessage":"User not authorised to perform operation."}}`
)

func BuildErrRespCannotDeleteIntegrationType(orgId int) string {

	return fmt.Sprintf(`{"type":"error", "message":{ "errorCode":1236, "errorMessage":"Failed to delete integration type. It is enabled for organisation: %d "}}`, orgId)
}
