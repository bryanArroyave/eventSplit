package singleton

import (
	"github.com/bryanArroyave/eventsplit/common/logger/enums"
	"github.com/bryanArroyave/eventsplit/common/logger/factory"
	"github.com/bryanArroyave/eventsplit/common/logger/ports"
)

var (
	logger ports.ILogger
)

func InitLogger(loggerType enums.LoggerType, serviceName string) {
	logger = factory.NewLogger(loggerType, serviceName)
}

func GetLogger() ports.ILogger {
	return logger
}
