package service

import "github.com/scys12/cake-store/internal/types"

type ICakeService interface {
	InsertCake(types.InsertCakeRequest) (*types.CakeResponse, error)
	UpdateCake(int64, types.InsertCakeRequest) (*types.CakeResponse, error)
	DeleteCake(int64) error
}
