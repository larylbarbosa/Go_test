// TODO: test suites and code stay in parallel (_test for every corresponding file)!
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yashdiniz/soa_infrastructure_management/service"
)

type server struct {
	router *gin.Engine
}

func new(router *gin.Engine, svc service.Service) *server {
	s := &server{router}
	s.routes(svc)
	return s
}

func (s *server) listen() {
	s.router.Run(":" + os.Getenv("PORT"))
}

func main() {

	// main() abstraction to better understand error handling.
	if err := run(); err != nil {
		log.Fatalf("%s\n", err) // will make a Fatal log and exit.
	}

}

func run() error {
	err := LoadConfig()
	if err != nil {
		log.Fatal("Failed to load environment variables...")
	}

	dbCred := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=enable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	svc := service.New(
		service.InitDB(dbCred),
	)
	svc.GetCognitoConfigs()
	svc.StartPartnerManagementServer()
	svc.GetInfrastructureMngmtAPIKey()

	server := new(
		gin.Default(),
		svc,
	)
	server.listen()
	return nil
}
