# Demo Service

This is a basic Spring Boot application that can be used to test / demo sba-cli.

## Getting Started

`demo-service` supports Docker multi-stage builds - This allows you to build and run the demo-service without having to setup JDK / Java on your machine.

```bash
# Build the service + image
$ docker-compose build demo

# Run the image
$ docker-compose up demo
```

## Build + Run

Requires Java 11.

```bash
./gradlew bootRun
```

## Build a fat-JAR file

```bash
# Build the JAR file
./gradlew clean build

# Execute it
java -jar build/libs/demo-*.jar
```

## Tests

```bash
./gradlew clean test
```

## Gradle Tasks

Check for available task in gradle

```bash
./gradlew tasks
```
