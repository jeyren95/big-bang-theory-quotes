package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeyren95/big-bang-theory-quotes/queries"
)

func GetAllQuotes(ctx *gin.Context) {
	allQuotes, err := queries.GetAllQuotes()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"quotes": allQuotes})
}
