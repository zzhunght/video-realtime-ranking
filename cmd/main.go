package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	services "github.com/zzhunght/realtime-video-ranking/internal/application"
	"github.com/zzhunght/realtime-video-ranking/internal/config"
	mq "github.com/zzhunght/realtime-video-ranking/internal/infrastructure/messaging"
	"github.com/zzhunght/realtime-video-ranking/internal/infrastructure/persistence/postgres"
	"github.com/zzhunght/realtime-video-ranking/internal/infrastructure/persistence/redis"
	"github.com/zzhunght/realtime-video-ranking/internal/interfaces/api/handler"
	"github.com/zzhunght/realtime-video-ranking/internal/interfaces/api/router"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Cannot load config ", err)
	}

	// set up infra
	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database ", err)
	}

	defer db.Close()

	rd, err := redis.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		log.Fatal("Cannot connect to redis ", err)
	}
	defer rd.Client.Close()

	//--------------------------------------
	// setup repository
	postgresVideoRepo := postgres.NewVideoRepository(db)
	redisRankingRepo := redis.NewRedisRankingRepository(rd.Client)
	redisCachedRepo := redis.NewRedisCachedRepository(rd.Client)

	// setup services
	rankingService := services.NewRankingService(redisRankingRepo, redisCachedRepo, postgresVideoRepo)

	// -------------------------------------

	// setup mq
	kafkaProducer := mq.NewKafkaProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	defer kafkaProducer.Close()
	eventConsumer := mq.NewScoreConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.TopicGroup, rankingService)
	go eventConsumer.Start(context.Background())

	// setup handlers
	rankingHandler := handler.NewRankingHanlder(rankingService, kafkaProducer)

	// setup router
	router := router.NewRouter(
		rankingHandler,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: router.SetupRouter(),
	}

	var shutDownChan = make(chan os.Signal, 1)

	signal.Notify(shutDownChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		log.Println("Server is running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-shutDownChan

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	fmt.Println("exiting...")

}
