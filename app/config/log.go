package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	logrus "github.com/sirupsen/logrus"
)

var logger = SetupLog()

// SetupLog - return logrus.
func SetupLog() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	return log
}
