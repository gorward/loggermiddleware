package loggermiddleware

import (
	"github.com/gorward/logger"
	"net/http"
	"time"
)

func LoggerMiddleware(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var log *logger.Logger

		log = logger.New(logger.Config{
			Level:  logger.Access,
			Access: "access.log",
		})

		start := time.Now()

		defer func() {
			log.Access(start, w, r)
		}()

		h.ServeHTTP(w, r)
	})

}
