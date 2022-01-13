package main

import (
	"github.com/joho/godotenv"
)

/*
	The Config struct is a placeholder, and has no practical importance.
	godotenv automatically sets environment variables,
	so they can be read via os.getEnv(...)
*/
type Config struct {
	GO_ENV string

	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`

	PORT         string `mapstructure:"PORT"`
	BASE_URL     string `mapstructure:"BASE_URL"`
	SERVICE_NAME string `mapstructure:"SERVICE_NAME"`

	SQS_ERROR_EVENT_QUEUE_URL   string `mapstructure:"SQS_ERROR_EVENT_QUEUE_URL"`
	SQS_PARTNER_EVENT_QUEUE_URL string `mapstructure:"SQS_PARTNER_EVENT_QUEUE_URL"`

	IOT_WEBHOOK_URL     string `mapstructure:"IOT_WEBHOOK_URL"`
	IOT_WEBHOOK_API_KEY string `mapstructure:"IOT_WEBHOOK_API_KEY"`

	ACAAS_WEBHOOK_URL     string `mapstructure:"ACAAS_WEBHOOK_URL"`
	ACAAS_WEBHOOK_API_KEY string `mapstructure:"ACAAS_WEBHOOK_API_KEY"`

	DIAGNOSTIC_WEBHOOK_URL     string `mapstructure:"DIAGNOSTIC_WEBHOOK_URL"`
	DIAGNOSTIC_WEBHOOK_API_KEY string `mapstructure:"DIAGNOSTIC_WEBHOOK_API_KEY"`

	HARDWARE_MANAGER_BASE_URL string `mapstructure:"HARDWARE_MANAGER_BASE_URL"`
	HARDWARE_MANAGER_API_KEY  string `mapstructure:"HARDWARE_MANAGER_API_KEY"`

	PARTNER_MANAGEMENT_WEBHOOK_URL     string `mapstructure:"PARTNER_MANAGEMENT_WEBHOOK_URL"`
	PARTNER_MANAGEMENT_WEBHOOK_API_KEY string `mapstructure:"PARTNER_MANAGEMENT_WEBHOOK_API_KEY"`

	ACAAS_SERVICE_HOST   string `mapstructure:"ACAAS_SERVICE_HOST"`
	COGNITO_REGION       string `mapstructure:"COGNITO_REGION"`
	COGNITO_USER_POOL_ID string `mapstructure:"COGNITO_USER_POOL_ID"`

	INFRASTRUCTURE_MANAGEMENT_API_KEY string `mapstructure:"INFRASTRUCTURE_MANAGEMENT_API_KEY"`
}

func LoadConfig() error {
	err := godotenv.Load(".env")
	return err
}
