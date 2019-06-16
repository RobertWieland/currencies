package main

import (
	"Currency/currency"
	"fmt"
)

func main() {

	date := "2019-06-16"
	baseCurrency := "USD"
	requestedCurrency := "JPY"

	currencyRate := currency.GetExchangeRate(date, baseCurrency, requestedCurrency)
	fmt.Println("currency rate from func:", currencyRate)
}
