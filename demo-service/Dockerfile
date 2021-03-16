FROM adoptopenjdk/openjdk11:jdk-11.0.10_9-alpine AS BUILD_STAGE

EXPOSE 8080

# Set working directory
WORKDIR /sba-cli/demo-service

# Copy demo-service's source code
COPY ./demo-service ./

# Copy the .git directory - required for generating/injecting git info during build time
COPY ./.git ./.git

# Check what was added to the container
RUN ls -la

# Build an artifact
RUN ./gradlew clean build

# TODO: switch to JRE
FROM adoptopenjdk/openjdk11:jdk-11.0.10_9-alpine 

# Set working directory
WORKDIR /sba-cli/demo-service

# Use the JAR file for execution
COPY --from=BUILD_STAGE /sba-cli/demo-service/build/libs/demo-*.jar /demo.jar

# Check what was added to the container
RUN ls -la

CMD ["java", "-Dspring.profiles.active=dockerized", "-jar", "/demo.jar"]