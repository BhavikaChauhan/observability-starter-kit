# Stage 1: Build the Spring Boot application
FROM maven:3.8.5-openjdk-17 AS build

WORKDIR /app

# Copy pom.xml and download dependencies first (for caching)
COPY pom.xml .
RUN mvn dependency:go-offline

# Copy source and build
COPY src ./src
RUN mvn clean package -DskipTests

# Stage 2: Run with OpenTelemetry Java Agent
FROM eclipse-temurin:17-jre

WORKDIR /app

# Copy the built jar from the previous stage
COPY --from=build /app/target/*.jar payment-service.jar

# Copy the OpenTelemetry Java Agent (ensure you have it locally!)
COPY opentelemetry-javaagent.jar opentelemetry-javaagent.jar

# Expose the port your Spring Boot app runs on
EXPOSE 3001

# Run the app with the OTEL agent
ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent.jar", "-jar", "payment-service.jar"]