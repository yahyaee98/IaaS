package api

import (
	"context"
	"encoding/json"
	"iaas/internal/data"
	"iaas/internal/log"
	"iaas/internal/repository"
	"net/http"
)

// GetItemsHandler finds proper results and returns back to the client
func GetItemsHandler(rep repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			respond(w, http.StatusBadRequest, data.ErrorResponse{Error: "malformed form input"})
			return
		}

		search := r.FormValue("search")
		if len(search) < 3 {
			respond(w, http.StatusBadRequest, data.ErrorResponse{Error: "search should be at least 3 characters"})
			return
		}

		items, err := rep.GetItems(search)
		if err != nil {
			respond(w, http.StatusInternalServerError, data.ErrorResponse{Error: "server has failed to fetch items"})
		}

		respond(w, http.StatusOK, data.Response{Items: items})
	}
}

// GetHealthHandler handles requests to /health. It considers itself healthy all the times, except the moments that
// we know we have to end everything to shutdown gracefully.
func GetHealthHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	readyState := true

	go func() {
		<-ctx.Done()
		readyState = false
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		if readyState {
			respond(w, http.StatusOK, nil)
			return
		}

		respond(w, http.StatusInternalServerError, nil)
	}
}

func respond(w http.ResponseWriter, code int, body interface{}) {
	byteBody, _ := json.Marshal(body)
	w.WriteHeader(code)
	if _, err := w.Write(byteBody); err != nil {
		log.Errorw("http server has failed to write response",
			"error", err,
		)
	}
}
