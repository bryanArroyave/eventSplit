package app

import (
	"context"

	appdtos "github.com/bryanArroyave/eventsplit/common/app/dtos"
	"github.com/bryanArroyave/eventsplit/common/logger/adapter/singleton"
	"github.com/bryanArroyave/eventsplit/common/logger/ports"
	mongoadapter "github.com/bryanArroyave/eventsplit/common/mongo"
	"github.com/bryanArroyave/eventsplit/common/mongo/dtos"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	logger         ports.ILogger
	mongoInstances map[string]*mongo.Database
}

func NewApp(config *appdtos.LoggerConfigDTO) *App {
	app := &App{}
	app.initLogger(config)
	return app
}

func (a *App) initLogger(config *appdtos.LoggerConfigDTO) {
	singleton.InitLogger(config.LoggerType, config.ServiceName)
	a.logger = singleton.GetLogger()
}

func (a *App) AddMongoConnection(name string, config *dtos.MongoConnectionDTO) *App {
	if a.mongoInstances == nil {
		a.mongoInstances = make(map[string]*mongo.Database)
	}

	db, err := mongoadapter.InitMongoDB(context.TODO(), config)
	if err != nil {
		panic(err)
	}

	a.mongoInstances[name] = db
	return a

}

// GETTERS
func (a *App) GetLogger() ports.ILogger {
	return a.logger
}

func (a *App) GetMongoConnection(name string) *mongo.Database {
	return a.mongoInstances[name]
}
