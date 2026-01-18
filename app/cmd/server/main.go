package main


import (
	"log"
	"os"
	"strconv"
	"time"
	"os/signal"
	"syscall"
	"context"
	"net/http"

	"ci-supplychain-playground/app/internal/httpapi"
	"ci-supplychain-playground/app/internal/version"
)

func main() {
	port := mustEnvInt("PORT", 8080)
	readTimeout := mustEnvDuration("READ_TIMEOUT", 5*time.Second)
	writeTimeout := mustEnvDuration("WRITE_TIMEOUT", 10*time.Second)

	mux :=httpapi.NewRouter(httpapi.RouterConfig{
		Version: version.Info(),
	})

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	go func() {
		log.Printf("Listening on %s (version: %s)", srv.Addr, version.Info().String())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func mustEnvInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("Invalid value for %s: %v", key, err)
	}
	return i
}

func mustEnvDuration(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Fatalf("Invalid value for %s: %v", key, err)
	}
	return d
}