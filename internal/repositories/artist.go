package repositories

import "github.com/escopodev/spotify-artist/internal/models"

type Artist interface {
	GetByUID(uid string) (models.Artist, error)
}
