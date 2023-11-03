package concrete

import (
	"cc-backend-root-console/app/boot/application/application/abstracts"
)

var instance abstracts.Application = nil

func App(args ...interface{}) interface{} {
	if instance == nil {
		instance = BuildAppInstance()
	}

	if len(args) == 1 {
		return instance.Get(args[0].(string))
	}
	return instance
}
