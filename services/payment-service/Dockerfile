# Stage 1: Build the Spring Boot application
FROM maven:3.8.5-openjdk-17 AS build

WORKDIR /app

# Copy only pom.xml first and download dependencies for better layer caching
COPY pom.xml .
RUN mvn dependency:go-offline

# Now copy the source code and build
COPY src ./src
RUN mvn clean package -DskipTests

# Stage 2: Runtime image with OTEL agent
FROM eclipse-temurin:17-jre

WORKDIR /app

# Copy the built Spring Boot JAR from stage 1
COPY --from=build /app/target/*.jar payment-service.jar

# Download OpenTelemetry Java Agent
RUN curl -L -o opentelemetry-javaagent.jar https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar

# Expose the app port
EXPOSE 3001

# Start the Spring Boot app with OTEL instrumentation
ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent.jar", "-jar", "payment-service.jar"]
