package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hugo1707/task-api/handler"
	"github.com/hugo1707/task-api/middleware"
)

func main() {
	l := log.New(os.Stdout, "[task-api]: ", log.LstdFlags|log.Lshortfile)

	mux := http.NewServeMux()
	mux.Handle("/home", handler.NewHome())
	mux.Handle("/slow-home", middleware.Slow(handler.NewHome(), 25*time.Second))
	mux.Handle("/products", handler.NewProduct(l))

	s := &http.Server{
		Addr:        ":9090",
		Handler:     mux,
		IdleTimeout: 30 * time.Second,
	}

	go runServer(s, l)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Wait for the signal
	<-c

	l.Println("server will shutdown in 30 seconds")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		l.Fatalf("server shutdown %s\n", err)
	}
}

func runServer(s *http.Server, l *log.Logger) {
	l.Printf("server listening on port %s\n", s.Addr)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Fatalf("server listen %s\n", err)
	}
}
