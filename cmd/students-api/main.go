package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sonu31/student-api/internal/config"
	student "github.com/sonu31/student-api/internal/http/handlers"
	"github.com/sonu31/student-api/internal/storage/sqlite"
)

func main() {
	fmt.Println("HI Go")
	//load confg
	cfg := config.MustLoad()
	storage, sqlerr := sqlite.New(cfg)
	if sqlerr != nil {
		log.Fatal(sqlerr)

	}
	slog.Info("Storage Initlialized ", slog.String("env", cfg.Env), slog.String("Verion", "1.0.0"))

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.Create(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	// database setup

	//setup router

	//setuo server
	serve := http.Server{
		Addr:    cfg.Adds,
		Handler: router,
	}
	slog.Info("Server Start", slog.String("address", cfg.Adds))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := serve.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server: %v", err)

		}

	}()
	// Gracefully server shutdow
	<-done

	slog.Info("Shutting down the server")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	err := serve.Shutdown(ctx)
	if err != nil {
		slog.Error("Faild to shutDown Server", slog.String("Error", err.Error()))

	}

	slog.Info("Server ShutDown Successfully")

}
