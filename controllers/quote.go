package controllers

import (
	"math/rand"
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
	"github.com/jeyren95/big-bang-theory-quotes/models"
	"github.com/jeyren95/big-bang-theory-quotes/queries"
)

// GetQuotes go doc
//
//	@Summary		Get quotes
//	@Description	get quotes by character, season or episode
//	@Router			/quotes [get]
//	@Produce		json
//	@Success		200			{object}	models.Quote
//	@Failure		500			{object}	models.ErrorResponse
//	@Param			character	query		string	false	"search quotes by character"	example("sheldon")
//	@Param			season		query		string	false	"search quotes by season"		example(1)
//	@Param			episode		query		string	false	"search quotes by episode"		example(1)
func GetQuotes(ctx *gin.Context) {
	character, characterMatch := ctx.GetQuery("character")
	season, seasonMatch := ctx.GetQuery("season")
	episode, episodeMatch := ctx.GetQuery("episode")

	titleConverter := cases.Title(language.English)
	character = titleConverter.String(character)

	if !seasonMatch && episodeMatch {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters"})
		return
	}

	var quotes []*models.Quote
	var err error

	if characterMatch && seasonMatch && episodeMatch {
		quotes, err = queries.GetQuotesByAllParams(character, season, episode)
	} else if characterMatch && seasonMatch {
		quotes, err = queries.GetQuotesByCharacterAndSeason(character, season)
	} else if seasonMatch && episodeMatch {
		quotes, err = queries.GetQuotesBySeasonAndEpisode(season, episode)
	} else if characterMatch {
		quotes, err = queries.GetQuotesByCharacter(character)
	} else if seasonMatch {
		quotes, err = queries.GetQuotesBySeason(season)
	} else {
		quotes, err = queries.GetAllQuotes()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"quotes": quotes})
}

// GetRandomQuote go doc
//
//	@Summary	Get a random quote
//	@Router		/quotes/random [get]
//	@Produce	json
//	@Success	200	{object}	models.Quote
//	@Failure	500	{object}	models.ErrorResponse
func GetRandomQuote(ctx *gin.Context) {
	count, err := queries.GetQuoteCount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	selected := rand.Intn(count)
	quote, err := queries.GetQuoteById(selected)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, quote)
}
