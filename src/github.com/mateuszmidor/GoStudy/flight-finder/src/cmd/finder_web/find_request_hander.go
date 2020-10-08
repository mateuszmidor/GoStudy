package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/application"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv"
)

func newFindRequestHandler() func(*gin.Context) {
	repo := csv.NewFlightsDataRepoCSV("../../../data/")
	finder := application.NewConnectionFinder(repo)

	return func(c *gin.Context) {
		from := getFromAirportCode(c.Request)
		to := getToAirportCode(c.Request)
		maxSegmentCount := getMaxSegmentsCount(c.Request)

		buff := &bytes.Buffer{}
		pathRenderer := application.NewPathRendererAsJSON(buff)
		err := finder.Find(from, to, maxSegmentCount, pathRenderer)

		// FIND ERROR
		if err != nil {
			log.Printf("%s -> %s: ERROR: %v\n", from, to, err)
			c.AbortWithStatusJSON(http.StatusBadRequest, errorToJSON(err))
			return
		}

		// FIND OK
		log.Printf("%s -> %s: OK\n", from, to)
		c.Render(
			http.StatusOK,
			render.Data{
				ContentType: "application/json",
				Data:        buff.Bytes(),
			})
	}

}

func getFromAirportCode(r *http.Request) string {
	return strings.ToUpper(r.FormValue("from"))
}

func getToAirportCode(r *http.Request) string {
	return strings.ToUpper(r.FormValue("to"))
}

func getMaxSegmentsCount(r *http.Request) int {
	count, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))
	return count
}

func errorToJSON(err error) interface{} {
	type ErrorJSON struct {
		Error string `json:"error"`
	}

	return ErrorJSON{Error: err.Error()}
}
