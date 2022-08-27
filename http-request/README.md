# 简介

一个可以捕获到http数据请求并打印的小工具

```
GET /index.php?id=123aaaaaaaa HTTP/1.1
Host: 127.0.0.1:9090
Connection: close
Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: en-US;q=0.9,en;q=0.8
Cache-Control: max-age=0
Connection: close
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36


POST /index.php HTTP/1.1
Host: 127.0.0.1:9090
Connection: close
Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: en-US;q=0.9,en;q=0.8
Cache-Control: max-age=0
Connection: close
Content-Length: 14
Content-Type: application/x-www-form-urlencoded
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36

id=123aaaaaaaa
POST /index.php HTTP/1.1
Host: 127.0.0.1:9090
Connection: close
Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: en-US;q=0.9,en;q=0.8
Cache-Control: max-age=0
Connection: close
Content-Length: 144
Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryLsORoJdVVGJg9vJG
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36

------WebKitFormBoundaryLsORoJdVVGJg9vJG
Content-Disposition: form-data; name="id"

123aaaaaaaa
------WebKitFormBoundaryLsORoJdVVGJg9vJG--
```