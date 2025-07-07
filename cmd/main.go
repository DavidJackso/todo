package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DavidJackso/TodoApi/internal/config"
	"github.com/DavidJackso/TodoApi/internal/database"
	Handler "github.com/DavidJackso/TodoApi/internal/handler"
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/DavidJackso/TodoApi/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic("no .env file found")
	}

	cfg := config.SetupConfig()

	db := database.ConnectToDb(cfg)
	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{}, &models.Tag{})

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := Handler.NewHandler(service)
	router := handler.InitRouting()

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port),
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
	}

	done := make(chan os.Signal, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatalf("failed start server:%v", err)
		}
		logrus.Info("Server started")
	}()

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown failed: %v", err)
	}

	logrus.Info("Server Shutdown")

}
