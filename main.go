package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/golang-web/internal/config"
	"example.com/golang-web/internal/http/router"
	"example.com/golang-web/internal/model"
	"example.com/golang-web/internal/repo"
	"example.com/golang-web/internal/services/impl"
)

func main() {
	db := config.InitDB()
	if err := db.AutoMigrate(&model.UserModel{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	userRepo := repo.NewUserRepo(db)
	authSvc := impl.NewAuthService(userRepo)
	userSvc := impl.NewUserService(userRepo)

	r := router.NewRouter(authSvc, userSvc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		slog.Info("server starting", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	slog.Info("shutting down", "signal", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "err", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		_ = sqlDB.Close()
	}

	slog.Info("server stopped")
}
