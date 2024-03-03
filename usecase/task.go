package usecase

import (
	"context"

	"github.com/ars0915/gogolook-exercise/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (h TaskHandler) ListTasks(ctx context.Context) (tasks []entity.Task, err error) {
	return h.db.ListTasks()
}

func (h TaskHandler) GetTask(ctx context.Context, id uint) (task entity.Task, err error) {
	task, err = h.db.GetTask(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return task, ErrorTaskNotFound
		}
		return task, errors.Wrap(err, "get task")
	}
	return
}

func (h TaskHandler) CreateTask(ctx context.Context, t entity.Task) (entity.Task, error) {
	return h.db.CreateTask(t)
}

func (h TaskHandler) UpdateTask(ctx context.Context, id uint, t entity.Task) (entity.Task, error) {
	if err := h.db.UpdateTask(id, t); err != nil {
		return entity.Task{}, err
	}

	task, err := h.db.GetTask(id)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (h TaskHandler) DeleteTask(ctx context.Context, id uint) error {
	return h.db.DeleteTask(id)
}
