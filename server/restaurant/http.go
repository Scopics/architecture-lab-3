package restaurant

import (
	"log"
	"net/http"

	"github.com/Scopics/architecture-lab-3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleFullMenu(store, rw)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleFullMenu(store *Store, rw http.ResponseWriter) {
	res, err := store.GetMenu()
	if err != nil {
		log.Printf("Error in a database query: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
