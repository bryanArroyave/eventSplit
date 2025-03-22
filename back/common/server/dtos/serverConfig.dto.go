package dtos

import (
	appadapter "github.com/bryanArroyave/eventsplit/common/app"
	"github.com/bryanArroyave/eventsplit/common/server/ports"
)

type APIRestServerConfigDTO struct {
	GlobalPrefix string
	Port         string
	App          *appadapter.App
}

type RouterHandler struct {
	Prefix string
	Router ports.IRouter
}
