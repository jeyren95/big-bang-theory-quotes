package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jeyren95/big-bang-theory-quotes/controllers"
	"github.com/jeyren95/big-bang-theory-quotes/db"
	_ "github.com/jeyren95/big-bang-theory-quotes/docs"
	"github.com/jeyren95/big-bang-theory-quotes/middlewares"
	"github.com/jeyren95/big-bang-theory-quotes/seed"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Big Bang Theory Quotes API
// @version		1.0
// @description	API that returns quotes from the sitcom show Big Bang Theory
// @host			localhost:8080
func main() {
	slog.Info("Started server")
	db.DbConnect()
	seed.Seed()

	router := gin.Default()
	router.GET(
		"/quotes",
		middlewares.CheckQueryParamsCount(3),
		middlewares.MatchQueryParamNames([]string{"character", "season", "episode"}),
		controllers.GetQuotes,
	)
	router.GET("/quotes/random", controllers.GetRandomQuote)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ipAddress := os.Getenv("IPADDR")
	router.Run(ipAddress)
}
