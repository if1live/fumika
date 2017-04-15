package main

import (
	"fmt"

	"github.com/if1live/fumika"
)

func main() {
	mainAladin()
	mainYes24()
}

func mainYes24() {
	api := fumika.NewYes24()
	isbn := "9788926790403"
	result, err := api.Search(isbn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Yes24 API")
	printSearchResult(&result)
}

func mainAladin() {
	api := fumika.NewAladin()
	isbn := "9788926790403"
	result, err := api.Search(isbn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Aladin API")
	printSearchResult(&result)
}

func printSearchResult(r *fumika.SearchResult) {
	fmt.Printf("Title : %s\n", r.Title)
	fmt.Printf("UnitPrice : %d\n", r.UnitPrice)
	fmt.Printf("PriceBest : %d\n", r.PriceBest)
	fmt.Printf("PriceGood : %d\n", r.PriceGood)
	fmt.Printf("PriceNormal : %d\n", r.PriceNormal)
}
