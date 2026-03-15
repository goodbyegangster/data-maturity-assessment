package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const htmxKey contextKey = "htmx"

// Htmx は HX-Request ヘッダーの有無を context に格納するミドルウェア。
func Htmx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), htmxKey, r.Header.Get("HX-Request") == "true")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// IsHtmx は context から htmx リクエストかどうかを返す。
func IsHtmx(r *http.Request) bool {
	v, _ := r.Context().Value(htmxKey).(bool)
	return v
}
