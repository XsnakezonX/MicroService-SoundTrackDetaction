package resources

import (
	"encoding/json"
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func searchTrack(w http.ResponseWriter, r *http.Request) {
	// println("test2")
	t := map[string]interface{}{}
	// println("test3")
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {

		if audio_file, ok := t["Audio"].(string); ok {
			// print("audio_file:")
			// print(audio_file)
			// print("audio_file:")
			if audio_file == "" {
				// panic(err) // wav file not found in DB
				// println("base64 wav file not exist locally or passed properly")
				w.WriteHeader(http.StatusNotFound) /* 404 */
				return
			}
			
			if results, err := service.Service(audio_file); err == nil {
				// if id is empty, then the wav file is not recognised by audd, return 404
				if results == "emptySong" {
					w.WriteHeader(http.StatusNotFound) /* 404 */
					return
				}

				u := map[string]interface{}{"Id": results}
				w.WriteHeader(http.StatusOK) /* 200 */
				json.NewEncoder(w).Encode(u)
				return
			} else {
				
				w.WriteHeader(http.StatusInternalServerError) /* 500 */
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest) /* 400 */
			return
		}
	}
	// println("Bad Request")
	w.WriteHeader(http.StatusBadRequest) /* 400 */
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* search audio */
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
