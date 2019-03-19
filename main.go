package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type request struct {
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Environ []string    `json:"environ"`
	Body    []byte      `json:"body"`
}

const EnvPrefix = "ECHO_"

func environ() []string {
	env := make([]string, 0)
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, EnvPrefix) {
			env = append(env, strings.TrimPrefix(e, EnvPrefix))
		}
	}
	return env
}

func handle(rw http.ResponseWriter, r *http.Request) {
	var err error
	rr := &request{}
	rr.Method = r.Method
	rr.Headers = r.Header
	rr.URL = r.URL.String()
	rr.Environ = environ()
	rr.Body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rrb, err := json.Marshal(rr)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(rrb)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}
