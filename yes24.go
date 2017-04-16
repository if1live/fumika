package fumika

import (
	"io"

	"errors"

	"net/http"

	"golang.org/x/net/html"
)

/*
isbn : 9788926790403
full url : http://www.yes24.com/Mall/buyback/Search?CategoryNumber=018&SearchWord=9788926790403&SearchDomain=BOOK,FOREIGN&BuybackAccept=N
short url : http://www.yes24.com/Mall/buyback/Search?SearchWord=9788926790403
*/

type Yes24API struct {
	client *http.Client
}

func NewYes24(client *http.Client) *Yes24API {
	return &Yes24API{
		client: client,
	}
}

func (api *Yes24API) SearchISBN(isbn string) (SearchResult, error) {
	sanitizedISBN, ok := sanitizeISBN(isbn)
	if !ok {
		return SearchResult{}, errors.New("invalid isbn : " + isbn)
	}
	uri := api.createURI(sanitizedISBN)
	reader, err := uriToReader(uri, api.client)
	if err != nil {
		return SearchResult{}, err
	}

	result := api.parse(reader)
	return result, nil
}

func (api *Yes24API) createURI(isbn string) string {
	url := "http://www.yes24.com/Mall/buyback/Search"
	qs := "?SearchWord=" + isbn
	return url + qs
}

func (api *Yes24API) parse(reader io.Reader) SearchResult {
	result := SearchResult{}

	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	keyIsbnNode := GetElementByID(doc, "hidSearchWord")
	for _, attr := range keyIsbnNode.Attr {
		if attr.Key == "value" {
			result.SearchedISBN = attr.Val
		}
	}

	titleNode := GetElementByClassName(doc, "bbG_name")
	if titleNode != nil {
		// 검색 결과 있음
		titleLinkNode := GetElementsByTagName(titleNode, "a")[0]
		for _, attr := range titleLinkNode.Attr {
			if attr.Key == "title" {
				result.Title = attr.Val
			}
		}

		priceDiv := GetElementByClassName(doc, "bbG_price")
		priceTbody := GetElementsByTagName(priceDiv, "tbody")[0]
		priceNodes := GetElementsByTagName(priceTbody, "td")
		for idx, priceNode := range priceNodes {
			textNode := priceNode.FirstChild
			text := textNode.Data
			price, _ := sanitizePrice(text)

			if idx == 0 {
				result.UnitPrice = price
			} else if idx == 1 {
				result.PriceBest = price
			} else if idx == 2 {
				result.PriceGood = price
			} else if idx == 3 {
				result.PriceNormal = price
			}
		}
	}

	return result
}
