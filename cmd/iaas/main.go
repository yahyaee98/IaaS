package main

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"iaas/internal/api"
	"iaas/internal/cache"
	"iaas/internal/log"
	"iaas/internal/repository"
	"iaas/internal/upstream"
	"iaas/pkg/google_books"
	"iaas/pkg/itunes"
	"os"
	"os/signal"
	"time"
)

// Configuration is an struct which will hold values from env.
type Configuration struct {
	Port                    int    `envconfig:"PORT" default:"8080"`
	RedisAddress            string `envconfig:"REDIS_ADDRESS" default:"127.0.0.1:6379"`
	RedisPassword           string `envconfig:"REDIS_PASSWORD" default:""`
	RedisDb                 int    `envconfig:"REDIS_DB" default:"0"`
	GoogleBooksApiKey       string `envconfig:"GOOGLE_BOOKS_API_KEY" default:""`
	ResultLimitPerContent   int    `envconfig:"RESULT_LIMIT_PER_CONTENT" default:"5"`
	UpstreamTimeout         int    `envconfig:"UPSTREAM_TIMEOUT" default:"5"`
	GracefulShutdownTimeout int    `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT" default:"15"` // in seconds
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer log.Sync()

	var c Configuration
	if err := envconfig.Process("", &c); err != nil {
		log.Errorw("failed to process env configuration",
			"error", err,
		)
	}

	// Handle graceful shutdowns.
	go handleShutdown(cancel)

	upstreams := []upstream.Upstream{
		upstream.NewGoogleBooksUpstream(
			google_books.NewGoogleBooks(
				c.ResultLimitPerContent,
				c.GoogleBooksApiKey,
				time.Duration(c.UpstreamTimeout)*time.Second,
				log.Errorw,
			),
			func(duration time.Duration) {},
		),
		upstream.NewItunesUpstream(
			itunes.NewItunes(
				c.ResultLimitPerContent,
				time.Duration(c.UpstreamTimeout)*time.Second,
				log.Errorw,
			),
			func(duration time.Duration) {},
		),
	}

	cacheInstance := cache.NewRedisCache(c.RedisAddress, c.RedisPassword, c.RedisDb)

	rep := repository.NewRepository(upstreams, cacheInstance)

	api.StartServer(ctx, fmt.Sprintf(":%d", c.Port), c.GracefulShutdownTimeout, rep)

}

// handleShutdown calls cancel func when an interrupt signal is received from OS which leads to
// the main application's context cancellation.
func handleShutdown(cancel func()) {
	c := make(chan os.Signal, 1)

	// Express interest in interrupt signals.
	signal.Notify(c, os.Interrupt)

	// Block until an interrupt signal is received from OS.
	<-c

	// Cancel the main application's context.
	cancel()
}
