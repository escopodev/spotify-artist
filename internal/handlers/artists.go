package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/escopodev/spotify-artist/internal/repositories"
)

type ArtistsGet struct {
	repo   repositories.Artist
	logger *slog.Logger
}

func NewArtistsGet(repo repositories.Artist, logger *slog.Logger) *ArtistsGet {
	return &ArtistsGet{
		repo:   repo,
		logger: logger,
	}
}

func (handler *ArtistsGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	artist, err := handler.repo.GetByUID(r.PathValue("uid"))
	if err != nil {
		handler.logger.Error(
			"get artist by uid",
			slog.String("uid", r.PathValue("uid")),
			slog.Any("err", err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(artist)
	if err != nil {
		handler.logger.Error(
			"marshal artist",
			slog.String("uid", r.PathValue("uid")),
			slog.Any("err", err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
