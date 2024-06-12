package main

import (
	"context"
	"github.com/rachitkawar/boilerplate-go/src/internal/server"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	"go.opencensus.io/trace"
	"os/signal"
	"syscall"
	"time"
)

import "github.com/rachitkawar/boilerplate-go/src/internal/database"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	start := time.Now()

	utils.InitialiseLogger()

	utils.Log.Info("started logger...")

	utils.Log.Info("loading env..")
	utils.InitConfig()
	utils.Log.Info("env loaded")

	utils.Log.Info("connecting db..")

	database.InitDB(ctx)

	apiServer := server.InitializeServer()

	go apiServer.Run(":8080")

	utils.Log.Info("time to load the app: ", time.Now().Sub(start), trace.StatusCodeUnknown)

	select {
	case <-ctx.Done():
		utils.Log.Info("shutting down everything")

		err := apiServer.Shutdown(ctx)
		if err != nil {
			utils.Log.Error("error shutting api server ", err)
		}

		utils.Log.Info("api server closed")

	}

}
