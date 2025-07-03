package service

import (
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/sirupsen/logrus"
)

type TodoService struct {
	rep *repository.Repository
}

func NewTodoService(repository *repository.Repository) *TodoService {
	return &TodoService{
		rep: repository,
	}
}

func (s *TodoService) CreateTask(task models.Task) (int, error) {
	id, err := s.rep.CreateTask(task)
	if err != nil {
		logrus.Info("error create task")
		return 0, err
	}
	return id, err
}

func (s *TodoService) GetTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (s *TodoService) DeleteTask(id int) error {
	return nil
}

func (s *TodoService) UpdateTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (s *TodoService) GetAllTask() ([]models.Task, error) {
	return nil, nil
}
