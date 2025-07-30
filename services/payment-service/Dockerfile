FROM eclipse-temurin:17-jdk-alpine

WORKDIR /app

# Copy your built JAR (make sure it matches actual file name)
COPY target/payment-service-1.0.0.jar app.jar

# Download the stable OpenTelemetry agent at build time
ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar opentelemetry-javaagent.jar

# Expose your application's port
EXPOSE 3001

# Use the agent when running
ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent.jar", "-jar", "app.jar"]
