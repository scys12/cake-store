package service

import (
	"github.com/scys12/cake-store/internal/repositories"
	"github.com/scys12/cake-store/internal/types"
	log "github.com/sirupsen/logrus"
)

type CakeService struct {
	cakeRepo repositories.ICakeRepo
}

func NewCakeService(repo repositories.ICakeRepo) ICakeService {
	return &CakeService{
		cakeRepo: repo,
	}
}

func (c *CakeService) InsertCake(data types.InsertCakeRequest) (*types.CakeResponse, error) {
	cake := &types.Cake{
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
	}

	id, err := c.cakeRepo.InsertCake(cake)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "InsertCake",
			"error":    err.Error(),
		}).Errorln("[CakeService] Failed to insert database")

		return nil, err
	}

	resp := types.CakeResponse{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
	}
	return &resp, nil
}
