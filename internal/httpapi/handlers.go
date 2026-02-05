package httpapi

import (
	"net/http"
	"riffwire/internal/store"
	"strconv"
	"strings"
)

func ItemsByIDHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/items/")
		if idStr == "" {
			writeErrorJSON(w, http.StatusBadRequest, "id is empty")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeErrorJSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		found, ok := s.GetItemByID(id)
		if !ok {
			writeErrorJSON(w, http.StatusNotFound, "item not found")
			return
		}

		writeJSON(w, http.StatusOK, found)
	}
}

func ItemsHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		writeJSON(w, http.StatusOK, s.ListItems())
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	writeJSON(w, http.StatusOK, map[string]string{"message": "Hello"})

}
