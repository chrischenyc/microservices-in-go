# ep-02: Hello World! http server

```bash
# start server at localhost:3000
go run main.go
```

```bash
curl http://localhost:3000 -d "Hello World" -v

# expected output
#
# *   Trying ::1...
# * TCP_NODELAY set
# * Connected to localhost (::1) port 3000 (#0)
# > POST / HTTP/1.1
# > Host: localhost:3000
# > User-Agent: curl/7.64.1
# > Accept: */*
# > Content-Length: 11
# > Content-Type: application/x-www-form-urlencoded
# >
# * upload completely sent off: 11 out of 11 bytes
# < HTTP/1.1 200 OK
# < Date: Sun, 12 Sep 2021 03:56:30 GMT
# < Content-Length: 54
# < Content-Type: text/plain; charset=utf-8
# <
# requested with
# method: POST
# path: /
# body: Hello World
# * Connection #0 to host localhost left intact
# * Closing connection 0
```
