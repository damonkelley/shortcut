package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"com.damonkelley/linkshortener/internal/database"
	"com.damonkelley/linkshortener/internal/graphql"
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

type ConstantDatabase struct{}

func (database *ConstantDatabase) Lookup(key string) (*url.URL, error) {
	return url.Parse("http://subdomain.example.com")
}

func RedirectMe(db database.Database) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		redirectTo, err := db.Lookup(r.URL.Path)

		if err != nil {
			http.NotFound(rw, r)
		} else {
			http.Redirect(rw, r, redirectTo.String(), 301)
		}
	})
}

func Welcome() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Welcome")
	})
}

func Api(db database.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := graphql.NewGraphQL(db).Execute(r.URL.Query().Get("query"))

		json.NewEncoder(w).Encode(result)
	})
}
func main() {
	db := &ConstantDatabase{}
	http.Handle("/welcome", WithLogging(Welcome()))
	http.Handle("/api", WithLogging(Api(db)))
	http.Handle("/", WithLogging(RedirectMe(db)))

	http.ListenAndServe(":8000", nil)
}
