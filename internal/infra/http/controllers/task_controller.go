package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskConroller struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskConroller {
	return TaskConroller{
		taskService: ts,
	}
}

func (c TaskConroller) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			BadRequest(w, err)
			return
		}

		task.UserId = user.Id
		task.Status = domain.New
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Created(w, tDto)
	}
}

func (c TaskConroller) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tsk := r.Context().Value(UserKey).(domain.Task)

		// Получение параметров поиска из запроса
		queryParams := r.URL.Query()
		status := queryParams.Get("status")
		title := queryParams.Get("title")

		// Вызов метода поиска в сервисе задач
		tasks, err := c.taskService.Find(tsk.Id, status, title)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tasksDto []resources.TaskDto
		for _, task := range tasks {
			var tDto resources.TaskDto
			tDto = tDto.DomainToDto(task)
			tasksDto = append(tasksDto, tDto)
		}

		Success(w, tasksDto)
	}
}

func (c TaskConroller) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tsk, err := requests.Bind(r, requests.UpdateUserRequest{}, domain.Task{})
		if err != nil {
			log.Printf("UserController: %s", err)
			BadRequest(w, err)
			return
		}

		t := r.Context().Value(UserKey).(domain.Task)
		t.Description = tsk.Description
		t.Title = tsk.Title
		tsk, err = c.taskService.Update(t)
		if err != nil {
			log.Printf("UserController: %s", err)
			InternalServerError(w, err)
			return
		}

		var userDto resources.TaskDto
		Success(w, userDto.DomainToDto(tsk))
	}
}

