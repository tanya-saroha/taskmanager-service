package api

import (
	"encoding/json"
	"net/http"

	"github.com/tanya-saroha/taskmanager-service/app"
)

func Error(status int, response interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		app.GetLogger().Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func Success(status int, response interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		app.GetLogger().Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
