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
	postgressSql "github.com/sonu31/student-api/internal/storage/postgrsSql"
)

func main() {
	fmt.Println("HI Go")
	//load confg
	cfg := config.MustLoad()

	// storage, sqlerr := sqlite.New(cfg)
	// if sqlerr != nil {
	// 	log.Fatal(sqlerr)

	// }
	// slog.Info("Storage Initlialized ", slog.String("env", cfg.Env), slog.String("Verion", "1.0.0"))
	// SQL
	// router := http.NewServeMux()
	// router.HandleFunc("POST /api/students", student.Create(storage))
	// router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	// router.HandleFunc("GET /api/students", student.GetList(storage))
	// router.HandleFunc("PUT /api/students/{id}", student.UpdateItme(storage))

	postgreSql, pSqlerr := postgressSql.New(cfg)
	if pSqlerr != nil {
		log.Fatal(pSqlerr)
	}

	//PostgrSql
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.Create(postgreSql))
	router.HandleFunc("GET /api/students/{id}", student.GetById(postgreSql))
	router.HandleFunc("GET /api/students", student.GetList(postgreSql))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateItme(postgreSql))

	// database setup

	//setup router

	//setuo server
	/* 	serve := http.Server{
	   		Addr:    cfg.Adds,
	   		Handler: router,
	   	}
	   	slog.Info("Server Start", slog.String("address", cfg.Adds))

	   	done := make(chan os.Signal, 1)

	   	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) */

	serve := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	slog.Info("Server Start", slog.String("address", cfg.HTTPServer.Address))

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
