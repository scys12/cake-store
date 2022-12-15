package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/scys12/cake-store/internal/server"
	"github.com/scys12/cake-store/internal/service"
	"github.com/scys12/cake-store/internal/types"
	log "github.com/sirupsen/logrus"
)

type CakeHandler struct {
	cakeService service.ICakeService
}

func NewCakeHandler(service service.ICakeService) CakeHandler {
	return CakeHandler{
		cakeService: service,
	}
}

func (c *CakeHandler) InsertCake(w http.ResponseWriter, r *http.Request) {
	var insertReq types.CakeRequest

	err := json.NewDecoder(r.Body).Decode(&insertReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "InsertCake",
		}).Info("[CakeHandler] Unable to decode request body")

		server.RenderError(w, http.StatusBadRequest, types.ErrBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(insertReq)
	errs := ValidateRequest(validate, err)
	if errs != nil {
		log.WithFields(log.Fields{
			"error":    errs[0],
			"function": "InsertCake",
		}).Info("[CakeHandler] Request validation failed")

		server.RenderError(w, http.StatusBadRequest, errs[0])
		return
	}

	resp, err := c.cakeService.InsertCake(insertReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "InsertCake",
		}).Errorln("[CakeHandler] Failed to insert cake")

		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp)
}

func (c *CakeHandler) UpdateCake(w http.ResponseWriter, r *http.Request) {
	var updateReq types.CakeRequest

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "UpdateCake",
		}).Errorln("[CakeHandler] Unable to decode request body")

		server.RenderError(w, http.StatusBadRequest, types.ErrBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(updateReq)
	errs := ValidateRequest(validate, err)
	if errs != nil {
		log.WithFields(log.Fields{
			"error":    errs[0],
			"function": "InsertCake",
		}).Info("[CakeHandler] Request validation failed")

		server.RenderError(w, http.StatusBadRequest, errs[0])
		return
	}

	resp, err := c.cakeService.UpdateCake(id, updateReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "UpdateCake",
		}).Errorln("[CakeHandler] Failed to update cake")

		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	server.RenderResponse(w, http.StatusOK, resp)
}

func (c *CakeHandler) DeleteCake(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	err := c.cakeService.DeleteCake(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "DeleteCake",
		}).Errorln("[CakeHandler] Failed to delete cake")

		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	resp := types.CakeResponse{
		ID: id,
	}
	server.RenderResponse(w, http.StatusOK, resp)
}

func (c *CakeHandler) GetCakeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	resp, err := c.cakeService.GetCakeDetail(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "GetCakeDetail",
		}).Errorln("[CakeHandler] Failed to get cake detail")

		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	server.RenderResponse(w, http.StatusOK, resp)
}

func (c *CakeHandler) GetListOfCake(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")
	sortBy = GetValidatedSortBy(sortBy)

	page := r.URL.Query().Get("page")
	pageNum, _ := strconv.ParseInt(page, 10, 64)

	resp, err := c.cakeService.GetListOfCake(sortBy, pageNum)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "GetListOfCake",
		}).Errorln("[CakeHandler] Failed to get list of cake")

		server.RenderError(w, http.StatusInternalServerError, err)
		return
	}
	server.RenderResponse(w, http.StatusOK, resp)
}
