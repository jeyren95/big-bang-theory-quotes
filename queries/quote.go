package queries

import (
	"database/sql"
	"fmt"

	"github.com/jeyren95/big-bang-theory-quotes/db"
	"github.com/jeyren95/big-bang-theory-quotes/models"
)

func scanRows(rows *sql.Rows) ([]*models.Quote, error) {
	var quotes []*models.Quote

	for rows.Next() {
		quote := new(models.Quote)
		err := rows.Scan(&quote.Id, &quote.Quote, &quote.Character, &quote.Season, &quote.Episode, &quote.Title)

		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func InsertQuote(quote models.Quote) error {
	statement, err := db.DB.Prepare("INSERT INTO quote(quote, character, season, episode, title) VALUES($1, $2, $3, $4, $5)")

	if err != nil {
		fmt.Print(err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(quote.Quote, quote.Character, quote.Season, quote.Episode, quote.Title)

	if err != nil {
		return err
	}

	return nil
}

func GetAllQuotes() ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote")

	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetQuotesByCharacter(character string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE character = $1")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(character)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetQuotesBySeason(season string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE season = $1")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(season)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetQuotesBySeasonAndEpisode(season string, episode string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE season = $1 AND episode = $2")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(season, episode)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetQuotesByCharacterAndSeason(character string, season string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE character = $1 AND season = $2")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(character, season)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetQuotesByAllParams(character string, season string, episode string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE character = $1 AND season = $2 AND episode = $3")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(character, season, episode)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		return nil, err
	}
	return quotes, nil
}
