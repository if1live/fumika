# fumika
중고책 매입가 검색 API

## Features
* 검색 가능한 인터넷 서점
    * 알라딘
    * Yes24
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

```golang
api := fumika.NewAladin()
isbn := "9788926790403"
result, err := api.Search(isbn)
if err != nil {
    panic(err)
}

fmt.Println("Aladin API")
printSearchResult(&result)
```

```
Aladin API
Title : 기어와라! 냐루코 양 1
UnitPrice : 6000
PriceBest : 600
PriceGood : 500
PriceNormal : 400
```

### Yes24

```golang
api := fumika.NewYes24()
isbn := "9788926790403"
result, err := api.Search(isbn)
if err != nil {
    panic(err)
}

fmt.Println("Yes24 API")
printSearchResult(&result)
```

```
Yes24 API
Title : 기어와라! 냐루코 양 1
UnitPrice : 6000
PriceBest : 600
PriceGood : 600
PriceNormal : 500
```