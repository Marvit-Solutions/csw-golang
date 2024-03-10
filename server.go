package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/cmd/app"
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

func startApp() {
	// initiate auth middleware
	err := auth.NewMiddlewareConfig(env)
	if err != nil {
		fmt.Println(err)
	}

	// initiate postgresql
	SQLMasterConn, err := database.InitDBSQL(env, "postgresql")
	if err != nil {
		fmt.Println(err)
	}

	//start gin engine
	engine := gin.New()

	//add swagger documentation
	swagger(engine)

	//add route endpoint and healthcheck
	healthCheck(engine)

	//call route
	req := request.RouteInit{Engine: engine, SQLMaster: SQLMasterConn, Env: env}
	app.NewRouteInit(req)

	//run server
	serverPort := env.GetString("server.address")
	log.Println("run on port ", serverPort)
	fmt.Printf("App running on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, engine); err != nil {
		log.Fatal(err)
	}
}

func healthCheck(engine *gin.Engine) {
	// root endpoint
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "CSW API Service")
	})

	// Healthcheck endpoint
	engine.GET("ping",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong!")
		})
}

func swagger(engine *gin.Engine) {
	if env.GetString("server.env") == "stage" {
		docs.SwaggerInfo.Host = "stage.api.csw.id"
	} else if env.GetString("server.env") == "local" {
		docs.SwaggerInfo.Host = "127.0.0.1:8080"
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
