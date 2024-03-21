package gamemodel

import "github.com/1001bit/OnlineCanvasGames/internal/database"

type Game struct {
	ID    string
	Title string
}

func All() ([]Game, error) {
	rows, err := database.DB.Query("SELECT id, title FROM games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []Game

	for rows.Next() {
		var game Game

		err := rows.Scan(&game.ID, &game.Title)
		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}
