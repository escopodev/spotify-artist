package repositories

import (
	"database/sql"

	"github.com/escopodev/spotify-artist/internal/models"
	"github.com/lib/pq"
)

type ArtistPostgreSQL struct {
	db *sql.DB
}

func NewArtistPostgreSQL(db *sql.DB) Artist {
	return &ArtistPostgreSQL{
		db: db,
	}
}

func (r *ArtistPostgreSQL) GetByUID(uid string) (models.Artist, error) {
	query := `
		SELECT 
			a.id, a.uid, a.name, a.popularity, a.genres,
			i.id, i.uid, i.width, i.height,
			u.id, u.username
		FROM artists a
		LEFT JOIN images i ON a.id = i.artist_id
		LEFT JOIN artist_user au ON a.id = au.artist_id
		LEFT JOIN users u ON au.user_id = u.id
		WHERE a.uid = $1
	`

	rows, err := r.db.Query(query, uid)
	if err != nil {
		return models.Artist{}, err
	}
	defer rows.Close()

	var artist models.Artist
	images := make(map[int64]struct{})
	users := make(map[int64]struct{})

	for rows.Next() {
		var image models.Image
		var user models.User

		err := rows.Scan(
			&artist.ID, &artist.UID, &artist.Name, &artist.Popularity, pq.Array(&artist.Genres),
			&image.ID, &image.UID, &image.Width, &image.Height,
			&user.ID, &user.UserName,
		)
		if err != nil {
			return models.Artist{}, err
		}

		if _, found := images[image.ID]; !found {
			images[image.ID] = struct{}{}
			image.Compute()
			artist.Images = append(artist.Images, image)
		}

		if _, found := users[user.ID]; !found {
			users[user.ID] = struct{}{}
			artist.Users = append(artist.Users, user)
		}
	}

	if err = rows.Err(); err != nil {
		return models.Artist{}, err
	}

	artist.Compute()
	return artist, nil
}
