package app

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tanya-saroha/taskmanager-service/config"
	"go.uber.org/zap"
)

var (
	db     *sqlx.DB
	logger *zap.SugaredLogger
)

func Init() {
	InitLogger()

	err := initDB()
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = zapLogger.Sugar()
}

func initDB() (err error) {
	dbConfig := config.Database()

	db, err = sqlx.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		logger.Errorw("Could not ping db", "msg", err.Error())
		return
	}

	db.SetMaxIdleConns(dbConfig.MaxPoolSize())
	db.SetMaxOpenConns(dbConfig.MaxOpenConns())
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)

	return
}

func GetDB() *sqlx.DB {
	return db
}

func GetLogger() *zap.SugaredLogger {
	return logger
}

func Close() {
	db.Close()
}
