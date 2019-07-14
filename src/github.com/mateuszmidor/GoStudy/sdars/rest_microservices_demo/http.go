package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func encodeBody(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func respond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		encodeBody(w, data)
	}
}

func respondErr(w http.ResponseWriter, status int, args ...interface{}) {
	respond(w, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func respondHTTPErr(w http.ResponseWriter, status int) {
	respondErr(w, status, http.StatusText(status))
}

func httpPut(url string, v interface{}) {
	var builder strings.Builder
	err := encodeBody(&builder, v)
	payload := builder.String()
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	if err != nil {
		log.Println(err)
		return
	}

	// request.SetBasicAuth("admin", "admin")
	request.ContentLength = int64(len(payload))
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	_ = response
	// else {
	// 	defer response.Body.Close()
	// 	contents, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
	// 	fmt.Println("   ", response.StatusCode)
	// 	hdr := response.Header
	// 	for key, value := range hdr {
	// 		fmt.Println("   ", key, ":", value)
	// 	}
	// 	fmt.Println(contents)
	// }
}
