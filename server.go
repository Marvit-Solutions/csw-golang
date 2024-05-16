package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/route"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/config/database"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
	"github.com/Marvit-Solutions/csw-golang/library/struct/request"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

var env config.Config

// startApp initializes the application.
func startApp() {
	// Initialize authentication middleware.
	err := auth.NewMiddlewareConfig(env)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize PostgreSQL database.
	SQLMasterConn, err := database.InitDBSQL(env, "postgresql")
	if err != nil {
		fmt.Println(err)
	}

	// Initialize Redis database.
	redisConn, err := database.InitDBRedis(env)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize MongoDB database.
	mongoConn, err := database.InitDBMongo(env)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize Gin engine.
	engine := gin.New()

	// Add Swagger documentation.
	swagger(engine)

	// Add route endpoints and healthcheck.
	healthCheck(engine)

	// Call route initialization.
	req := request.RouteInit{Engine: engine, SQLMaster: SQLMasterConn, Redis: redisConn, Mongo: mongoConn, Env: env}
	route.NewRouteInit(req)

	// Run server.
	serverPort := env.GetString("server.address")
	log.Println("run on port ", serverPort)
	fmt.Printf("App running on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, engine); err != nil {
		log.Fatal(err)
	}
}

// healthCheck sets up root and health check endpoints.
func healthCheck(engine *gin.Engine) {
	// Root endpoint.
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "CSW API Service")
	})

	// Healthcheck endpoint.
	engine.GET("ping",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong!")
		})
}

// swagger sets up Swagger documentation endpoint.
func swagger(engine *gin.Engine) {
	if env.GetString("server.env") == "stage" {
		docs.SwaggerInfo.Host = "stage.api.csw.id"
	} else if env.GetString("server.env") == "local" {
		docs.SwaggerInfo.Host = "127.0.0.1:8080"
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
