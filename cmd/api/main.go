package api

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
	"github.com/Shakezidin/pkg/config"
)

func main() {
	cfg, err := config.Configuration()
	if err != nil {
		log.Fatalf("unable to load config file, aborting")
	}
	reader := kafka.NewKafkaReader()
	group, ctx := errgroup.WithContext(context.Background())
	log.Println("sending to email writer")
	group.Go(func() error {
		return reader.EmailWriter(ctx, cfg)
	})
	err = group.Wait()
	if err != nil {
		return
	}
	log.Println("reading from kafka complete")
}
