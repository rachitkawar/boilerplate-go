package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rachitkawar/boilerplate-go/src/common"
	"go.opencensus.io/trace"
	"log"
	"time"
)

func main() {

	start := time.Now()

	common.InitialiseLogger()

	common.Log.Info("started logger...")

	common.Log.Info("loading env..")
	common.InitConfig()
	common.Log.Info("env loaded")

	m, err := migrate.New(
		"file://"+common.GetEnv("MIGRATION_PATH"),
		common.GetEnv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

	common.Log.Info("time to load the migrations: ", time.Now().Sub(start), trace.StatusCodeUnknown)
	common.Log.Info("shutting down migrations")
}
