package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/escopodev/spotify-artist/internal/handlers"
	"github.com/escopodev/spotify-artist/internal/repositories"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	port := os.Getenv("APP_PORT")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	logger.Info("application is starting", slog.String("port", port))

	artistRepository := repositories.NewArtistPostgreSQL(db)
	artistsHandler := handlers.NewArtistsGet(artistRepository, logger)

	mux := http.NewServeMux()
	mux.Handle("GET /api/v1/artists/{uid}", artistsHandler)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
