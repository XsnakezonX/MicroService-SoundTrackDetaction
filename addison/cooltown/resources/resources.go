package resources

import (
	"cooltown/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func coolTown(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{}
	var id string
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if audio_file, ok := t["Audio"].(string); ok {
			if audio_file == "" {
				// panic(err) // wav file not found in DB
				// println("Search microservice: Invalid AUDIO file Input")
				w.WriteHeader(http.StatusNotFound) /* 404 */
				return
			}
			if result, err := service.Service(audio_file); err == nil {
				id = result

			} else {
				if err != nil {
					// status code 500
					// println("Search microservice: Internal Server Error, track maybe unrecognisable")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

			}
		} else {
			// status code 400
			// println("Cooltown microservice: Bad Request")
			w.WriteHeader(http.StatusBadRequest) /* 400 */
			return
		}
	} else {
		// status code 400
		// println("Cooltown microservice: Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// read track from DB

	if results, err := service.ReadService(id); err == nil {
		if results == "" {
			// status code 404
			w.WriteHeader(http.StatusNotFound)
			return
		}

		u := map[string]interface{}{"Audio": results}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
		return
	} else {
		if err != nil {
			// status code 500
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* comunicates to other miroservices */
	r.HandleFunc("/cooltown", coolTown).Methods("POST")
	return r
}
