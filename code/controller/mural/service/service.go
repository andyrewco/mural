package service

import (
	"fmt"
	"mural/db"
	"mural/model"
)

func ResetSelected(all_tiles [][]model.Tile) [][]model.Tile {
	first_row := all_tiles[0]
	size := len(first_row)
	new_tiles := model.NewTiles(size)

	for _, row := range all_tiles {
		for _, tile := range row {
			tile := model.Tile{
				Penalty: tile.Penalty,
				I: tile.I,
				J: tile.J,
				Selected:  false,
				Flipped: tile.Flipped,
			}

			new_tiles[tile.I][tile.J] = tile
		}
	}

	return new_tiles
}

func ComputeShareable(
	session model.Session,
	current_game model.Game,
) string {
	text := fmt.Sprintf("Mural #%d Score: %d\n", current_game.GameKey, session.CurrentScore)

	// need to make tiles
	for _, row := range session.Board.Tiles {
		for _, tile := range row {
			if tile.Flipped {
				text += "⬜"
			} else {
				text += "🟪"
			}
		}
		text += "\n"
	}


	return text
}

func GetCorrectAnswer(answers []model.Answer) model.Answer {
	var answer model.Answer
	for _, a := range answers {
		if a.IsCorrect {
			answer = a
		}
	}

	return answer
}

func GetCurrentMural(
	user_key string,
) (*model.Mural, error) {
	current_game, err := db.DAL.GetCurrentGameInfo()
	if err != nil {
		return nil, err
	}

	current_session, err := db.DAL.GetGameSessionForUser(user_key)
	if err != nil {
		return nil, err
	}

	user_stats, _ := db.DAL.GetStatsForUser(user_key)
	return &model.Mural{
		Game: *current_game,
		Session: *current_session,
		UserStats: user_stats,
	}, nil
}