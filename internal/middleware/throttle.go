package middleware

import (
	"context"
	"fmt"
	"net/http"
	"throttle/internal/storage"

	"github.com/rs/zerolog"
)

const throttleBound = 3

func Throttle(next http.HandlerFunc, log zerolog.Logger, cache storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("User-Id")

		num, err := cache.Get(context.Background(), fmt.Sprintf("user_%s", userId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if num >= throttleBound {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}
