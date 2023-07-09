package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"testing"

	. "github.com/Eun/go-hit"
)

func Test_Simple(t *testing.T) {
	// Expected response:
	// {
	// 	"table":"A",
	// 	"currency":"euro",
	// 	"code":"EUR",
	// 	"rates":[
	// 	   {
	// 		  "no":"129/A/NBP/2023",
	// 		  "effectiveDate":"2023-07-06",
	// 		  "mid":4.4754
	// 	   }
	// 	]
	//  }

	var returnedDate string
	Test(t,
		Description("Get EUR/PLN exchange rate for 2023-07-06 from api.nbp.pl"),
		Get("http://api.nbp.pl/api/exchangerates/rates/A/EUR/2023-07-06?format=json"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Headers("Content-Type").Contains("application/json; charset=utf-8"),
		Expect().Body().JSON().JQ(".table").Equal("A"),
		Expect().Body().JSON().JQ(".currency").Equal("euro"),
		Expect().Body().JSON().JQ(".code").Equal("EUR"),
		Expect().Body().JSON().JQ(".rates").Len().Equal(1),
		Expect().Body().JSON().JQ(".rates[0].mid").Equal(4.4754),
		Store().Response().Body().JSON().JQ(".rates[0].effectiveDate").In(&returnedDate), // store received response in variable,
	)
}

func Test_CustomHttpClient(t *testing.T) {
	// Expected response:
	// {
	// 	"table":"A",
	// 	"currency":"euro",
	// 	"code":"EUR",
	// 	"rates":[
	// 	   {
	// 		  "no":"129/A/NBP/2023",
	// 		  "effectiveDate":"2023-07-06",
	// 		  "mid":4.4754
	// 	   }
	// 	]
	//  }

	// prepare support for cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	// prepare http client
	client := &http.Client{
		Jar: jar,
	}

	steps := []IStep{
		HTTPClient(client),
		BaseURL("http://api.nbp.pl"),
	}
	Test(t, append(steps,
		Description("Get EUR/PLN exchange rate for 2023-07-06 from api.nbp.pl, custom HTTP client with cookies suppoprt"),
		Get("/api/exchangerates/rates/A/EUR/2023-07-06?format=json"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Headers("Content-Type").Contains("application/json; charset=utf-8"),
		Expect().Body().JSON().JQ(".table").Equal("A"),
	)...)
}
