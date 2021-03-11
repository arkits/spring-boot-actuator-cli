# Spring Boot Actuator CLI

`sba-cli` is a command-line utility to interface with a [Spring Boot](https://spring.io/projects/spring-boot) application's [Actuator](https://docs.spring.io/spring-boot/docs/current/reference/html/production-ready-features.html) endpoint.

## Usage

```bash
$ ./sba-cli info -U http://localhost:8080

{
    "title": "user-service"
}
```

## Development

### Building + Running Locally

```bash
# Download Go dependencies
$ go mod download

# Run sba-cli
$ go run main.go
```
