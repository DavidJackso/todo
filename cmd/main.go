package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DavidJackso/TodoApi/internal/database"
	Handler "github.com/DavidJackso/TodoApi/internal/handler"
	"github.com/DavidJackso/TodoApi/internal/models"
	"github.com/DavidJackso/TodoApi/internal/repository"
	"github.com/DavidJackso/TodoApi/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {

	db := database.ConnectToDb()
	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{}, &models.Tag{})
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := Handler.NewHanlder(service)
	router := handler.InitRouting()

	//TODO: Move to config
	srv := http.Server{
		Addr:    "localhost:8081",
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	go func() {
		srv.ListenAndServe()
	}()

	logrus.Info("Server started")

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	logrus.Info("Server Shutdown")

}
