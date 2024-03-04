package main

import (
	api "api"
	"os"
	db "sqlc"
	util "util"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// "github.com/golang-migrate/migrate/v4"
// _ "github.com/golang-migrate/migrate/v4/database/postgres"
// _ "github.com/golang-migrate/migrate/v4/source/file"

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	// fmt.Println("Templates:", config.Templates)

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	postgres, err := db.InitialiseDB(config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	queries := db.New(postgres)
	// runDBMigration(config.MigrationURL, config.DBSource)

	newHandler := db.NewHandlers(queries)

	server, err := api.CreateServer(config, newHandler)
	if err != nil {
		log.Fatal().Err(err)
	}

	httpServer := server.SetupServer()

	api.StartServer(httpServer)
}
