package goods

import "net/http"

func handleGetAllItems(w http.ResponseWriter, r *http.Request) {

}

func handleAddItem(w http.ResponseWriter, r *http.Request) {

}

func Routers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/itmes", handleGetAllItems)
	mux.HandleFunc("/itmes/add", handleAddItem)
	return mux
}
