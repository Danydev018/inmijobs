package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/api"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/core"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/database"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Fatal Error connecting to database: %v", err)
	}


	authRepository := repository.NewAuthRepository(*db)
	connRepository := repository.NewConnectionRepository(db)

	
	authService := core.NewAuthService(*authRepository)


	pingHandler := api.NewPingHandler(*authService)
	connHandler := api.NewConnectionHandler(connRepository)
	

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Route("/api", func(r chi.Router) {
        r.Get("/ping", pingHandler.Ping)
        
        // Rutas de Conexiones (Amigos)
        r.Route("/connections", func(r chi.Router) {
            r.Get("/test", connHandler.Ping)          
            r.Post("/", connHandler.CreateConnection)      
            r.Put("/{id}", connHandler.UpdateConnection)   
            r.Delete("/{id}", connHandler.DeleteConnection)
        })
    })

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}