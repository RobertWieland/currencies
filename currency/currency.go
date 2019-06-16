package currency

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Currencies struct
// All of these must be Capital Letters in order for the Unmarshal to work
type Currencies struct {
	Base  string
	Rates map[string]float64
	Date  string
}

// GetExchangeRate returns the exange rate for a date, baseCurrence, and a requestedCurrency
func GetExchangeRate(date string, baseCurrency string, requestedCurrency string) float64 {

	currencyURL := "https://api.exchangeratesapi.io/"

	resp, err := http.Get(currencyURL + date + "?base=" + baseCurrency)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	resp.Body.Close()

	// If "var testAnything interface{}" used it will turn the entire json into a struct
	// so, you do not need to know what format the response body will be in. If you do know
	// the format of the response body then create your own struct.
	currencies := &Currencies{}

	// Takes the string "body" and turns it into a slice of bytes.
	// Then unmarsals the byte slice and stores it at the address of Currencies
	json.Unmarshal([]byte(body), currencies)

	return currencies.Rates[requestedCurrency]
}
