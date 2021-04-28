package main

import (
	"net/http"
	"net/url"

	api "com.damonkelley/shortcut/internal/api/adapter"
	"com.damonkelley/shortcut/internal/infrastructure"
	shortcut "com.damonkelley/shortcut/internal/shortcut/adapter"
)

type ConstantDatabase struct{}

func (database *ConstantDatabase) Lookup(key string) (*url.URL, error) {
	return url.Parse("http://tyleemarsh.com")
}

func main() {
	db := &ConstantDatabase{}
	http.Handle("/api", infrastructure.WithLogging(api.HTTPAPIAdapter(db)))
	http.Handle("/", infrastructure.WithLogging(shortcut.HTTPShortCutAdapter(db)))

	http.ListenAndServe(":8000", nil)
}
