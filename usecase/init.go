package usecase

import "github.com/ars0915/gogolook-exercise/repo"

func InitHandler(db repo.App) Handler {
	// TODO: inject all usecase service by option pattern

	h := newHandler()

	return h
}
