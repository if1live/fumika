package fumika

import (
	"bytes"
	"strconv"
)

func sanitizeISBN(code string) (string, bool) {
	var buffer bytes.Buffer
	for _, c := range code {
		if c >= '0' && c <= '9' {
			buffer.WriteRune(c)
		}
	}
	isbn := buffer.String()

	if len(isbn) == 10 || len(isbn) == 13 {
		return isbn, true
	}
	return "", false
}

func sanitizePrice(text string) (int, bool) {
	var buffer bytes.Buffer
	for _, c := range text {
		if c >= '0' && c <= '9' {
			buffer.WriteRune(c)
		}
	}
	priceText := buffer.String()
	price, err := strconv.Atoi(priceText)
	if err != nil {
		return 0, false
	}
	return price, true
}
