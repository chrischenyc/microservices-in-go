# ep-05: use Gorilla framework

```bash
# start server at localhost:3000
go run main.go

# send os termination signal
ctrl^c

# expected output
#
# [api] 2021/09/14 17:07:04 received OS termination signal, gracefully shut down
# [api] 2021/09/14 17:07:04 http: Server closed
```

```bash
# testing in another terminal

curl http://localhost:3000/products | jq
# [
#   {
#     "id": 1,
#     "name": "Sobe - Lizard Fuel",
#     "description": "Networked 3rd generation emulation",
#     "sku": "WAU32AFD0FN874428",
#     "createdAt": "2021-01-04T12:34:56",
#     "updatedAt": "2021-01-04T12:34:56"
#   },
#   {
#     "id": 2,
#     "name": "Cherries - Bing, Canned",
#     "description": "Extended hybrid hierarchy",
#     "sku": "JM1GJ1T6XF1611223",
#     "price": 100,
#     "createdAt": "2021-03-04T12:34:56",
#     "updatedAt": "2021-03-04T12:34:56"
#   },
#   {
#     "id": 3,
#     "name": "Syrup - Monin, Amaretta",
#     "description": "Adaptive global contingency",
#     "sku": "19UUA65515A343666",
#     "price": 200,
#     "createdAt": "2021-05-04T12:34:56",
#     "updatedAt": "2021-06-04T12:34:56"
#   }
# ]

curl http://localhost:3000/products -X POST -d '{"name":"new product", "description":"a brand new product", "sku": "abcdefg12345"}'
curl http://localhost:3000/products | jq
#
# [
#   {
#     "id": 1,
#     "name": "Sobe - Lizard Fuel",
#     "description": "Networked 3rd generation emulation",
#     "sku": "WAU32AFD0FN874428",
#     "createdAt": "2021-01-04T12:34:56",
#     "updatedAt": "2021-01-04T12:34:56"
#   },
#   {
#     "id": 2,
#     "name": "Cherries - Bing, Canned",
#     "description": "Extended hybrid hierarchy",
#     "sku": "JM1GJ1T6XF1611223",
#     "price": 100,
#     "createdAt": "2021-03-04T12:34:56",
#     "updatedAt": "2021-03-04T12:34:56"
#   },
#   {
#     "id": 3,
#     "name": "new name",
#     "description": "",
#     "sku": "",
#     "createdAt": "",
#     "updatedAt": ""
#   },
#   {
#     "id": 4,
#     "name": "new product",
#     "description": "a brand new product",
#     "sku": "abcdefg12345",
#     "createdAt": "",
#     "updatedAt": ""
#   }
# ]

curl http://localhost:3000/products/3 -X PUT -d '{"name":"new name"}'
curl http://localhost:3000/products | jq
#
# [
#   {
#     "id": 1,
#     "name": "Sobe - Lizard Fuel",
#     "description": "Networked 3rd generation emulation",
#     "sku": "WAU32AFD0FN874428",
#     "createdAt": "2021-01-04T12:34:56",
#     "updatedAt": "2021-01-04T12:34:56"
#   },
#   {
#     "id": 2,
#     "name": "Cherries - Bing, Canned",
#     "description": "Extended hybrid hierarchy",
#     "sku": "JM1GJ1T6XF1611223",
#     "price": 100,
#     "createdAt": "2021-03-04T12:34:56",
#     "updatedAt": "2021-03-04T12:34:56"
#   },
#   {
#     "id": 3,
#     "name": "new name",
#     "description": "",
#     "sku": "",
#     "createdAt": "",
#     "updatedAt": ""
#   }
# ]

curl http://localhost:3000/products/3 -X DELETE
curl http://localhost:3000/products | jq
# [
#   {
#     "id": 1,
#     "name": "Sobe - Lizard Fuel",
#     "description": "Networked 3rd generation emulation",
#     "sku": "WAU32AFD0FN874428",
#     "createdAt": "2021-01-04T12:34:56",
#     "updatedAt": "2021-01-04T12:34:56"
#   },
#   {
#     "id": 2,
#     "name": "Cherries - Bing, Canned",
#     "description": "Extended hybrid hierarchy",
#     "sku": "JM1GJ1T6XF1611223",
#     "price": 100,
#     "createdAt": "2021-03-04T12:34:56",
#     "updatedAt": "2021-03-04T12:34:56"
#   }
# ]

curl http://localhost:3000/products/3 -X DELETE
# product not found
```

## other references I read

- [Gorilla Mux](https://github.com/gorilla/mux): Gorilla Mux to Golang feels like Express.js to Node.js
