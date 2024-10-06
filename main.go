package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("Server is running on", cfg.Address)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start server")
	}

}
