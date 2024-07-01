package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	c "scheduler-api/config"
	"scheduler-api/db"
	r "scheduler-api/routes"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.LambdaFunctionURLRequest) (Response, error) {

	return Response{Body: "It works!", StatusCode: 200}, nil
}

func main() {
	c.EnvSetup()

	//	enableCors(&w)
	db.Connect()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	//e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},//
	//github.com/labstack/echo/v4/middleware	}))

	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	//		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//	}))
	r.InitRoutes(e)
	//	isLambda := os.Getenv("LAMBDA")

	if "TRUE" == "TRUE" {
		//		lambdaAdapter := &LambdaAdapter{Echo: e}
		//		lambda.Start(lambdaAdapter.Handler)
		lambda.Start(Handler)
	} else {
		e.Logger.Fatal(e.Start(":3600"))
	}
	//e.Logger.Fatal(e.Start(":3500"))
	//lambda.Start(Handler)
}
