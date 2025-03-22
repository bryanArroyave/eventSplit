package ports

import "github.com/bryanArroyave/eventsplit/common/logger/dtos"

type ILogger interface {
	Info(message string, fields ...*dtos.LoggerFieldsDTO)
	Error(message string, fields ...*dtos.LoggerFieldsDTO)
	Warn(message string, fields ...*dtos.LoggerFieldsDTO)
}
