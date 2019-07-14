package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		EncodeBody(w, data)
	}
}

func RespondErr(w http.ResponseWriter, status int, args ...interface{}) {
	Respond(w, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func RespondHTTPErr(w http.ResponseWriter, status int) {
	RespondErr(w, status, http.StatusText(status))
}

func HttpPut(url string, v interface{}) {
	var builder strings.Builder
	err := EncodeBody(&builder, v)
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
