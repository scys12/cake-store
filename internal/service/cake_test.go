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
		req  types.InsertCakeRequest
		resp *types.CakeResponse
	}{
		{
			name: "Insert cake successfully",
			err:  nil,
			req: types.InsertCakeRequest{
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
			req: types.InsertCakeRequest{
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
		req  types.InsertCakeRequest
		resp *types.CakeResponse
	}{
		{
			name: "Update cake successfully",
			err:  nil,
			req: types.InsertCakeRequest{
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
			req: types.InsertCakeRequest{
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
