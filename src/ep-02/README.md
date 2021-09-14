# ep-02: Hello World! http server

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

curl http://localhost:3000 -d "Hello World"

# expected output
#
# requested with
# method: GET
# path: /
# body:

curl http://localhost:3000/goodbye

# expected output
#
# Goodbye!
```

## other references I read

- go channel
  - <https://gobyexample.com/channels>
  - <https://gobyexample.com/channel-buffering>
- go signal
  - <https://gobyexample.com/signals>
