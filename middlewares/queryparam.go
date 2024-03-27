package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func CheckQueryParamsCount(desiredCount int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		queryParams := ctx.Request.URL.Query()

		if len(queryParams) == 0 {
			ctx.Next()
		}

		if len(queryParams) > desiredCount {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func MatchQueryParamNames(desiredNames []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		queryParams := ctx.Request.URL.Query()

		for p := range queryParams {
			if !slices.Contains(desiredNames, p) {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters"})
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
