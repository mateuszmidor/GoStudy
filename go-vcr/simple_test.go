package main

import (
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
	"io"
	"net/http"
	"strings"
	"testing"
)

func Test_getEurExchangeRatesForDate_returnRates(t *testing.T) {
	// given
	const url = "http://api.nbp.pl/api/exchangerates/rates/A/EUR/2023-07-06?format=json"
	const cachedResponsePath = "fixtures/api-nbp-pl"
	const expectedRate = "4.4754"

	rec, err := recorder.New(cachedResponsePath)
	if err != nil {
		t.Fatal(err)
	}
	defer rec.Stop() // Make sure recorder is stopped once done with it
	if rec.Mode() != recorder.ModeRecordOnce {
		t.Fatal("Recorder should be in ModeRecordOnce")
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("User-Agent", "golang") // api.nbp.pl requires User-Agent to be set, or 403 "Request forbidden by administrative rules" is returned
	client := rec.GetDefaultClient()

	// when
	resp, err := client.Do(req)

	// then
	if err != nil {
		t.Fatalf("Failed to get url %s: %s", url, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	bodyContent := string(body)
	if !strings.Contains(bodyContent, expectedRate) {
		t.Errorf("EUR/PLN rate %s not found in response: %s", expectedRate, bodyContent)
	}
}
