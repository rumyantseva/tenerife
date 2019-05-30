package application

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HomeHandler(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Catch a home request")

		w.WriteHeader(http.StatusOK)
	}
}
