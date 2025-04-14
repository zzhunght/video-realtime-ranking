package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zzhunght/realtime-video-ranking/internal/config"
	"github.com/zzhunght/realtime-video-ranking/internal/interfaces/api/router"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Cannot load config ", err)
	}
	fmt.Printf("Config %v", cfg)
	router := router.NewRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router.SetupRouter(),
	}

	fmt.Println("Server is running on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
