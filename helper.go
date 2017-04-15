package fumika

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
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

func uriToReader(uri string) (io.Reader, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// euc kr-> utf-8
	// aladin, yes24 모두 euc-kr 기반이라서 일단 하드코딩
	utf8buf := euckrToUTF8(&buf)
	reader := bytes.NewReader(utf8buf.Bytes())
	return reader, nil
}

func euckrToUTF8(input *bytes.Buffer) *bytes.Buffer {
	var utf8buf bytes.Buffer
	wr := transform.NewWriter(&utf8buf, korean.EUCKR.NewDecoder())
	wr.Write(input.Bytes())
	wr.Close()
	return &utf8buf
}
