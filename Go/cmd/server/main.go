package main

import(
	"log"
	"net/http"
	// "fmt"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"

	controllers "github.com/jayschoen/iWant-slack-bot/internal"
)

func get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, has_id := vars["id"]

	if !has_id {

		json.NewEncoder(w).Encode(controllers.Get_all_wants())

	} else {

		id, err := strconv.Atoi(id)
		if err != nil || id < 1 {
			panic( err.Error() )
		}

		json.NewEncoder(w).Encode(controllers.Get_want_by_id(id))

	}
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write( []byte( `{"message": "POST"}` ) )
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write( []byte( `{"message": "PUT"}` ) )
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write( []byte( `{"message": "DELETE"}` ) )
}

/*
func notFound(w http.ResponseWriter, r *http.Request) {
}
*/

func executeTests(w http.ResponseWriter, r *http.Request) {

	test := `{"message": "yay tests"}`

	controllers.Tests()

	json.NewEncoder(w).Encode(test)
}

func main() {

	controllers.OpenDatabase()

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/get-wants/", get)
	r.HandleFunc("/get-wants/{id}", get)
	r.HandleFunc("/create-want", post)
	r.HandleFunc("/update-want/{id}", put)
	r.HandleFunc("/delete-want/{id}", delete)
	// r.HandleFunc("/", notFound)

	r.HandleFunc("/tests", executeTests)

	log.Fatal(http.ListenAndServe(":8080", r) )
}