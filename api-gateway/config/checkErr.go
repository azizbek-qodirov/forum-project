package config

import (
	"api-gateway/config/logger"
)

type ErrorManager struct {
	logger *logger.Logger
}

func NewErrorManager(logger *logger.Logger) *ErrorManager {
	return &ErrorManager{logger: logger}
}

func (e *ErrorManager) CheckErr(err error, line int) {
	if err != nil {
		e.logger.ERROR.Panicln(err.Error() + string(line))
	}
}
