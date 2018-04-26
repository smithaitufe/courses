package services

import (
	"os"

	"github.com/op/go-logging"
	"github.com/smithaitufe/courses/context"
)

func NewLogger(config *context.Config) *logging.Logger {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(config.Logger.Format)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.INFO, "")
	if config.Logger.DebugMode {
		backendLeveled.SetLevel(logging.DEBUG, "")
	}

	logging.SetBackend(backendLeveled)
	logger := logging.MustGetLogger(config.App.Name)
	return logger
}
