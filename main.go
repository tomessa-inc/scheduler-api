package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	c "scheduler-api/config"
	"scheduler-api/db"
	r "scheduler-api/routes"
	"strings"
)

var echoLambda *echoadapter.EchoLambda

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.LambdaFunctionURLRequest) (Response, error) {

	return Response{Body: "It works!", StatusCode: 200}, nil
}

func Handler3(ctx context.Context, request events.LambdaFunctionURLRequest) (Response, error) {

	return Response{Body: "It works!", StatusCode: 200}, nil
}

func Handler2(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
	///	return Response{Body: "It works!", StatusCode: 200}, nil
}
func Handler1(ctx context.Context, requests events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return echoLambda.ProxyWithContext(ctx, requests)
	///	return Response{Body: "It works!", StatusCode: 200}, nil
}

type Server struct {
}

func main() {
	c.EnvSetup()

	//	enableCors(&w)
	db.Connect()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	//	echoLambda = echoadapter.New(e)
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
	isLambda := os.Getenv("LAMBDA")

	//server.Mount(e)

	if isLambda == "TRUE" {
		//	lambdaAdapter := &LambdaAdapter{Echo: e}
		//	lambda.Start(lambdaAdapter.Handler)
		//		lambdaAdapter := &LambdaAdapter{Echo: e}
		//		lambda.Start(lambdaAdapter.Handler)
		server := wrapRouter(e)

		lambda.Start(server)
	} else {
		e.Logger.Fatal(e.Start(":3600"))
	}
	//e.Start()
	//e.Logger.Fatal(e.Start(":3500"))
	//lambda.Start(Handler)
}

func wrapRouter(e *echo.Echo) func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		body := strings.NewReader(request.Body)
		fmt.Println("body")

		fmt.Println(body)

		req := httptest.NewRequest(request.HTTPMethod, request.Path, body)
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}

		fmt.Println("req")

		fmt.Println(req)
		fmt.Println("req2")
		fmt.Println("req3")
		goo := "ffff"
		fmt.Println(goo)
		q := req.URL.Query()
		for k, v := range request.QueryStringParameters {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		res := rec.Result()
		responseBody, err := io.ReadAll(res.Body)
		fmt.Println("responseBody")
		fmt.Println(responseBody)
		if err != nil {
			return formatAPIErrorResponse(http.StatusInternalServerError, res.Header, err.Error())
		}

		return formatAPIResponse(res.StatusCode, res.Header, string(responseBody))
	}
}

func formatAPIResponse(statusCode int, headers http.Header, responseData string) (events.APIGatewayProxyResponse, error) {
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"

	return events.APIGatewayProxyResponse{
		Body:       responseData,
		Headers:    responseHeaders,
		StatusCode: statusCode,
	}, nil
}

func formatAPIErrorResponse(statusCode int, headers http.Header, err string) (events.APIGatewayProxyResponse, error) {
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	for key, value := range headers {
		responseHeaders[key] = ""

		if len(value) > 0 {
			responseHeaders[key] = value[0]
		}
	}

	responseHeaders["Access-Control-Allow-Origin"] = "*"
	responseHeaders["Access-Control-Allow-Headers"] = "origin,Accept,Authorization,Content-Type"

	return events.APIGatewayProxyResponse{
		Body:       err,
		Headers:    responseHeaders,
		StatusCode: statusCode,
	}, nil
}
