package fumika

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_Aladin_createURI(t *testing.T) {
	cases := []struct {
		isbn string
		uri  string
	}{
		{"9788926790403", "http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx?KeyWord=9788926790403"},
	}
	for _, c := range cases {
		api := NewAladin(nil)
		got := api.createURI(c.isbn)
		if got != c.uri {
			t.Errorf("createURI - expected %q, got %q", c.uri, got)
		}
	}
}

func Test_Aladin_parse(t *testing.T) {
	cases := []struct {
		filepath string
		result   SearchResult
	}{
		{
			"testdata/aladin_9788926790403.html",
			SearchResult{
				SearchedISBN: "9788926790403",
				Title:        "기어와라! 냐루코 양 1",
				UnitPrice:    6000,
				PriceBest:    600,
				PriceGood:    500,
				PriceNormal:  400,
			},
		},
		{
			"testdata/aladin_8926790401.html",
			SearchResult{
				SearchedISBN: "8926790401",
				Title:        "기어와라! 냐루코 양 1",
				UnitPrice:    6000,
				PriceBest:    600,
				PriceGood:    500,
				PriceNormal:  400,
			},
		},
		{
			"testdata/aladin_9999999999999.html",
			SearchResult{
				SearchedISBN: "9999999999999",
				Title:        "",
				UnitPrice:    0,
				PriceBest:    0,
				PriceGood:    0,
				PriceNormal:  0,
			},
		},
	}

	for _, c := range cases {
		api := NewAladin(nil)
		b, err := ioutil.ReadFile(c.filepath)
		if err != nil {
			panic(err)
		}

		r := bytes.NewReader(b)
		got := api.parse(r)

		if got != c.result {
			t.Errorf("parse - expected %Q, got %Q", c.result, got)
		}
	}
}
