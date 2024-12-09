**Lightweight Cloud Storage Solution.** \\
This project implements a simple and efficient cloud storage solution using Go, designed to help users manage files seamlessly. The system is built with a RESTful API architecture, providing robust functionality for uploading, retrieving, and managing files.

Features:
File Management: Secure file upload, retrieval, and deletion functionality.
Scalable Architecture: Designed with scalability in mind, suitable for lightweight cloud storage needs.

Technologies Used:
Programming Language: Go (Golang)
Database: PostgreSQL for reliable data storage and retrieval.
API Framework: Gin for efficient and lightweight API development.
Containerization: Docker for simplified deployment and consistent environments.

## Getting started

  ```sh
  $ git clone github.com/st107853/cloud
  $ go get
  ```
Then you need to change the settings in .env

To run the code, use the `go run` command, like:

```sh
$ go run example.go
```

## PUT:
```sh
$ curl -X PUT -d 'VALUE' -v http://localhost:8080/v1/KEY
```

## GET:
```sh
$ curl http://localhost:8080/v1/KEY
```

## DELATE:
```sh
$ curl -X DELETE -v http://localhost:8080/v1/KEY
```