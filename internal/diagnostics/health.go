package diagnostics

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReadinessHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Readiness probe")

		w.WriteHeader(http.StatusOK)
	}
}


func LivenessHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Liveness probe")

		w.WriteHeader(http.StatusOK)
	}
}

