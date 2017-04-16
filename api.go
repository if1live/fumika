package fumika

type SearchResult struct {
	// metadata
	SearchedISBN string `json:"searched_isbn"`
	Title        string `json:"title"`

	UnitPrice   int `json:"unit_price"`
	PriceBest   int `json:"price_best"`
	PriceGood   int `json:"price_good"`
	PriceNormal int `json:"price_normal"`
}
