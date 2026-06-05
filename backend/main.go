package main

import (
	"backend/database"
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	sugar := logger.GetLogger().Sugar()

	sugar.Info("Server initialize")

	if err := godotenv.Load(); err != nil {
		sugar.Warnf("Warning: .env file not found, relying on environment variables: %v", err)
	} else {
		sugar.Info("Success to load .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		sugar.Fatal("DATABASE_URL is not set in environment variables")
	}

	db, err := database.Connect(dbURL)
	if err != nil {
		sugar.Fatal(err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db.Pool)

	h := &handler.Handler{DB: db, ProductRepo: productRepo}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)

	srv := &http.Server{
		Addr:         ":" + config.GetPort(),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		sugar.Info("Server is running...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalf("Error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	sugar.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		sugar.Fatalf("Server forced to shutdown: %v", err)
	}

	sugar.Info("Server exited")
}
