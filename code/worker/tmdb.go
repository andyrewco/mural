package worker

import (
	"log/slog"
	"mural/controller/movie"
	"mural/db"
)

type TMDBWorker struct { }

func NewTMDBWorker() TMDBWorker{
	return TMDBWorker{}
}

func (tw TMDBWorker) CacheAnswers() {
	slog.Info("Caching Answers From TMDB")
	movie_controller := movie.NewTMDBController()

	// get current page
	current_page, err := db.DAL.GetCurrentMoviePageFromDB()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	next_page := *current_page + 1
	answers, err := movie_controller.GetAnswers(next_page)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// cache answers
	err = db.DAL.CacheAnswersInDatabase(answers)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// set next page
	err = db.DAL.SetCurrentMoviePageFromDB()
	if err != nil {
		slog.Error(err.Error())
		return
	}
}