package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const port = ":5500"

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/cp{cp}", getDireccion)
	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))

}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root page"))
}

func getDireccion(w http.ResponseWriter, r *http.Request) {
	cp := mux.Vars(r)["cp"]
	//resp, err := http.Get("https://api.copomex.com/query/info_cp/" + cp + "?type=simplified&token=pruebas")
	resp, err := http.Get("https://api.copomex.com/query/info_cp/" + cp + "?type=simplified&token=c5b05a80-9d5a-4752-92c8-ee85ba8f83be") // Token real limitado a 50 pruebas
	if err != nil {
		log.Fatalln(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		} else {
			//We Read the response body on the line below.
			w.Header().Set("content-type", "application/json")
			w.Write(body)
		}
	}
}
