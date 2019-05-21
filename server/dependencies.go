package server

import (
	"fmt"
	"github.com/tanya-saroha/taskmanager-service/app"
	"github.com/tanya-saroha/taskmanager-service/db"
)

type dependencies struct {
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB, logger)
	fmt.Println(appDB, logger, dbStore)
	return dependencies{}, nil
}
