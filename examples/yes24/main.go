package main

import (
	"fmt"
	"net/http"

	"github.com/if1live/fumika"
)

func main() {
	client := http.Client{}
	api := fumika.NewYes24(&client)
	isbn := "9788926790403"
	result, err := api.SearchISBN(isbn)
	if err != nil {
		panic(err)
	}

	fmt.Println("title: ", result.Title)
	fmt.Println("unit price: ", result.UnitPrice)
	fmt.Println("price best: ", result.PriceBest)
	fmt.Println("price good: ", result.PriceGood)
	fmt.Println("price normal: ", result.PriceNormal)
}
