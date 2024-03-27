package controllers

import (
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
	"github.com/jeyren95/big-bang-theory-quotes/models"
	"github.com/jeyren95/big-bang-theory-quotes/queries"
)

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
