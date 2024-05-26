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

		var orgDto resources.TaskDto
		Success(w, orgDto.DomainToDto(tsk))
	}
}

func (c TaskConroller) Update() http.HandlerFunc {
	return c.Update()
}

func (c TaskConroller) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tsk := r.Context().Value(UserKey).(domain.User)

		err := c.taskService.Delete(tsk.Id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
