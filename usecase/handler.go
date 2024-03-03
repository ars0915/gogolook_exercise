package usecase

import (
	"context"

	"github.com/ars0915/gogolook-exercise/repo"
)

type AppHandler struct {
	Task
}

type NewHandlerOption func(*AppHandler)

func newHandler(optFn ...NewHandlerOption) *AppHandler {
	h := &AppHandler{}

	for _, o := range optFn {
		o(h)
	}

	return h
}

func WithinTransaction(ctx context.Context, s repo.App, tFunc func(ctx context.Context) error) error {
	var isNewTx bool
	tx := repo.ExtractTx(ctx)
	// there is no tx inject to ctx, so new one
	if tx == nil {
		tx = s.Begin()
		defer tx.Rollback()
		isNewTx = true
	}

	// run callback
	if err := tFunc(repo.InjectTx(ctx, tx)); err != nil {
		return err
	}

	// tx extract from ctx don't need commit
	if isNewTx {
		return tx.Commit()
	}
	return nil
}

type TaskHandler struct {
	db repo.App
}

func NewTaskHandler(db repo.App) *TaskHandler {
	return &TaskHandler{
		db: db,
	}
}

func WithTask(i Task) func(h *AppHandler) {
	return func(h *AppHandler) { h.Task = i }
}
