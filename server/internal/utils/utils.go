package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response map[string]any

func WriteResult(w http.ResponseWriter, result interface{}, code int) error {
	w.WriteHeader(code)
	json, err := json.Marshal(result)
	if err != nil {
		return err
	}
	_, err = w.Write(json)
	if err != nil {
		return err
	}
	return nil
}

func ReadJson(r *http.Request, expected any) error {
	text, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(text, expected)
	if err != nil {
		return err
	}
	return nil
}

/*
extracts params from url

for example:
path "users/:id"
url "users/342"

returns 342
*/
func ParseParams(path string) func(url string) (map[string]string, error) {
	parsedPath := strings.Split(path, "/")
	routes := map[string]int{}
	for i := 0; i < len(parsedPath); i++ {
		elem := parsedPath[i]
		if elem != "" && elem[0] == ':' {
			routes[elem[1:]] = i
		}
	}
	return func(url string) (map[string]string, error) {
		parsedUrl := strings.Split(url, "/")
		if len(parsedUrl) != len(parsedPath) {
			return nil, fmt.Errorf("invalid url")
		}
		res := map[string]string{}
		for value, key := range routes {
			tra := parsedUrl[key]
			res[value] = tra
		}
		// fmt.Println(res)
		return res, nil
	}
}

func Error(w http.ResponseWriter, err string, status int) {
	WriteResult(w, Response{"err": err}, status)
}

func Result(w http.ResponseWriter, res any) {
	WriteResult(w, Response{"res": res}, 200)
}

func RequireFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
