package worker

func (s MuralScheduler) RegisterWorkers(
) error {
	// new mural worker
	mural_worker := NewMuralWorker()

	// new tmdb worker
	tmdb_worker := NewTMDBWorker()

	// // register session worker
	// s.Scheduler.WaitForSchedule().Every(1).Day().At("23:59").Do(mural_worker.SetupNewGame)

	// // register session worker
	// s.Scheduler.WaitForSchedule().Every(1).Day().At("23:59").Do(mural_worker.ResetGameSessions)

	// // register session worker
	// s.Scheduler.Every(1).Day().At("22:00").Do(tmdb_worker.CacheAnswers)


	// register session worker
	s.Scheduler.Every(1).Minute().Do(mural_worker.SetupNewGame)

	// register session worker
	s.Scheduler.Every(1).Minute().Do(mural_worker.ResetGameSessions)
	s.Scheduler.Every(1).Minute().Do(tmdb_worker.CacheAnswers)

	// register session worker
	s.Scheduler.Every(1).Minute().Do(tmdb_worker.CacheAnswers)
	return nil
}