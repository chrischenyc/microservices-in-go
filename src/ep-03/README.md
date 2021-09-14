# ep-03: REST API basic

```bash
# start server at localhost:3000
go run main.go

# send os termination signal
ctrl^c

# expected output
#
# [api] 2021/09/14 17:07:04 received OS termination signal, gracefully shut downinterrupt
# [api] 2021/09/14 17:07:04 http: Server closed
```

```bash
# testing in another terminal

curl http://localhost:3000/products | jq

# expected output
#
# [
#   {
#     "id": "d5ce1cfc-a82a-4300-9655-eac0bde41f71",
#     "name": "Sobe - Lizard Fuel",
#     "description": "Networked 3rd generation emulation",
#     "sku": "WAU32AFD0FN874428",
#     "createdAt": "2021-01-04T12:34:56",
#     "updatedAt": "2021-01-04T12:34:56"
#   },
#   {
#     "id": "15823ea7-0a3f-4cdf-af26-2b119c113938",
#     "name": "Cherries - Bing, Canned",
#     "description": "Extended hybrid hierarchy",
#     "sku": "JM1GJ1T6XF1611223",
#     "price": 100,
#     "createdAt": "2021-03-04T12:34:56",
#     "updatedAt": "2021-03-04T12:34:56"
#   },
#   {
#     "id": "4e983def-db40-4334-9ade-a35d7bd1ab14",
#     "name": "Syrup - Monin, Amaretta",
#     "description": "Adaptive global contingency",
#     "sku": "19UUA65515A343666",
#     "price": 200,
#     "createdAt": "2021-05-04T12:34:56",
#     "updatedAt": "2021-06-04T12:34:56"
#   }
# ]

curl http://localhost:3000/products -X POST -v

# expected output
#
# *   Trying ::1...
# * TCP_NODELAY set
# * Connected to localhost (::1) port 3000 (#0)
# > POST /products HTTP/1.1
# > Host: localhost:3000
# > User-Agent: curl/7.64.1
# > Accept: */*
# >
# < HTTP/1.1 405 Method Not Allowed
# < Date: Tue, 14 Sep 2021 11:22:38 GMT
# < Content-Length: 0
# <
# * Connection #0 to host localhost left intact
# * Closing connection 0
```

## other references I read

- [Go Struct tag](https://github.com/golang/go/wiki/Well-known-struct-tags): I've been using it without knowing it has a name. I really like how easy it makes JSON parse/encode with data object
- [json.Encoder](https://pkg.go.dev/encoding/json#Encoder.Encode)
