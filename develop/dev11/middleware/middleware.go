package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(logger *logrus.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Warn(err)
		}
		switch r.URL.Path {

		case "/create_event":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Infof("Создание встречи для юзера с ID=%s. Встреча на %s. Название встречи: %s", r.Form.Get("user_id"), r.Form.Get("date"), r.Form.Get("event_name"))

		case "/update_event":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Infof("Обновляем время встречи для юзера с ID=%s. Встреча %s перенесена на %s", r.Form.Get("user_id"), r.Form.Get("event_name"), r.Form.Get("date"))

		case "/delete_event":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Warnf("Удаляем встречу %s у юзера с ID=%s", r.Form.Get("event_name"), r.Form.Get("user_id"))

		case "/events_for_day", "/events_for_week", "/events_for_month":
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Infof("Получаем информацию о встречах юзера с ID=%s", r.Form.Get("user_id"))
		}
		next(w, r)
	}
}
