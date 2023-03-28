package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	localSearch = "http://localhost:3001/search"
	localTrack  = "http://localhost:3000/tracks/"
)

func Service(audio_file string) (string, error) {
	client := &http.Client{}
	m, b := map[string]string{"Audio": audio_file}, new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	if req, err := http.NewRequest("POST", localSearch, b); err == nil {
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == http.StatusOK {
				t := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&t); err == nil {
					if id, err := idOf(t); err == nil {
						return replaceSpaces(id), nil
					}
				}
			} else if rsp.StatusCode == http.StatusNotFound {
				return "", nil
			}
		}
	}
	return "", errors.New("Service")
}

func idOf(t map[string]interface{}) (string, error) {
	if id, ok := t["Id"].(string); ok {
		return id, nil
	}
	return "", errors.New("idOf")
}

// a functon that replace all spaces of a string parameter with "+"
func replaceSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "+")
}

// return the track audio with the given id
func ReadService(id string) (string, error) {
	client := &http.Client{}
	if req, err := http.NewRequest("GET", localTrack+id, nil); err == nil {
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == http.StatusOK {
				t := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&t); err == nil {
					if audio, err := audioOf(t); err == nil {
						return audio, nil
					}
				}
			} else if rsp.StatusCode == http.StatusNotFound {
				// println("Track Microservice: track not found in DB, create the track first.")
				return "", nil
			}
		}
	}
	return "", errors.New("ReadServiceice")
}

func audioOf(t map[string]interface{}) (string, error) {

	if audio, ok := t["Audio"].(string); ok {
		return audio, nil
	}

	return "", errors.New("audioOf")
}
