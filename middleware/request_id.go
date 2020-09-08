package middleware

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"pkg.dsb.dev/requestid"
)

// RequestID is a middleware that reuses or creates a request id for each HTTP request. The
// id is also added to the request context using the requestid package.
func RequestID() mux.MiddlewareFunc {
	const key = "X-Request-ID"
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get(key)
			if id == "" {
				id = uuid.Must(uuid.NewV4()).String()
				r.Header.Set(key, id)
			}

			if span := opentracing.SpanFromContext(r.Context()); span != nil {
				span.SetTag("http.request_id", id)
			}

			ctx := requestid.ToContext(r.Context(), id)
			w.Header().Set(key, id)
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
