package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rioten353/student_api/internal/config"
)

func main() {

	//load config

	cfg := config.MustLoad()

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	//setup server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server is running on " + cfg.Address)
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Server is not running: ", slog.String("error ", err.Error()))
		}
	}()

	<-done

	slog.Info("Shuddown server: " + cfg.Address)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancle()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

}
