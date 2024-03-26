package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jeyren95/big-bang-theory-quotes/models"
	"github.com/jeyren95/big-bang-theory-quotes/queries"
)

func Seed() {
	data, err := os.ReadFile("quotes.json")

	if err != nil {
		panic("Failed to read JSON file")
	}

	var quotes []models.Quote
	if err = json.Unmarshal(data, &quotes); err != nil {
		panic("Failed to parse JSON data")
	}

	for i := 0; i < len(quotes); i++ {
		queries.InsertQuote(quotes[i])
	}

	fmt.Print("Successfully migrated seed data\n")
}
