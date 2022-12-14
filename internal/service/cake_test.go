package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/scys12/cake-store/internal/repositories/mocks"
	"github.com/scys12/cake-store/internal/service"
	"github.com/scys12/cake-store/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_InsertCake(t *testing.T) {
	defID := int64(1)
	testcases := []struct {
		name string
		err  error
		req  types.CakeRequest
		resp *types.CakeResponse
	}{
		{
			name: "Insert cake successfully",
			err:  nil,
			req: types.CakeRequest{
				Title: "Lemon",
			},
			resp: &types.CakeResponse{
				ID:    defID,
				Title: "Lemon",
			},
		},
		{
			name: "Failed to insert cake",
			err:  errors.New("error"),
			req: types.CakeRequest{
				Title: "Lemon",
			},
			resp: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mocks.NewMockICakeRepo(controller)
			repo.EXPECT().InsertCake(gomock.Any()).Return(defID, tc.err).AnyTimes()

			svc := service.NewCakeService(repo)
			resp, err := svc.InsertCake(tc.req)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.resp, resp)

		})
	}
}

func Test_UpdateCake(t *testing.T) {
	defID := int64(1)
	testcases := []struct {
		name string
		err  error
		req  types.CakeRequest
		resp *types.CakeResponse
	}{
		{
			name: "Update cake successfully",
			err:  nil,
			req: types.CakeRequest{
				Title: "Lemon",
			},
			resp: &types.CakeResponse{
				ID:    defID,
				Title: "Lemon",
			},
		},
		{
			name: "Failed to update cake",
			err:  errors.New("error"),
			req: types.CakeRequest{
				Title: "Lemon",
			},
			resp: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mocks.NewMockICakeRepo(controller)
			repo.EXPECT().UpdateCake(gomock.Any()).Return(tc.err).AnyTimes()

			svc := service.NewCakeService(repo)
			resp, err := svc.UpdateCake(defID, tc.req)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.resp, resp)

		})
	}
}

func Test_DeleteCake(t *testing.T) {
	defID := int64(1)
	testcases := []struct {
		name string
		err  error
	}{
		{
			name: "Delete cake successfully",
			err:  nil,
		},
		{
			name: "Failed to delete cake",
			err:  errors.New("error"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mocks.NewMockICakeRepo(controller)
			repo.EXPECT().DeleteCake(gomock.Any()).Return(tc.err).AnyTimes()

			svc := service.NewCakeService(repo)
			err := svc.DeleteCake(defID)
			assert.Equal(t, tc.err, err)

		})
	}
}

func Test_GetCakeDetail(t *testing.T) {
	defID := int64(1)
	testcases := []struct {
		name string
		err  error
		resp *types.CakeResponse
		cake *types.Cake
	}{
		{
			name: "Get cake successfully",
			err:  nil,
			resp: &types.CakeResponse{
				ID:    1,
				Title: "Lemon",
			},
			cake: &types.Cake{
				ID:    1,
				Title: "Lemon",
			},
		},
		{
			name: "Failed to get cake",
			err:  errors.New("error"),
			resp: nil,
			cake: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mocks.NewMockICakeRepo(controller)
			repo.EXPECT().GetCakeDetail(gomock.Any()).Return(tc.cake, tc.err).AnyTimes()

			svc := service.NewCakeService(repo)
			resp, err := svc.GetCakeDetail(defID)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.resp, resp)

		})
	}
}
