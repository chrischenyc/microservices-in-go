# Building Microservices with Go

I'm follow along Nick Jackson's [Building Microservices with Go](https://www.youtube.com/channel/UC2V1SxXFUa5YxVJvTsrCgyg) series on YouTube.

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
- brush Kubernetes skills
- deep dive in Microservices architecture based on container and API
- learn gRPC
- try GCP
- last but not least: learn Nick Jackson's style (teaching, not hair)

## Logs

- 2021-09-12: init commit!
- 2021-09-12: [ep-01](./src/ep-01) hello world http server
- 2021-09-14: [ep-02](./src/ep-02) refactor to http handler modules. shut down server gracefully

## Tricks

- auto-compile and re-start the server on file changes. unlike Javascript, Go community doesn't seem to have a go-to solution.

  ```bash
  # https://techinscribed.com/5-ways-to-live-reloading-go-applications/

  npx nodemon --exec go run main.go --signal SIGTERM
  ```
