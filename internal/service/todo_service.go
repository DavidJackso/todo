package service

import (
	"fmt"

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
	task, err := s.rep.TaskRepository.GetTask(id)
	if err != nil {
		return task, err
	}
	fmt.Print(task)
	return task, nil
}

func (s *TodoService) DeleteTask(id int) error {
	err := s.rep.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) UpdateTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (s *TodoService) GetTasks() ([]models.Task, error) {
	return nil, nil

}
