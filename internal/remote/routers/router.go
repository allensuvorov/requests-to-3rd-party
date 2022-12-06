package routers

func NewRouter(req handlers.ReqHandler) chi.Router {

	r := chi.NewRouter()

	r.Get("/{id}", req.Get)
	r.Post("/", req.Create)
	return r
}
