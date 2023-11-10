package movie

import (
	"fmt"
	"mural/api"
	"mural/config"
	"mural/model"
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)

type TMDBController struct {
	TMDBApi *tmdb.TMDb
}

func NewTMDBAPI() *tmdb.TMDb {
	config := tmdb.Config{
		APIKey:   os.Getenv(config.EnvTMDBKey),
		Proxies:  nil,
		UseProxy: false,
	}

	return tmdb.Init(config)
}

func NewTMDBController() *TMDBController {
	api := NewTMDBAPI()
	return &TMDBController{
		TMDBApi: api,
	}
}


func (mc TMDBController) GetAnswers(
	page_number int,
) ([]model.Answer, error) {
	parameters := map[string]string{
		"page": fmt.Sprintf("%d", page_number),
	}

	movie_results, err := mc.TMDBApi.DiscoverMovie(parameters)
	if err != nil {
		return nil, err
	}


	var answers []model.Answer
	for i, mov := range movie_results.Results {
		movie := model.Movie{
			Poster: mov.PosterPath,
			ID: mov.ID,
			Name: mov.Title,
			Description: mov.Overview,
			ReleaseDate: mov.ReleaseDate,
		}

		answer := model.Answer{
			Movie: movie,
			Selected: false,
			IsCorrect: i == api.RandomAnswerKey ,
		}

		answers = append(answers, answer)
	}


	return answers, nil
}

func (mc TMDBController) GetAnswersByDecade(
	page_number int,
	decade string,
) ([]model.Answer, error) {
	lower_bound, upper_boud := getDecadeBounds(decade)

	parameters := map[string]string{
		"page": fmt.Sprintf("%d", page_number),
		"primary_release_date.lte": upper_boud,
		"primary_release_date.gte": lower_bound,
	}

	movie_results, err := mc.TMDBApi.DiscoverMovie(parameters)
	if err != nil {
		return nil, err
	}


	var answers []model.Answer
	for i, mov := range movie_results.Results {
		movie := model.Movie{
			Poster: mov.PosterPath,
			ID: mov.ID,
			Name: mov.Title,
			Description: mov.Overview,
			ReleaseDate: mov.ReleaseDate,
		}

		answer := model.Answer{
			Movie: movie,
			Selected: false,
			IsCorrect: i == api.RandomAnswerKey ,
		}

		answers = append(answers, answer)
	}


	return answers, nil
}