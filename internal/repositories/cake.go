package repositories

import (
	"database/sql"
	"fmt"

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
		}).Errorln("[CakeRepo] Failed to insert cake")
		return 0, types.ErrFailedInsertCake
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to get ID")

		return 0, types.ErrFailedGetLastIDDB
	}

	return id, nil
}

func (c *CakeRepo) UpdateCake(cake *types.Cake) error {
	res, err := c.db.Exec(`
	UPDATE cake SET
		title = ?,
		description = ?,
		rating = ?,
		image = ?
	WHERE
		id = ?
	`,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image,
		cake.ID,
	)

	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to update cake")
		return types.ErrFailedUpdateCake
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to get rows affected")

		return types.ErrFailedRowsAffected
	}

	if rowsAffected < 1 {
		return types.ErrNoRowsAffected
	}

	return nil
}

func (c *CakeRepo) DeleteCake(id int64) error {
	res, err := c.db.Exec(`
	DELETE 
	FROM cake
	WHERE
		id = ?
	`,
		id,
	)

	if err != nil {
		log.WithFields(log.Fields{
			"function": "DeleteCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to delete cake")

		return types.ErrFailedDeleteCake
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "DeleteCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to get rows affected")

		return types.ErrFailedRowsAffected
	}

	if rowsAffected < 1 {
		return types.ErrNoRowsAffected
	}

	return nil
}

func (c *CakeRepo) GetCakeDetail(id int64) (*types.Cake, error) {
	cake := new(types.Cake)

	if err := c.db.QueryRow(`
	SELECT 
		id, 
		title, 
		description, 
		rating, 
		image, 
		created_at, 
		updated_at
	FROM cake
	WHERE
		id = ?
	`,
		id,
	).Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdatedAt,
	); err != nil {
		log.WithFields(log.Fields{
			"function": "GetCakeDetail",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to get cake detail")

		return nil, types.ErrGetCakeDetail
	}

	return cake, nil
}

func (c *CakeRepo) GetListOfCake(sortBy string, pageNum int64) ([]types.Cake, error) {
	pageLimit := 5
	var cakes []types.Cake

	dbQuery := `
	SELECT 
		id, 
		title, 
		description, 
		rating, 
		image, 
		created_at, 
		updated_at
	FROM cake
	`

	if len(sortBy) > 0 {
		dbQuery = fmt.Sprint(dbQuery, "ORDER BY ", sortBy, " ASC ")
	} else {
		dbQuery = fmt.Sprint(dbQuery, "ORDER BY rating ASC, title ASC ")
	}
	dbQuery = fmt.Sprint(dbQuery, "LIMIT ", pageLimit, " OFFSET ", pageLimit*(int(pageNum)-1))

	rows, err := c.db.Query(dbQuery)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetListOfCake",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to get list of cake")

		return nil, types.ErrGetCakeList
	}
	defer rows.Close()

	for rows.Next() {
		cake := types.Cake{}
		err = rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cakes, nil
}

func (c *CakeRepo) CountAllCakes() (int64, error) {
	var count int64

	if err := c.db.QueryRow(`
	SELECT 
		COUNT(id)
	FROM cake
	`).Scan(&count); err != nil {
		log.WithFields(log.Fields{
			"function": "CountAllCakes",
			"error":    err.Error(),
		}).Errorln("[CakeRepo] Failed to count cakes")

		return 0, types.ErrCountCakes
	}

	return count, nil
}
