package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	Find(t domain.Task) domain.Task
	Update(r domain.Task) (domain.Task, error)
	Delete(id uint64) error
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository)  TaskService {
	return taskService{
		taskRepo: tr,
	}
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Find(id uint64) (interface{}, error) {
	tsk, err := s.taskRepo.Find(id)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return nil, err
	}

	return tsk, nil
}

func (s taskService) Update(r domain.Task) (domain.Task, error) {
	room, err := s.taskRepo.Update(r)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}

	return room, nil
}

func (s taskService) Delete(id uint64) error {
	err := s.taskRepo.Delete(id)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return err
	}

	return nil
}

 