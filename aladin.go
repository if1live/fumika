package fumika

import (
	"errors"
	"io"

	"golang.org/x/net/html"
)

/*
isbn : 9788926790403
full url : http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx?ActionType=1&SearchTarget=Book&KeyWord=9788926790403&x=0&y=0
simple url : http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx?KeyWord=9788926790403
알라딘 앱도 뜯어봤는데 결과가 HTML이었다. html 파싱을 피할수 없다면 그냥 데탑 기준으로 쓰자
*/

type AladinAPI struct {
}

func NewAladin() *AladinAPI {
	return &AladinAPI{}
}

func (api *AladinAPI) Search(isbn string) (SearchResult, error) {
	sanitizedISBN, ok := sanitizeISBN(isbn)
	if !ok {
		return SearchResult{}, errors.New("invalid isbn : " + isbn)
	}
	uri := api.createURI(sanitizedISBN)
	reader, err := uriToReader(uri)
	if err != nil {
		return SearchResult{}, err
	}

	result := api.parse(reader)
	return result, nil
}

func (api *AladinAPI) createURI(isbn string) string {
	url := "http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx"
	qs := "?KeyWord=" + isbn
	return url + qs
}

func (api *AladinAPI) parse(reader io.Reader) SearchResult {
	result := SearchResult{}

	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	keyIsbnNode := GetElementByID(doc, "KeyISBN")
	for _, attr := range keyIsbnNode.Attr {
		if attr.Key == "value" {
			result.SearchedISBN = attr.Val
		}
	}

	searchRoot := GetElementByID(doc, "searchResult")
	if searchRoot != nil {
		// searchResult 밑에는 tbody로 검색결과가 붙는다
		for c := searchRoot.FirstChild; c != nil; c = c.NextSibling {
			// tbody는 1개의 자식과 3개의 td를 가진다
			if c.Type != html.ElementNode {
				continue
			}

			titleLinkNode := GetElementByClassName(c, "c2b_b")
			titleNode := GetElementsByTagName(titleLinkNode, "strong")[0]
			title := titleNode.FirstChild.Data
			result.Title = title

			priceNodes := GetElementsByClassName(c, "c2b_tablet3")
			for idx, priceNode := range priceNodes {
				textNode := priceNode.FirstChild
				priceText := textNode.Data
				price, _ := sanitizePrice(priceText)

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
	}

	return result
}
