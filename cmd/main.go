package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	Handler "github.com/DavidJackso/TodoApi/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {

	handler := Handler.NewHanlder()
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
