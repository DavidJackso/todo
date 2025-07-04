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

func (s *TodoService) CreateTask(task models.Task, userID int) (int, error) {
	id, err := s.rep.CreateTask(task, userID)
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
	return task, nil
}

func (s *TodoService) DeleteTask(id int) error {
	err := s.rep.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) UpdateTask(id int, updTask models.Task) (models.Task, error) {
	task, err := s.rep.UpdateTask(id, updTask)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *TodoService) GetTasks(id int) ([]models.Task, error) {
	tasks, err := s.rep.GetTasks(id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return tasks, nil
}
