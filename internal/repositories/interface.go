package repositories

import "github.com/scys12/cake-store/internal/types"

type ICakeRepo interface {
	InsertCake(*types.Cake) (int64, error)
	UpdateCake(*types.Cake) error
	DeleteCake(int64) error
	GetCakeDetail(int64) (*types.Cake, error)
	GetListOfCake(string, int64) ([]types.Cake, error)
	CountAllCakes() (int64, error)
}
