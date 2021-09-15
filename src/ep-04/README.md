# ep-04: REST API basic CRUD

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

curl http://localhost:3000/products/ | jq
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

curl http://localhost:3000/products/3 -X PUT -d '{"name":"new name"}'
curl http://localhost:3000/products/ | jq
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
```

```bash
curl http://localhost:3000/products/3 -X DELETE
curl http://localhost:3000/products/ | jq
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

- [Let's Go - Routing Requests](https://lets-go.alexedwards.net/sample/02.04-routing-requests.html): I wanted to do wildcard routing, i.e.: /products/*, it turned out I just needed to use /products/ with a slash at the end, it's Subtree Pattern

- [fmt.Errorf vs errors.New](https://www.reddit.com/r/golang/comments/6ffrie/fmterrorf_or_errorsnew/): I'm still not quite clear, need to read this [PDF](https://dave.cheney.net/paste/gocon-spring-2016.pdf)

- [how to merge 2 structs with same type](https://stackoverflow.com/questions/47395430/merge-fields-two-structs-of-same-type): I wish I could have something equivalent to Javascript spread:

  ```javascript
  c = {...a, ...b}
  ```
