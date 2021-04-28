package adapter

import (
	"encoding/json"
	"net/http"

	"com.damonkelley/shortcut/internal/api"
	"com.damonkelley/shortcut/internal/shortcut"
)

func HTTPAPIAdapter(db shortcut.ReadWrite) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := api.NewGraphQL(db).Execute(r.URL.Query().Get("query"))

		json.NewEncoder(w).Encode(result)
	})
}
