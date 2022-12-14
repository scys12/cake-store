package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type headerData struct {
	ProcessTime  int64  `json:"process_time_ms"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type response struct {
	Header       headerData  `json:"header"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data"`
}

func RenderResponse(w http.ResponseWriter, statusCode int, data interface{}, startTime time.Time) {
	res := response{
		Header: headerData{
			ProcessTime: time.Since(startTime).Milliseconds(),
		},
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	d, _ := json.Marshal(res)
	w.Write(d)
}

func RenderError(w http.ResponseWriter, statusCode int, err error, startTime time.Time) {
	res := response{
		Header: headerData{
			ProcessTime:  time.Since(startTime).Milliseconds(),
			ErrorMessage: err.Error(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	d, _ := json.Marshal(res)
	w.Write(d)
}
