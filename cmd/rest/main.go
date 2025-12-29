package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ilyasa1211/url-shortener-demo/cmd/rest/handler"
	"github.com/ilyasa1211/url-shortener-demo/internal/application"
	"github.com/ilyasa1211/url-shortener-demo/internal/infrastructure/database/sqlite"
)

func main() {
	host := flag.String("host", "0.0.0.0", "HTTP server listen host")
	port := flag.Int("port", 8080, "HTTP server listen port")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	log.Println("Connecting to database")
	db := sqlite.Connect()
	log.Println("Successfully connected to database")

	siteRepo := sqlite.NewSiteRepository(db)
	siteService := application.NewSiteService(siteRepo)
	siteHandler := handler.NewSiteHandler(siteService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK\n")
	})
	mux.HandleFunc("GET /sites", siteHandler.Index)
	mux.HandleFunc("GET /sites/{aliasUrl}", siteHandler.Show)
	mux.HandleFunc("POST /sites", siteHandler.Create)
	mux.HandleFunc("PUT /sites/{aliasUrl}", siteHandler.Update)
	mux.HandleFunc("DELETE /sites/{aliasUrl}", siteHandler.Delete)

	addr := fmt.Sprintf("%s:%d", *host, *port)

	httpSrv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	httpErr := make(chan error, 1)

	go func() {
		httpErr <- httpSrv.ListenAndServe()
	}()

	log.Printf("HTTP server listening on: %s\n", addr)

	select {
	case <-ctx.Done():
		log.Println("Signal interrupt received")
		log.Println("Shutting down HTTP server")
		if err := httpSrv.Shutdown(context.Background()); err != nil {
			log.Println("Failed to shutdown server")
		}
	case <-httpErr:
		stop()
	}

}
