package main

import (
	"net/http"

	apiadapter "com.damonkelley/shortcut/internal/api/adapter"
	"com.damonkelley/shortcut/internal/infrastructure"
	"com.damonkelley/shortcut/internal/shortcut"
	"com.damonkelley/shortcut/internal/shortcut/adapter"
)

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
