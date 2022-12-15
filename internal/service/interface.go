package service

import "github.com/scys12/cake-store/internal/types"

type ICakeService interface {
	InsertCake(types.CakeRequest) (*types.CakeResponse, error)
	UpdateCake(int64, types.CakeRequest) (*types.CakeResponse, error)
	DeleteCake(int64) error
	GetCakeDetail(int64) (*types.CakeResponse, error)
	GetListOfCake(string, int64) (*types.CakesResponse, error)
}
