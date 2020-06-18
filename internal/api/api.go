package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"iaas/internal/log"
	"iaas/internal/repository"
	"net/http"
	"time"
)

// StartServer Starts the http api server
func StartServer(ctx context.Context, addr string, gracefulShutdownTimeout int, rep repository.Repository) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := mux.NewRouter()

	r.
		HandleFunc("/api/results", GetItemsHandler(rep)).
		Methods("GET")

	r.
		Handle("/api/metrics", promhttp.Handler()).
		Methods("GET")

	r.
		HandleFunc("/api/health", GetHealthHandler(ctx)).
		Methods("GET")

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Infow("HTTP server started.",
		"address", addr,
	)

	// Run the server in a goroutine so it doesn't block.
	go func(srv *http.Server, cancel context.CancelFunc) {
		if err := srv.ListenAndServe(); err != nil {
			log.Infow("underlying http server has reported an error",
				"error", err,
			)
			cancel()
		}
	}(srv, cancel)

	// Wait till we get a signal that we have to finish.
	<-ctx.Done()

	// Create a deadline to wait for.
	ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(gracefulShutdownTimeout))
	defer cancel()

	// Block till requests are cleared.
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorw("http server has failed to shutdown gracefully")
		return
	}

	log.Infow("http server shutdown")

	return
}
