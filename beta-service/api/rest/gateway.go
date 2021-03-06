package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ProvideRoutes(rest *Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
		rest.Authenticator,
	)
	router.Mount("/betas", restRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.Authorizer)
	r.Get("/{betaID}", rest.GetOne)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.Post("/", rest.Post)
	r.Post("/multiply", rest.Post)
	r.Patch("/", rest.Patch)
	r.Get("/count", rest.Count)
	r.Post("/count", rest.Count)
	return
}
