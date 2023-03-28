package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	api_url   = "https://api.audd.io/"
	// api_token = "" // insert api key here
	
)

func Service(audio_file string) (string, error) {
	data := url.Values{
		"audio":     {audio_file},
		"return":    {"apple_music,spotify"},
		"api_token": {api_token},
	}
	response, _ := http.PostForm(api_url, data)
	defer response.Body.Close()
	// print("response:")
	// print(response.StatusCode)
	if response.StatusCode != 200 {
		return "", errors.New("Service")
	}
	// body, _ := ioutil.ReadAll(response.Body)
	t := map[string]interface{}{}
	if err := json.NewDecoder(response.Body).Decode(&t); err == nil {
		
		if t["status"].(string) == "error" {
			// print("\n")
			// print(t["status"].(string))
			// print(t["result"])
			// print("\n")
			
			return "", errors.New("Service")
		}
		if trackid, err := trackidOf(t); err == nil {
			// print("printed")
			// return replaceSpaces(trackid), nil
			return trackid, nil
		} else {
			return "emptySong", nil
		}
	} else {
		return "", err
	}
}

// source: https://web.archive.org/web/20191201001550/https://docs.audd.io/

func trackidOf(t map[string]interface{}) (string, error) {

	if result, ok := t["result"].(map[string]interface{}); ok {
		if title, ok := result["title"].(string); ok {
			// print("decode")
			// print(title)
			// print("decode")
			return title, nil
		} else {
			print("no result field")
			return "", errors.New("trackidOf")
		}
	} // result field response is an empty array or null when the wav file is not recognised by audd
	return "", errors.New("trackidOf")
}
