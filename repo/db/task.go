package db

import (
	"gorm.io/gorm"

	"github.com/ars0915/gogolook-exercise/entity"
)

func (s *AppRepo) ListTasks() (t []entity.Task, err error) {
	if err = s.db.Find(&t).Error; err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return t, nil
}

func (s *AppRepo) CreateTask(t entity.Task) (entity.Task, error) {
	if err := s.db.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (s *AppRepo) UpdateTask(id uint, t entity.Task) error {
	return s.db.Model(entity.Task{}).Where(`"id" = ?`, id).Updates(t).Error
}

func (s *AppRepo) GetTask(id uint) (task entity.Task, err error) {
	if err = s.db.Where(`"id" = ?`, id).First(&task).Error; err != nil {
		return task, err
	}
	return
}

func (s *AppRepo) DeleteTask(id uint) (err error) {
	if err = s.db.Where(`"id" = ?`, id).Delete(&entity.Task{}).Error; err != nil {
		return err
	}
	return
}
