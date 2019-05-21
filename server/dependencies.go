package server

import (
	"github.com/tanya-saroha/taskmanager-service/app"
	"github.com/tanya-saroha/taskmanager-service/db"
)

type dependencies struct {
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB, logger)

	return dependencies{
	}, nil
}
