package handler

import (
	"encoding/json"
	"net/http"
	"time"

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
		}).Errorln("[CakeHandler] Unable to decode request body")

		server.RenderError(w, http.StatusBadRequest, err, time.Now())
		return
	}

	resp, err := c.cakeService.InsertCake(insertReq)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"function": "InsertCake",
		}).Errorln("[CakeHandler] Failed to insert cake")

		server.RenderError(w, http.StatusBadRequest, err, time.Now())
		return
	}
	server.RenderResponse(w, http.StatusCreated, resp, time.Now())
}
