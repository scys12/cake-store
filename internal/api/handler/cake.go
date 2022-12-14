package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
	var insertReq types.InsertCakeRequest

	err := json.NewDecoder(r.Body).Decode(&insertReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "InsertCake",
		}).Info("[CakeHandler] Unable to decode request body")

		server.RenderError(w, http.StatusBadRequest, err, time.Now())
		return
	}

	resp, err := c.cakeService.InsertCake(insertReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "InsertCake",
		}).Errorln("[CakeHandler] Failed to insert cake")

		server.RenderError(w, http.StatusInternalServerError, err, time.Now())
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp, time.Now())
}

func (c *CakeHandler) UpdateCake(w http.ResponseWriter, r *http.Request) {
	var updateReq types.InsertCakeRequest

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "UpdateCake",
		}).Errorln("[CakeHandler] Unable to decode request body")

		server.RenderError(w, http.StatusBadRequest, err, time.Now())
		return
	}

	resp, err := c.cakeService.UpdateCake(id, updateReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "UpdateCake",
		}).Errorln("[CakeHandler] Failed to update cake")

		server.RenderError(w, http.StatusInternalServerError, err, time.Now())
		return
	}
	server.RenderResponse(w, http.StatusOK, resp, time.Now())
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

		server.RenderError(w, http.StatusInternalServerError, err, time.Now())
		return
	}
	resp := types.CakeResponse{
		ID: id,
	}
	server.RenderResponse(w, http.StatusOK, resp, time.Now())
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

		server.RenderError(w, http.StatusInternalServerError, err, time.Now())
		return
	}
	server.RenderResponse(w, http.StatusOK, resp, time.Now())
}
