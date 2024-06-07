package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go-http/docs"
	entrypoint "go-http/internal/entrypoints/http"
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func (r *Router) LoadRouters() {
	routes := chi.NewRouter()

	routes.Get("/", entrypoint.Index)

	routes.Route("/test", func(rout chi.Router) {
		//r.Use(ArticleCtx)
		rout.Get("/", entrypoint.Test)
	})

	//r.Ctx.Mount("/swagger", httpSwagger.WrapHandler)
	r.Ctx.Mount("/api", routes)

	r.Ctx.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
	))
}

func (r *Router) ListenServer() {
	http.ListenAndServe(":3000", r.Ctx)
}
