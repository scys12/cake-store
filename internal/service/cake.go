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

func (c *CakeService) InsertCake(data types.CakeRequest) (*types.CakeResponse, error) {
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
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}
	return &resp, nil
}

func (c *CakeService) UpdateCake(id int64, data types.CakeRequest) (*types.CakeResponse, error) {
	cake := &types.Cake{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
	}

	err := c.cakeRepo.UpdateCake(cake)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateCake",
			"error":    err.Error(),
		}).Errorln("[CakeService] Failed to update database")

		return nil, err
	}

	resp := types.CakeResponse{
		ID:          cake.ID,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}
	return &resp, nil
}

func (c *CakeService) DeleteCake(id int64) error {
	err := c.cakeRepo.DeleteCake(id)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "UpdateCake",
			"error":    err.Error(),
		}).Errorln("[CakeService] Failed to delete cake")

		return err
	}

	return nil
}

func (c *CakeService) GetCakeDetail(id int64) (*types.CakeResponse, error) {
	cake, err := c.cakeRepo.GetCakeDetail(id)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetCakeDetail",
			"error":    err.Error(),
		}).Errorln("[CakeService] Failed to get cake detail")

		return nil, err
	}

	resp := types.CakeResponse(*cake)
	return &resp, nil
}

func (c *CakeService) GetListOfCake(sortBy string, pageNum int64) (*types.CakesResponse, error) {
	if pageNum < 1 {
		pageNum = 1
	}

	cakes, err := c.cakeRepo.GetListOfCake(sortBy, pageNum)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetCakeDetail",
			"error":    err.Error(),
		}).Errorln("[CakeService] Failed to get cake detail")

		return nil, err
	}

	totalData, err := c.cakeRepo.CountAllCakes()
	if err != nil {
		log.WithFields(log.Fields{
			"function": "GetCakeDetail",
			"error":    err.Error(),
		}).Errorln("[CakeService] Failed to count cakes")

		return nil, err
	}

	var cr []types.CakeResponse
	for _, c := range cakes {
		r := types.CakeResponse(c)
		cr = append(cr, r)
	}

	resp := &types.CakesResponse{
		Cakes:            cr,
		TotalData:        totalData,
		TotalDataPerPage: 5,
	}

	return resp, nil
}
