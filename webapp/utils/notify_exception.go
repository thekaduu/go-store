package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func NotifyException(err error, request http.Request, clientIP string) {

	log.WithError(err).WithFields(log.Fields{
		"path":     request.URL.Path,
		"method":   request.Method,
		"clientIP": clientIP,
	}).Error()
}
