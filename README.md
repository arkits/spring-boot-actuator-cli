# Spring Boot Actuator CLI

sba-cli is a command-line utility to work with a [Spring Boot](https://spring.io/projects/spring-boot) application's [Actuator](https://docs.spring.io/spring-boot/docs/current/reference/html/production-ready-features.html) endpoint.

![actuator screenshot](./docs/screenshots/actuator.png)

## Usage

### Quick Start

sba-cli is available as a stand-alone binary. Pre-built binaries are available in the [Releases](https://github.com/arkits/spring-boot-actuator-cli/releases).

The most basic usage of sba-cli is to check the info of a Spring Boot app.

```bash
# ./sba-cli info -U <baseURL to your app>
$ ./sba-cli info -U http://localhost:8080

{
    "title": "user-service"
}
```

sba-cli aims to provides a more legible output for most of the default Actuator endpoints. Currently support endpoints are - `info`, `env`.

Users can still reach unsupported or custom endpoints and make use of the Inventory functionality by using the `custom` command.

```bash
# ./sba-cli custom \
#    -E <custom endpoint's path ID> \
#    -U <baseURL to your app> \
#    --skip-pretty-print
$ ./sba-cli custom \
    -E prometheus \
    -U http://localhost:8080 \
    --skip-pretty-print

# HELP jvm_buffer_total_capacity_bytes An estimate of the total capacity of the buffers in this pool
# TYPE jvm_buffer_total_capacity_bytes gauge
jvm_buffer_total_capacity_bytes{application="user-service",id="mapped",} 0.0
jvm_buffer_total_capacity_bytes{application="user-service",id="direct",} 81920.0
...
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

Multiple specific services can be passed as a comma-separated string. sba-cli will iterate and print the responses for each.

```bash
$ ./sba-cli info -S user-service,order-service

{
    "title": "user-service"
}

{
    "title": "order-service"
}
```

### Debugging with Verbose

```bash
$ ./sba-cli custom -V \
    --url http://localhost:8080 \
    --auth-header "Basic YXJraXRzOmh1bnRlcjI=" \
    --endpoint health
>>> url was set... assuming impromptu definition
>>> GET http://localhost:8080
>>> Authorization: Basic YXJraXRzOmh1bnRlcjI=
>>> Proto: HTTP/1.1 Status: 200
{
  "status": "UP"
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
