package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func extractTime(html string) string {
	REGEXP := "[0-9][0-9]:[0-9][0-9]:[0-9][0-9]"
	r, _ := regexp.Compile(REGEXP)
	return r.FindString(html)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
func main() {
	TIME_WWW := "https://www.worldtimeserver.com/current_time_in_PL.aspx"
	resp, err := http.Get(TIME_WWW)
	exitOnError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	exitOnError(err)

	html := string(body)
	time := extractTime(html)
	fmt.Printf("Time in Poland: %s\n", time)
}
