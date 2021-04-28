package adapter

import (
	"net/http"

	"com.damonkelley/shortcut/internal/shortcut"
)

func HTTPShortCutAdapter(db shortcut.Lookup) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		redirectTo, err := db.Lookup(r.URL.Path)

		if err != nil {
			http.NotFound(rw, r)
		} else {
			http.Redirect(rw, r, redirectTo.String(), 301)
		}
	})
}
