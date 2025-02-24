package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi/v5"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"

	"github.com/beabys/go-template/internal/app"
	cfg "github.com/beabys/go-template/internal/app/config"
	"github.com/beabys/go-template/pkg/logger"
	"go.uber.org/zap"
)

var chiLambda *chiadapter.ChiLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if chiLambda == nil {

		app := app.New()

		// Setup configurations
		configs := cfg.New()
		err := configs.LoadConfigs()
		if err != nil {
			fmt.Errorf("error setting configs: %v", err)
			panic(err)
		}
		app.Config = configs

		// SetLogger
		loggerConfigs := &logger.DefaultLoggerConfig{}
		logger, err := logger.NewDefaultLogger(loggerConfigs)
		if err != nil {
			fmt.Errorf("error setting logger: %v", err)
			panic(err)
		}
		app.SetLogger(logger)

		h, err := app.SetLambdaMuxHandler()
		if err != nil {
			app.Logger.Fatal("error setting mux handler", zap.Error(err))
			panic(err)
		}
		// casting as chiadapter only accepts chi.Mux
		// the implementation of SetLambdaMuxHandler returns http.Handler
		// even though we are working with chi.Mux
		// we use http.Handler to make it more flexible and to avoid dependency on chi for the tests
		// in case to change to standard library mux, we can change the implementation in this function
		mux := h.(*chi.Mux)
		chiLambda = chiadapter.New(mux)
	}

	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
