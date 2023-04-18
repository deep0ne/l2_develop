package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(logger *logrus.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/create_event":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("Создание встречи для юзера")
		case "/update_event":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("Обновляем время встречи для юзера")
		case "/delete_event":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Warn("Удаляем встречу у юзера")
		case "/events_for_day", "/events_for_week", "/events_for_month":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("Получаем информацию о встречах юзера")
		}
		next(w, r)
	}
}
