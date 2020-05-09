package main

import (
	"encoding/json"
	"net/http"
)

type shortURL struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

// MapHandler func
func MapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		if dest, ok := pathToURLs[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

//JSONHandler func
func JSONHandler(jsonString string, fallback http.Handler) (http.HandlerFunc, error) {

	//PARSING JSON
	parsedJSON, err := parseJSON([]byte(jsonString))

	if err != nil {
		return nil, err
	}

	//BUILDING MAP
	pathToURLs := buildMap(parsedJSON)

	return MapHandler(pathToURLs, fallback), nil

}

func parseJSON(data []byte) ([]shortURL, error) {
	var pathURL []shortURL

	err := json.Unmarshal(data, &pathURL)

	if err != nil {
		return nil, err
	}

	return pathURL, nil
}

func buildMap(parsedJSON []shortURL) map[string]string {
	pathToURLs := make(map[string]string)

	for _, pu := range parsedJSON {
		pathToURLs[pu.Path] = pu.URL
	}

	return pathToURLs
}
