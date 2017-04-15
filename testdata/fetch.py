#!/usr/bin/env python3

import requests

targets = (
    {
        "url": "http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx?KeyWord=8926790401",
        "file": "aladin_8926790401.html",
    },
    {
        "url": "http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx?KeyWord=9999999999999",
        "file": "aladin_9999999999999.html",
    },
    {
        "url": "http://off.aladin.co.kr/shop/usedshop/wc2b_search.aspx?KeyWord=9788926790403",
        "file": "aladin_9788926790403.html",
    },
    {
        "url": "http://www.yes24.com/Mall/buyback/Search?SearchWord=9999999999",
        "file": "yes24_9999999999.html",
    },
    {
        "url": "http://www.yes24.com/Mall/buyback/Search?SearchWord=9788926790403",
        "file": "yes24_9788926790403.html",
    },
)

for target in targets:
    url = target["url"]
    filename = target["file"]

    f = open(filename, 'w', encoding='utf-8')
    r = requests.get(url)
    f.write(r.text)
    f.close()
