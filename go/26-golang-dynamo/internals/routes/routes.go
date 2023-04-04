package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	serviceconfig "github.com/jsierrab3991/scripts/26-golang-dynamo/config"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/handlers/health"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/handlers/product"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/repositories/adapter"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceconfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouter(repository adapter.Interface) *chi.Mux {
	r.setConfigRouters()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := health.NewHandler(repository)
	r.router.Route("/health", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := product.NewHandler(repository)
	r.router.Route("/product", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) setConfigRouters() {
	r.EnableCORS()
	r.EnabledLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestId()
	r.EnableRealIP()
}

func (r *Router) EnabledLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestId() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
