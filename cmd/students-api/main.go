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
)

func main() {
	fmt.Println("HI Go")
	//load confg
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello to stu api sk"))
	})

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
