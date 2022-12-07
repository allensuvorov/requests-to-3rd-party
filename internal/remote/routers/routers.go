package routers

import "jobsearch/companies/axxonsoft/projects/3rdparty/internal/remote/handlers"

func NewRouter(task handlers.TaskHandler) chi.Router {

	r := chi.NewRouter()

	r.Get("/task/{taskID}", task.Get)
	r.Post("/task", task.Create)
	return r
}
