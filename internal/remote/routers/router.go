package routers

func NewRouter(task handlers.ReqHandler) chi.Router {

	r := chi.NewRouter()

	r.Get("/task/{taskID}", task.Get)
	r.Post("/task", task.Create)
	return r
}
