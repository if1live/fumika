package fumika

type SearchResult struct {
	// metadata
	SearchedISBN string
	Title        string

	UnitPrice   int
	PriceBest   int
	PriceGood   int
	PriceNormal int
}
