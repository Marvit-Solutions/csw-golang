package main

import (
	"csw-golang/cmd/app"
	"csw-golang/library/config"
	"csw-golang/library/config/database"
	"csw-golang/library/struct/request"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var env config.Config

func startApp() {
	SQLMasterConn, err := database.InitDBSQL(env, "postgresql")
	if err != nil {
		fmt.Println(err)
	}

	//start gin engine
	engine := gin.New()

	//add route endpoint and healthcheck
	healthCheck(engine)

	//call route per modul
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
		return
	})

	// Healthcheck endpoint
	engine.GET("ping",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong!")
			return
		})
}
