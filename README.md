

# fumika
중고책 매입가 검색 API

[![Build Status](https://travis-ci.org/if1live/fumika.svg?branch=master)](https://travis-ci.org/if1live/fumika)

## Features
* 검색 가능한 인터넷 서점
    * 알라딘 : http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx 파싱
    * Yes24 : http://www.yes24.com/Mall/buyback/Search 파싱
* 검색 가능한 키워드
    * ISBN
* 검색 가능한 정보
    * 제목
    * 정가
    * 매입가 (최상)
    * 매입가 (상)
    * 매입가 (중)


## Test
```bash
# (optional) download testdata 
cd testdata
python fetch.py

go test
```

## Usage

### Aladin

`examples/aladin/main.go`

```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/if1live/fumika"
)

func main() {
	client := http.Client{}
	api := fumika.NewAladin(&client)
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
```

```
title:  기어와라! 냐루코 양 1
unit price:  6000
price best:  600
price good:  500
price normal:  400
```

### Yes24

`examples/yes24/main.go`

```golang
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
```

```
title:  기어와라! 냐루코 양 1
unit price:  6000
price best:  600
price good:  600
price normal:  500
```

## Development Note
### Build README
```bash
go get github.com/if1live/maya
maya -file=document/README.mkdn -mode=pelican > README.md
```