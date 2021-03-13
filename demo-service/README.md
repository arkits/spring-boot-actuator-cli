# Demo Service

This is a basic Spring Boot application that can be used to test / demo sba-cli.

## Getting Started

Requires Java 11.

## Build + Run

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
