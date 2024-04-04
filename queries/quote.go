package queries

import (
	"database/sql"
	"log/slog"

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
		slog.Error("InsertQuote", "Prepare statement error", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(quote.Quote, quote.Character, quote.Season, quote.Episode, quote.Title)
	if err != nil {
		slog.Error("InsertQuote", "Execute statement error", err)
		return err
	}

	return nil
}

func GetQuoteById(id int) (*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE id = $1")
	if err != nil {
		slog.Error("GetQuoteById", "Prepare statement error", err)
		return nil, err
	}

	defer statement.Close()

	quote := new(models.Quote)
	if err = statement.QueryRow(id).Scan(&quote.Id, &quote.Quote, &quote.Character, &quote.Season, &quote.Episode, &quote.Title); err != nil {
		slog.Error("GetQuoteById", "Query row error", err)
		return nil, err
	}

	return quote, err
}

func GetAllQuotes() ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote")
	if err != nil {
		slog.Error("GetAllQuotes", "Prepare statement error", err)
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		slog.Error("GetAllQuotes", "Query error", err)
		return nil, err
	}
	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		slog.Error("GetAllQuotes", "Scan rows error", err)
		return nil, err
	}

	return quotes, nil
}

func GetQuotesByCharacter(character string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE character LIKE CONCAT('%', $1, '%')")
	if err != nil {

		slog.Error("GetQuotesByCharacter", "Prepare statement error", err)
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(character)
	if err != nil {
		slog.Error("GetQuotesByCharacter", "Query error", err)
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		slog.Error("GetQuotesByCharacter", "Scan rows error", err)
		return nil, err
	}

	return quotes, nil
}

func GetQuotesBySeason(season string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE season = $1")
	if err != nil {
		slog.Error("GetQuotesBySeason", "Prepare statement error", err)
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(season)
	if err != nil {
		slog.Error("GetQuotesBySeason", "Query error", err)
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		slog.Error("GetQuotesBySeason", "Scan rows error", err)
		return nil, err
	}

	return quotes, nil
}

func GetQuotesBySeasonAndEpisode(season string, episode string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE season = $1 AND episode = $2")
	if err != nil {
		slog.Error("GetQuotesBySeasonAndEpisode", "Prepare statement error", err)
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(season, episode)
	if err != nil {
		slog.Error("GetQuotesBySeasonAndEpisode", "Query error", err)
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		slog.Error("GetQuotesBySeasonAndEpisode", "Scan rows error", err)
		return nil, err
	}

	return quotes, nil
}

func GetQuotesByCharacterAndSeason(character string, season string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE character LIKE CONCAT('%', $1, '%') AND season = $2")
	if err != nil {
		slog.Error("GetQuotesByCharacterAndSeason", "Prepare statement error", err)
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(character, season)
	if err != nil {
		slog.Error("GetQuotesByCharacterAndSeason", "Query error", err)
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		slog.Error("GetQuotesByCharacterAndSeason", "Scan rows error", err)
		return nil, err
	}

	return quotes, nil
}

func GetQuotesByAllParams(character string, season string, episode string) ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote WHERE character LIKE CONCAT('%', $1, '%') AND season = $2 AND episode = $3")
	if err != nil {
		slog.Error("GetQuotesByAllParams", "Prepare statement error", err)
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(character, season)
	if err != nil {
		slog.Error("GetQuotesByAllParams", "Query error", err)
		return nil, err
	}

	defer rows.Close()

	quotes, err := scanRows(rows)
	if err != nil {
		slog.Error("GetQuotesByAllParams", "Scan rows error", err)
		return nil, err
	}
	return quotes, nil
}

func GetQuoteCount() (int, error) {
	statement, err := db.DB.Prepare("SELECT COUNT(*) FROM quote")
	if err != nil {
		slog.Error("GetQuoteCount", "Prepare statement error", err)
		return -1, err
	}

	defer statement.Close()

	var count int

	if err = statement.QueryRow().Scan(&count); err != nil {
		slog.Error("GetQuoteCount", "Query row error", err)
		return -1, err
	}

	return count, err
}
