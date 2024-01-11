package main

import (
	"async-queue/internal/config"
	"async-queue/internal/services/emailService"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	cnf := config.Get()

	redisConnection := asynq.RedisClientOpt{
		Addr:     cnf.Redis.Addr,
		Password: cnf.Redis.Pass,
	}

	emailService := emailService.NewEmailService(cnf)

	worker := asynq.NewServer(redisConnection, asynq.Config{
		Concurrency: 4,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(emailService.SendEmailQueue())

	if err := worker.Run(mux); err != nil {
		log.Printf("err: %v", err)
	}
}
