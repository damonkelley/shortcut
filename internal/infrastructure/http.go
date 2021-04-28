package infrastructure

import (
	"fmt"
	"net/http"
)

type ResponseRecorder struct {
	http.ResponseWriter
	status int
}

func (rw *ResponseRecorder) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func WithLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		responseRecorder := &ResponseRecorder{
			ResponseWriter: rw,
			status:         200,
		}

		h.ServeHTTP(responseRecorder, r)

		fmt.Printf("%s %s - %d\n", r.Method, r.URL.Path, responseRecorder.status)
	})
}
