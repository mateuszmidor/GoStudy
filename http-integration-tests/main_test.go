package main

import (
	"net/http"
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
	)
}

func Test_Example(t *testing.T) {
	// TODO
}
