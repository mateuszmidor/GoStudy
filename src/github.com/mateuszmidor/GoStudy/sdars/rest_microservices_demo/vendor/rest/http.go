package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func DecodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func EncodeBody(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func Respond(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	w.WriteHeader(httpStatusCode)
	if data != nil {
		EncodeBody(w, data)
	}
}

func RespondErr(w http.ResponseWriter, httpStatusCode int, args ...interface{}) {
	Respond(w, httpStatusCode, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func RespondHTTPErr(w http.ResponseWriter, httpStatusCode int) {
	RespondErr(w, httpStatusCode, http.StatusText(httpStatusCode))
}

func HttpPut(url string, v interface{}) error {
	var builder strings.Builder
	err := EncodeBody(&builder, v)
	if err != nil {
		return err
	}
	payload := builder.String()
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	if err != nil {
		return err
	}

	// request.SetBasicAuth("admin", "admin")
	request.ContentLength = int64(len(payload))
	_, err = client.Do(request)
	return err
}
