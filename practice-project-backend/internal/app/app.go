package app

import (
	"fmt"
	"net/http"

	"github.com/NutsBalls/practice-project-backend/internal/endpoint"
	"github.com/NutsBalls/practice-project-backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
	c *chi.Mux
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.c = chi.NewRouter()
	a.c.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	a.c.Get("/status", a.e.Status)
	a.c.Get("/parse", a.e.Parse)
	a.c.Get("/regions", a.e.Regions)
	a.c.Get("/jobs", a.e.Vacancies)
	return a, nil
}

func (a *App) Run(port int) error {

	portStr := fmt.Sprintf("http://localhost:%d", port)

	fmt.Println("server running at", portStr)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.c,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start http server %v", err)
	}

	return nil
}
