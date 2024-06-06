package main

import (
	"context"
	"github.com/rachitkawar/boilerplate-go/src/common"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain/jwt"
	"github.com/rachitkawar/boilerplate-go/src/internal/server"
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

	common.InitialiseLogger()

	common.Log.Info("started logger...")

	common.Log.Info("loading env..")
	common.InitConfig()
	common.Log.Info("env loaded")

	common.Log.Info("connecting db..")

	database.InitDB(ctx)

	jwtDomain := &jwt.TokenMaster{SecretKey: "f"}

	apiServer := server.InitializeServer(jwtDomain)

	go apiServer.Run(":8080")

	common.Log.Info("time to load the app: ", time.Now().Sub(start), trace.StatusCodeUnknown)

	select {
	case <-ctx.Done():
		common.Log.Info("shutting down everything")

		err := apiServer.Shutdown(ctx)
		if err != nil {
			common.Log.Error("error shutting api server ", err)
		}

		common.Log.Info("api server closed")

	}
}
