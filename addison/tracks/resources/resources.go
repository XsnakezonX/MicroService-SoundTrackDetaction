package resources

import (
	"encoding/json"
	"net/http"
	"tracks/repository"

	"github.com/gorilla/mux"
)

func updateTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t repository.Track

	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if id == t.Id {
			if n := repository.Update(t); n > 0 {
				w.WriteHeader(http.StatusNoContent) /* No Content */
			} else if n := repository.Insert(t); n > 0 {
				w.WriteHeader(http.StatusCreated) /* Created */
			} else {
				w.WriteHeader(http.StatusInternalServerError) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(http.StatusBadRequest) /* Bad Request */
		}
	} else {
		w.WriteHeader(http.StatusBadRequest) /* Bad Request */
	}
}

func readTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if t, n := repository.Read(id); n > 0 {
		d := repository.Track{Id: t.Id, Audio: t.Audio}
		w.WriteHeader(http.StatusOK) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if n == 0 {
		w.WriteHeader(http.StatusNotFound) /* Not Found */
	} else {
		w.WriteHeader(http.StatusInternalServerError) /* Internal Server Error */
	}
}

func listTrack(w http.ResponseWriter, r *http.Request) {
	if ts, n := repository.List(); n >= 0 {
		w.WriteHeader(http.StatusOK) /* OK */
		json.NewEncoder(w).Encode(ts)
		// } else if n == 0 {
		// 	w.WriteHeader(http.StatusNotFound) /* Not Found */
	} else {
		w.WriteHeader(http.StatusInternalServerError) /* Internal Server Error */
	}
}

func deleteTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if n := repository.Delete(id); n > 0 {
		w.WriteHeader(http.StatusNoContent) /* No Content */
	} else if n == 0 {
		w.WriteHeader(http.StatusNotFound) /* Not Found */
	} else {
		w.WriteHeader(http.StatusInternalServerError) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Creating Tracks or Store */
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	/* Listing Tracks */
	r.HandleFunc("/tracks", listTrack).Methods("GET")
	/* Reading Tracks */
	r.HandleFunc("/tracks/{id}", readTrack).Methods("GET")
	/* Deleting Tracks */
	r.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")
	return r
}
