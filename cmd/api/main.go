package main

import (
	"log"

	"github.com/Shakezidin/pkg/config"
	"github.com/Shakezidin/pkg/rabbitmq"
)

func main() {
	cfg, err := config.Configuration()
	if err != nil {
		log.Fatalf("unable to load config file, aborting")
	}
	rabbitmq.ConsumeConfirmationMessages(cfg)

}
