package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rahmatrdn/go-skeleton/config"
	"github.com/rahmatrdn/go-skeleton/internal/queue"
	"github.com/rahmatrdn/go-skeleton/internal/queue/consumer"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
}

type GoSkeletonWorker struct {
	ctx   context.Context
	queue queue.Queue
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("[Worker] topic not found, please use 'go run cmd/worker/main.go your.topic-key'")
	}

	log.Println("Starting WORKER")

	var app GoSkeletonWorker
	var err error

	app.ctx = context.Background()
	cfg := config.NewConfig()

	app.queue, err = config.NewRabbitMQInstance(app.ctx, &cfg.RabbitMQOption)
	if err != nil {
		log.Fatal(err)
	}

	// MongoDB Repository
	// logMongoRepo := mongodb.NewLogRepository(app.mongoDB)

	// Consumer
	// logConsumer := consumer.NewLogConsumer(context.Background(), logMongoRepo)
	// exampleConsumer := consumer.NewExampleConsumer(context.Background(), logMongoRepo)
	registrationConsumer := consumer.NewRegistrationConsumer(context.Background())

	var interrupt = make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	switch os.Args[1] {
	// case queue.ProcessSyncLog:
	// 	log.Printf("[Worker] Listening to %v", queue.ProcessSyncLog)
	// 	go app.queue.HandleConsumedDeliveries(queue.ProcessSyncLog, logConsumer.ProcessSyncLog)
	// case queue.ProcessExample:
	// 	log.Printf("[Worker] Listening to %v", queue.ProcessExample)
	// 	go app.queue.HandleConsumedDeliveries(queue.ProcessExample, exampleConsumer.Process)
	case queue.ProcessRegistration:
		log.Printf("[Worker] Listening to %v", queue.ProcessRegistration)
		go app.queue.HandleConsumedDeliveries(queue.ProcessRegistration, registrationConsumer.Process)
	default:
		log.Fatalf("[Worker] topic not found : %v", os.Args[1])
	}

	<-interrupt
	log.Println("Shutting down the Worker...")

	if err = app.queue.Close(); err != nil {
		log.Printf("Fail shutting down Worker: %s\n", err.Error())
	} else {
		log.Println("Worker successfully shutdown")
	}

}
