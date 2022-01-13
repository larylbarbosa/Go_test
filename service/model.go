package service

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

/*
	Initializes the database
*/
func InitDB(dbCred string) *sql.DB {

	dsn := dbCred
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// database.Migrator().AutoMigrate(IdpSubjects{}, IdpKeys{}, IdentityScopes{}, IdentityRefreshTokens{}, IdentityCredentials{}, Clients{})

	return database
}
