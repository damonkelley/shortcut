package main

import (
	"net/http"
	"net/url"

	apiadapter "com.damonkelley/shortcut/internal/api/adapter"
	"com.damonkelley/shortcut/internal/infrastructure"
	"com.damonkelley/shortcut/internal/shortcut"
	"com.damonkelley/shortcut/internal/shortcut/adapter"
)

type ConstantDatabase struct{}

func (database *ConstantDatabase) Lookup(key string) (*url.URL, error) {
	return url.Parse("http://tyleemarsh.com")
}

func main() {
	shortcuts := shortcut.NewShortcuts(
		shortcut.NewRandomKeyGenerator(
			shortcut.DefaultRandomKeyConfig(),
		),
	)

	http.Handle("/api", infrastructure.WithLogging(apiadapter.HTTPAPIAdapter(shortcuts)))
	http.Handle("/", infrastructure.WithLogging(adapter.HTTPShortCutAdapter(shortcuts)))

	http.ListenAndServe(":8000", nil)
}
