FROM openjdk:17-jdk-slim

WORKDIR /app

# Copy your application JAR
COPY target/payment-service-0.0.1-SNAPSHOT.jar app.jar

# Copy the OpenTelemetry agent JAR
COPY opentelemetry-javaagent.jar opentelemetry-javaagent.jar

# Run the app with the agent
ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent.jar", "-jar", "app.jar"]
