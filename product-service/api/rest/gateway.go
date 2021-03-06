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
	)
	router.Mount("/products/sneakers", restRoutes(rest))
	router.Mount("/health", healthRoutes(rest))
	return router
}

func restRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(rest.auth.Authenticator)
	r.Use(rest.auth.UserSetter)
	r.Get("/{sneakerId}", rest.GetOne)
	r.Get("/query", rest.Get)
	r.Get("/", rest.Get)
	r.Post("/query", rest.Get)
	r.Get("/count", rest.Count)
	r.Post("/count", rest.Count)
	r.With(rest.auth.Authorizer).Post("/", rest.Post)
	r.With(rest.auth.Authorizer).Put("/", rest.Put)
	r.With(rest.auth.Authorizer).Put("/{sneakerId}/images", rest.PutImages)
	r.With(rest.auth.Authorizer).Patch("/", rest.Patch)
	r.With(rest.auth.Authorizer).Delete("/{sneakerId}", rest.Delete)
	return
}

func healthRoutes(rest *Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/live", rest.HealthZ)
	r.Get("/ready", rest.ReadyZ)
	return
}
