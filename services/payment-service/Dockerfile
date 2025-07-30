FROM eclipse-temurin:17-jdk-alpine

WORKDIR /app

# Copy your compiled Spring Boot JAR
COPY target/payment-service-1.0.0.jar app.jar

# Download the OpenTelemetry Java agent JAR directly into the image
ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v1.38.0/opentelemetry-javaagent-1.38.0.jar opentelemetry-javaagent.jar

# Set the entrypoint to use the OTEL agent
ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent.jar", "-jar", "app.jar"]