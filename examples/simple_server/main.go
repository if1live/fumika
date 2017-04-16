package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/if1live/fumika"
)

/*
sample url
/aladin/9788926790403
/yes24/9788926790403
*/

func handleSearchAPI(w http.ResponseWriter, r *http.Request, path string, api fumika.SearchAPI) {
	code := r.URL.Path[len(path):]
	isbn, ok := fumika.SanitizeISBN(code)
	if !ok {
		http.Error(w, "{}", http.StatusBadRequest)
		return
	}
	result, err := api.SearchISBN(isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.Title == "" {
		http.Error(w, string(b), http.StatusNotFound)
		return
	}
	w.Write(b)
}

func handlerAladin(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	api := fumika.NewAladin(&client)
	handleSearchAPI(w, r, "/aladin/", api)
}

func handlerYes24(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	api := fumika.NewYes24(&client)
	handleSearchAPI(w, r, "/yes24/", api)
}

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "server port")
}
func main() {
	flag.Parse()

	http.HandleFunc("/aladin/", handlerAladin)
	http.HandleFunc("/yes24/", handlerYes24)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
