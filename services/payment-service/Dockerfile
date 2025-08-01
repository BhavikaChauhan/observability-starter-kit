FROM eclipse-temurin:17-jdk-alpine

WORKDIR /app

COPY target/payment-service-1.0.0.jar app.jar

ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar otel-javaagent.jar

EXPOSE 8083

ENTRYPOINT ["java", "-javaagent:otel-javaagent.jar", "-jar", "app.jar"]