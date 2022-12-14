package repositories

import (
	"database/sql"

	"github.com/scys12/cake-store/internal/types"
	log "github.com/sirupsen/logrus"
)

type CakeRepo struct {
	db *sql.DB
}

func NewCakeRepo(db *sql.DB) ICakeRepo {
	return &CakeRepo{
		db: db,
	}
}

func (c *CakeRepo) InsertCake(cake *types.Cake) (int64, error) {
	var id int64
	res, err := c.db.Exec(`
	INSERT INTO cake (
		title,
		description,
		rating,
		image
	) VALUES (
		?,
		?,
		?,
		?
	)`,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image)

	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Problem querying to database")
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Problem querying to database")

		return 0, err
	}

	return id, nil
}
