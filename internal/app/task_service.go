package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	Find(t domain.Task) domain.Task
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
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

func (s taskService) Find(t domain.Task, userId int, status string, title string) ([]domain.Task, error) {
	task, err := s.taskRepo.Find(t, userId, status, title)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return []domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Update(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Update(t)
	if err != nil {
		log.Printf("UserService: %s", err)
		return domain.Task{}, err
	}

	return task, nil
}

// func (s userService) Delete(id uint64) error {
// 	err := s.userRepo.Delete(id)
// 	if err != nil {
// 		log.Printf("UserService: %s", err)
// 		return err
// 	}

	// return nil
// }
 