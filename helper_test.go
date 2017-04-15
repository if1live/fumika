package fumika

import (
	"testing"
)

func Test_sanitizeISBN(t *testing.T) {
	cases := []struct {
		code string
		isbn string
		ok   bool
	}{
		{"9788926790403", "9788926790403", true},
		{"978-89-267-9040-3", "9788926790403", true},
		{"978 89 267 9040 3", "9788926790403", true},

		{"978892679040", "", false},
		{"97889267904033", "", false},

		{"8926790401", "8926790401", true},
	}

	for _, c := range cases {
		got, ok := sanitizeISBN(c.code)
		if ok != c.ok || got != c.isbn {
			t.Errorf("sanitizeISBN - expected (%q, %q), got (%q, %q)", c.isbn, c.ok, got, ok)
		}
	}
}

func Test_sanitizePrice(t *testing.T) {
	cases := []struct {
		text  string
		price int
		ok    bool
	}{
		{"6,000원", 6000, true},
		{"600원", 600, true},
	}
	for _, c := range cases {
		got, ok := sanitizePrice(c.text)
		if ok != c.ok || got != c.price {
			t.Errorf("sanitizePrice - expected (%q, %q), got (%q, %q)", c.price, c.ok, got, ok)
		}
	}
}
