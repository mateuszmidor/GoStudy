package main

import ("fmt"
		"os"
		"regexp"
		"io/ioutil"
		"net/http"
)

func extract_time(html string) string {
	REGEXP := "[0-9][0-9]:[0-9][0-9]:[0-9][0-9]"
	r, _ := regexp.Compile(REGEXP)
	return r.FindString(html)
}

func exit_on_error(err error) {
	if (err != nil) {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
func main() {
	TIME_WWW := "https://www.worldtimeserver.com/current_time_in_PL.aspx"
	resp, err := http.Get(TIME_WWW)
	exit_on_error(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	exit_on_error(err)

	html := string(body)
	time := extract_time(html)
	fmt.Printf("Time in Poland: %s\n", time)
}
