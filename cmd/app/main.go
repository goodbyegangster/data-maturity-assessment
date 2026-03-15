package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"data-maturity-assessment/internal/handler"
	"data-maturity-assessment/internal/middleware"
	"data-maturity-assessment/internal/service"
)

func main() {
	svc, err := service.NewMaturityService("data")
	if err != nil {
		log.Fatalf("failed to initialize service: %v", err)
	}

	topHandler := handler.NewTopHandler(svc)
	assessmentHandler := handler.NewAssessmentHandler(svc)
	resultHandler := handler.NewResultHandler(svc)

	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.Htmx)

	// 静的ファイル
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// ページエンドポイント
	r.Get("/", topHandler.Top)
	r.Get("/assessment", assessmentHandler.Assessment)
	r.Post("/result", resultHandler.Result)

	// Partial エンドポイント (htmx 専用)
	r.Get("/partials/questions", assessmentHandler.Questions)
	r.Get("/partials/model-description", topHandler.ModelDescription)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
