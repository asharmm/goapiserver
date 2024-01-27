package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/asharmm/goapiserver/internal/middleware"
)

func Handlers(r *chi.Mux) {
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router){
		
		//middleware for /account route
		router.Use(middleware.Autherization)

		router.Get("/coins",GetCoinBalance)
	})
}