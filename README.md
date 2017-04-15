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
# download testdata
cd testdata
python fetch.py

go test
```

## Usage
### Aladin

```golang
api := fumika.CreateAladin()
isbn := "9788926790403"
result, found := api.Search(isbn)
if found {
    fmt.Printf("%Q", result)
} else {
    fmt.Println("not found")
}
```

### Yes24

```golang
api := fumika.CreateYes24()
isbn := "9788926790403"
result, found := api.Search(isbn)
if found {
    fmt.Printf("%Q", result)
} else {
    fmt.Println("not found")
}
```