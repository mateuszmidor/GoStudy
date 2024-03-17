package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	bitcoinPriceInUsd := getBitcoinPriceInCurrency("USD")
	printAssetPrice(bitcoinPriceInUsd, "USD")

	bitcoinPriceInEur := getBitcoinPriceInCurrency("EUR")
	printAssetPrice(bitcoinPriceInEur, "EUR")

	bitcoinPriceInPln := getBitcoinPriceInCurrency("PLN")
	printAssetPrice(bitcoinPriceInPln, "PLN")
}

func printAssetPrice(price float64, currency string) {
	fmt.Printf("Bitcoin price in %s: %.2f\n", currency, price)
}

func getBitcoinPriceInCurrency(currency string) float64 {
	// Define the URL for the CoinGecko API
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP GET request:", err)
		return 0 // Return 0 or another default value in case of an error
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0 // Return 0 or another default value in case of an error
	}

	type Response struct {
		BPI struct {
			USD struct {
				RateFloat float64 `json:"rate_float"`
			} `json:"USD"`
			GBP struct {
				RateFloat float64 `json:"rate_float"`
			} `json:"GBP"`
			EUR struct {
				RateFloat float64 `json:"rate_float"`
			} `json:"EUR"`
		} `json:"bpi"`
	}

	// Unmarshal the JSON response into the struct
	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return 0 // Return 0 or another default value in case of an error
	}

	var price float64
	switch currency {
	case "USD":
		price = data.BPI.USD.RateFloat
	case "GBP":
		price = data.BPI.GBP.RateFloat
	case "EUR":
		price = data.BPI.EUR.RateFloat
	case "PLN":
		usdToPlnRate := getUSDRateInPLN()
		price = data.BPI.USD.RateFloat * usdToPlnRate
	default:
		fmt.Println("Unsupported currency:", currency)
		return 0 // Return 0 or another default value in case of an error
	}

	return price
}

// here is URL and response from API that returns USD rate in PLN for given date:
// URL: http://api.nbp.pl/api/exchangerates/rates/a/USD/2023-08-14?format=json
// Response: {"table":"A","currency":"dolar amerykański","code":"USD","rates":[{"no":"156/A/NBP/2023","effectiveDate":"2023-08-14","mid":4.0525}]}
func getUSDRateInPLN() float64 {
	type ExchangeRateResponse struct {
		Rates []struct {
			Mid float64 `json:"mid"`
		} `json:"rates"`
	}

	// Get today's date in the required format
	today := time.Now()

	var resp *http.Response
	var err error

	for i := 0; i < 3; i++ {
		formattedDate := today.Format("2006-01-02")

		url := fmt.Sprintf("http://api.nbp.pl/api/exchangerates/rates/a/USD/%s?format=json", formattedDate)

		// Make the HTTP GET request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("NewRequest failed: " + err.Error())
			return 0
		}

		req.Header.Set("User-Agent", "golang")
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("GET failed: " + err.Error())
			return 0
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			today = today.AddDate(0, 0, -1)
			continue
		}
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll failed: " + err.Error())
		return 0
	}

	// Unmarshal the JSON response into the struct
	var data ExchangeRateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Unmarshall failed: " + err.Error())
		fmt.Println(string(body))
		return 0
	}

	// Return the USD rate in PLN
	fmt.Println("currency exchange rate from:", today.Format("2006-01-02"))
	return data.Rates[0].Mid
}
