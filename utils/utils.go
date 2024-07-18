package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// to parse an incoming JSON according to a pre-defined struct
func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err) // used to terminate program when unable to handle the error gracefully
	} else if err := json.Unmarshal([]byte(body), &x); err != nil {
		panic(err)
	}
}
