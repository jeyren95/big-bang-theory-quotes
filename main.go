package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jeyren95/big-bang-theory-quotes/controllers"
	"github.com/jeyren95/big-bang-theory-quotes/db"
	"github.com/jeyren95/big-bang-theory-quotes/middlewares"
	"github.com/jeyren95/big-bang-theory-quotes/seed"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
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

	ipAddress := os.Getenv("IPADDR")
	router.Run(ipAddress)
}
