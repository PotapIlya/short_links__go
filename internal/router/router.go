package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	//_ "go-http/docs"
	entrypoint "go-http/internal/entrypoints/http"
	"log"
	"net/http"
)

type Router struct {
	Ctx chi.Router
}

func (r *Router) InitRouters() {
	r.Ctx = chi.NewRouter()
}

func (r *Router) LoadMiddleware() {
	//r.Ctx.Use(middleware.RequestID)
	r.Ctx.Use(middleware.Logger)
	r.Ctx.Use(middleware.Recoverer)
	r.Ctx.Use(middleware.URLFormat)
}

func (r *Router) LoadRouters() {
	routes := chi.NewRouter()

	routes.Get("/", entrypoint.Index)

	routes.Route("/test", func(rout chi.Router) {
		//r.Use(ArticleCtx)
		rout.Get("/", entrypoint.Test)
	})

	r.Ctx.Mount("/api", routes)

	r.Ctx.Mount("/swagger", httpSwagger.WrapHandler)
}

func (r *Router) ListenServer() {
	if err := http.ListenAndServe(":3000", r.Ctx); err != nil {
		log.Printf("Failed to launch api server:%+v\n", err)
	}
}
