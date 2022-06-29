package router

import (
	"github.com/go-chi/chi/v5"
	"princebillygk.portfolio.io/controller"
)

func SetupPortfolio(r *chi.Mux, c controller.Portfolio) {
	r.Get("/", c.Get)
}
