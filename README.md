# Spring Boot Actuator CLI

sba-cli is a command-line utility to work with a [Spring Boot](https://spring.io/projects/spring-boot) application's [Actuator](https://docs.spring.io/spring-boot/docs/current/reference/html/production-ready-features.html) endpoint.

## Usage

### Most Basic

The most basic usage of sba-cli is to check the info of a Spring Boot app.

```bash
# ./sba-cli info -U <baseURL to your app>
$ ./sba-cli info -U http://localhost:8080

{
    "title": "user-service"
}
```

### Inventory

Chances are that you are managing multiple micro-services. sba-cli is designed to support this is use case by allowing the user to supply an Inventory. An Inventory can be defined in a `config.yaml` that must be placed in the same directory as sba-cli. Refer to the [config.sample.yaml](./config.sample.yaml)

After defining multiple services in your `config.yaml`, you can refer to a specific service by passing it's name in `-S` flag.

```bash
# ./sba-cli info -S <name of a specific service>
$ ./sba-cli info -S user-service

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
