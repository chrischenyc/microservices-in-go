# Building Microservices with Go

I'm following along Nick Jackson's [Building Microservices with Go](https://www.youtube.com/channel/UC2V1SxXFUa5YxVJvTsrCgyg) series on YouTube.

[![Building Microservices with Go - Nick Jackson](https://i.ytimg.com/vi/VzBGi_n65iU/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLC8YgF-MxFlwypyOgx-L2wY9yAFdg)](https://www.youtube.com/channel/UC2V1SxXFUa5YxVJvTsrCgyg)

I love his style of teaching.

---

## Background

I started using Go since Apr 2021, but mostly with [AWS SDK for Go V2](https://github.com/aws/aws-sdk-go-v2) and [Lambda in Go](https://github.com/aws/aws-lambda-go).

Come to think of it, my previous project was also a serverless one on AWS, but with Javascript.

So last time I developed commerical container-based system was around 2019-2020. And it was a [monolith server](https://www.capturedlabs.com/journal#mr-yum-2019-2020) backed by TypeScript, Node.js, PostgreSQL, and of course, Kubernetes. We ran into the textbook [Monolith Hell](https://livebook.manning.com/book/microservices-patterns/chapter-1/), a prevailing architecture struggle described in Chris Richardson's book [Microservices Patterns](https://www.manning.com/books/microservices-patterns). I'm reading this book by the way.

[![Microservices Patterns - Chris Richardson](https://images.manning.com/360/480/resize/book/b/dc43dfc-e43d-419d-b577-3809c6967442/Richardson-MP-HI.png)](https://www.manning.com/books/microservices-patterns)

## Goals

- learn more Go
- brush up Docker skills
- brush up Kubernetes skills
- brush up building REST APIs
- deep dive in Microservices with container and REST API
- learn gRPC
- try GCP
- last but not least: learn Nick Jackson's style (teaching, not hair)

## Logs

- 2021-09-12: init commit!
- 2021-09-12: [ep-01](./src/ep-01) hello world http server
- 2021-09-14: [ep-02](./src/ep-02) refactor to http handler modules. shut down server gracefully
- 2021-09-14: [ep-03](./src/ep-03) basic REST api
- 2021-09-16: [ep-04](./src/ep-04) basic REST api CRUD
- 2021-09-26: [ep-05](./src/ep-05) use [Gorilla Mux](https://github.com/gorilla/mux) to bootstrap REST routes. it is very similar to [express.js](https://expressjs.com/).
- 2021-10-02: [ep-06](./src/ep-06) use [Go validator](https://github.com/go-playground/validator) to validate REST request input. it is a different approach from [joi](https://github.com/sideway/joi), where the validation rules are tied with HTTP handler middleware.
- 2021-10-04: [ep-07](./src/ep-07) generate Swagger API docs.

## Tricks

- auto-compile and re-start the server on file changes. unlike Javascript, Go community doesn't seem to have a go-to solution.

  ```bash
  # https://techinscribed.com/5-ways-to-live-reloading-go-applications/

  npx nodemon --exec go run main.go --signal SIGTERM
  ```

- run go tests recursively in all sub directories

  ```bash
  go test -v ./...
  ```
