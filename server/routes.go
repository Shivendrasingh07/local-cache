package server

import (
	"github.com/go-chi/chi"
)

func (srv *Server) initializeRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(cache chi.Router) {
		//		cache.Use(middleware.AuthMiddleware)
		cache.Post("/register", srv.Signup)
		cache.Post("/login", srv.Login)
		//cache.Post("/set", srv.set)
		//cache.Get("/get/{key}", srv.Get)
		//cache.Delete("/delete/{key}", srv.Delete)
	})

	return r
}
