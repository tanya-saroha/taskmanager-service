package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanya-saroha/taskmanager-service/api"
	"github.com/tanya-saroha/taskmanager-service/config"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (taskRouter *mux.Router) {
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	fmt.Println(v1)

	router := mux.NewRouter()
	taskRouter = router.PathPrefix("/tasks").Subrouter()

	// ping
	taskRouter.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// //docs
	// docStrip := http.StripPrefix("/profiles/docs/", http.FileServer(http.Dir("./swaggerui/")))
	// profileRouter.PathPrefix("/docs/").Handler(docStrip)
	//config
	// taskRouter.HandleFunc("/config", api.Config()).Methods(http.MethodGet).Headers(versionHeader, v1)
	return
}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	api.Success(http.StatusOK, api.Response{Message: "pong"}, w)
}
