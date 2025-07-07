package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/sirupsen/logrus"
)

// TODO: добавить валидацию данных
type TodoService struct {
	rep *repository.Repository
}

func NewTodoService(repository *repository.Repository) *TodoService {
	return &TodoService{
		rep: repository,
	}
}

func (s *TodoService) CreateTask(task models.Task, userID uint) (uint, error) {
	id, err := s.rep.CreateTask(task, userID)
	if err != nil {
		logrus.WithError(err).Error("failed create task")
		return 0, err
	}
	return id, err
}

func (s *TodoService) GetTask(id uint, userID uint) (models.Task, error) {
	task, err := s.rep.GetTask(id, userID)
	if err != nil {
		logrus.WithError(err).Error("failed get task")
		return models.Task{}, err
	}
	return task, nil
}

func (s *TodoService) DeleteTask(id uint, userID uint) error {
	err := s.rep.DeleteTask(id, userID)
	if err != nil {
		logrus.WithError(err).Error("failed to delete task")
		return err
	}
	return nil
}

func (s *TodoService) UpdateTask(id, userID uint, updTask models.Task) (models.Task, error) {
	task, err := s.rep.UpdateTask(id, userID, updTask)
	if err != nil {
		logrus.WithError(err).Error("failed update task")
		return models.Task{}, err
	}
	return task, nil
}

func (s *TodoService) GetTasks(userID uint) ([]models.Task, error) {
	tasks, err := s.rep.GetTasks(userID)
	if err != nil {
		logrus.WithError(err).Error("failed to get tasks")
		return nil, err
	}

	return tasks, nil
}
