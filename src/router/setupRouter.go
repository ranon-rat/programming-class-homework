package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/proyectoClaseProgramacion/src/controllers"
)

func SetupRouter() error {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.String())
		if _, ok := r.URL.Query()["test"]; !ok {

			http.Redirect(rw, r, "/?test", http.StatusMovedPermanently)
		}
		http.ServeFile(rw, r, "view/index.html")
	})

	r.HandleFunc("/upload/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "view/upload.html")
	})
	r.HandleFunc("/purchases/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "view/buy.html")
	})

	r.HandleFunc("/api/", controllers.SendApi)
	r.HandleFunc(`/public/{file:[\w\W\/]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	log.Println("running on http://localhost:8080")

	return http.ListenAndServe(":8080", r)
}
