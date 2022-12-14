package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/scys12/cake-store/internal/api/handler"
	"github.com/scys12/cake-store/internal/service/mocks"
	"github.com/scys12/cake-store/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_InsertCake(t *testing.T) {
	testcases := []struct {
		name        string
		err         error
		resultCode  int
		requestBody string
	}{
		{
			name:        "Insert cake successfully",
			err:         nil,
			resultCode:  http.StatusCreated,
			requestBody: `{"title": "Lemon", "description": "cheesecake", "rating": 12}`,
		},
		{
			name:        "Failed Bad Request Body",
			err:         errors.New("error"),
			resultCode:  http.StatusBadRequest,
			requestBody: `{"title": "Lemon", "description": "cheesecake", "rating": }`,
		},
		{
			name:        "Failed to insert cake",
			err:         errors.New("error"),
			resultCode:  http.StatusInternalServerError,
			requestBody: `{"title": "Lemon", "description": "cheesecake", "rating": 12}`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/cakes", strings.NewReader(tc.requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			controller := gomock.NewController(t)
			defer controller.Finish()

			service := mocks.NewMockICakeService(controller)
			respBody := &types.CakeResponse{}
			service.EXPECT().InsertCake(gomock.Any()).Return(respBody, tc.err).AnyTimes()

			h := handler.NewCakeHandler(service)
			h.InsertCake(resp, req)
			assert.Equal(t, tc.resultCode, resp.Code)
		})
	}
}

func Test_UpdateCake(t *testing.T) {
	testcases := []struct {
		name        string
		err         error
		resultCode  int
		requestBody string
		id          int64
	}{
		{
			name:        "Update cake successfully",
			id:          1,
			err:         nil,
			resultCode:  http.StatusCreated,
			requestBody: `{"title": "Lemon", "description": "cheesecake", "rating": 12}`,
		},
		{
			name:        "Failed Bad Request Body",
			id:          1,
			err:         errors.New("error"),
			resultCode:  http.StatusBadRequest,
			requestBody: `{"title": "Lemon", "description": "cheesecake", "rating": }`,
		},
		{
			name:        "Failed to update cake",
			id:          1,
			err:         errors.New("error"),
			resultCode:  http.StatusInternalServerError,
			requestBody: `{"title": "Lemon", "description": "cheesecake", "rating": 12}`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPatch, fmt.Sprint("/cakes/", tc.id), strings.NewReader(tc.requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			controller := gomock.NewController(t)
			defer controller.Finish()

			service := mocks.NewMockICakeService(controller)
			respBody := &types.CakeResponse{}
			service.EXPECT().UpdateCake(gomock.Any(), gomock.Any()).Return(respBody, tc.err).AnyTimes()

			h := handler.NewCakeHandler(service)
			h.UpdateCake(resp, req)
			assert.Equal(t, tc.resultCode, resp.Code)
		})
	}
}
