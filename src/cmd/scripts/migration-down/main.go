package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"go.opencensus.io/trace"
	"log"
	"time"
)

func main() {

	start := time.Now()

	utils.InitialiseLogger()

	utils.Log.Info("started logger...")

	utils.Log.Info("loading env..")
	utils.InitConfig()
	utils.Log.Info("env loaded")

	m, err := migrate.New(
		"file://"+utils.GetEnv("MIGRATION_PATH"),
		utils.GetEnv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

	utils.Log.Info("time to load the migrations: ", time.Now().Sub(start), trace.StatusCodeUnknown)
	utils.Log.Info("shutting down migrations")
}
