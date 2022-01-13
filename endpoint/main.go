package endpoint

import "github.com/yashdiniz/soa_infrastructure_management/service"

type ErrorOutput struct {
	Type         string `json:"type"`
	StatusCode   int    `json:"errorCode" `
	ErrorMessage string `json:"errorMessage" `
}
type Endpoint struct {
	svc service.Service
}

/*
	Creates a New Endpoint
*/
func New(svc service.Service) *Endpoint {
	return &Endpoint{
		svc,
	}
}
