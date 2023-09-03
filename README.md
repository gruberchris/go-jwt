# Go API With JWT Auth Example

This is a quick tutorial on how to create a simple RESTful API using Go. The API will use JWT for authentication and demonstrate route authorization.

## Dependencies

- [Gin](https://gin-gonic.com/docs/quickstart/) - HTTP web framework
- [Gorm](https://gorm.io/docs/index.html) - ORM
- [Go-mysql-driver](https://gorm.io/docs/connecting_to_the_database.html#MySQL) - MySQL driver
- [JWT](https://github.com/golang-jwt/jwt#jwt-go) - JWT authentication
- [Godotenv](https://github.com/joho/godotenv#godotenv--) - Environment variables

## Compile Daemon for Go

```shell
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

### Run

First use the *build* script in the Makefile to build the binary. Then, it may be launched for development by using the following script in the shell:

```shell
CompileDaemon --command="./bin/go-jwt"
```

Alternatively, the binary may be launched directly:

```shell
./bin/go-jwt
```

or by using the *run* script in the Makefile:

```shell
make run
```