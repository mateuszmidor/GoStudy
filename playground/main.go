package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	priceByCurrency, ok := getBitcoinPrices()
	if !ok {
		return 0
	}

	if currency == "PLN" {
		if pln, ok := priceByCurrency["pln"]; ok {
			return pln
		}
		usdToPlnRate := getUSDRateInPLN()
		return priceByCurrency["usd"] * usdToPlnRate
	}

	price, ok := priceByCurrency[strings.ToLower(currency)]
	if !ok {
		fmt.Println("Unsupported currency:", currency)
		return 0
	}
	return price
}

func getBitcoinPrices() (map[string]float64, bool) {
	u, err := url.Parse("https://api.coingecko.com/api/v3/simple/price")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, false
	}
	q := u.Query()
	q.Set("ids", "bitcoin")
	q.Set("vs_currencies", "usd,eur,gbp,pln")
	u.RawQuery = q.Encode()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(u.String())
	if err != nil {
		fmt.Println("Error making HTTP GET request:", err)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		b, _ := io.ReadAll(resp.Body)
		fmt.Printf("Unexpected HTTP status: %s\n%s\n", resp.Status, string(b))
		return nil, false
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, false
	}

	// Example:
	// {"bitcoin":{"usd":123,"eur":456,"gbp":789,"pln":999}}
	var data map[string]map[string]float64
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, false
	}
	m, ok := data["bitcoin"]
	if !ok || len(m) == 0 {
		fmt.Println("Unexpected response:", string(body))
		return nil, false
	}
	return m, true
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
