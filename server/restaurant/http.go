package restaurant

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Scopics/architecture-lab-3/server/tools"
)

type ClientOrder struct {
	Table int          `json:"table"`
	Items []*OrderItem `json:"items"`
}

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetMenu(rw, store)
		} else if r.Method == "POST" {
			handleAddNewOrder(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleGetMenu(rw http.ResponseWriter, store *Store) {
	res, err := store.GetMenu()
	if err != nil {
		log.Printf("Error in a database query: %s", err)
		tools.SendJsonInternalError(rw)
		return
	}
	tools.SendJsonOk(rw, res)
}

func handleAddNewOrder(r *http.Request, rw http.ResponseWriter, store *Store) {
	var clientOrder *ClientOrder
	if err := json.NewDecoder(r.Body).Decode(&clientOrder); err != nil {
		log.Printf("Decoding json caused an error: %s", err)
		tools.SendJsonBadRequest(rw, "Unable to render JSON")
		return
	}

	resOrder, err := store.AddNewOrder(clientOrder.Table, clientOrder.Items)
	if err == nil {
		tools.SendJsonOk(rw, resOrder)
	} else {
		log.Printf("Error writing data to the database: %s", err)
		tools.SendJsonInternalError(rw)
	}
}
