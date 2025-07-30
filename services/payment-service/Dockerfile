FROM eclipse-temurin:17-jdk

WORKDIR /app

# Copy Java app and agent
COPY target/*.jar app.jar
COPY docker/opentelemetry-javaagent.jar .

# Add agent
ENV JAVA_TOOL_OPTIONS="-javaagent:/app/opentelemetry-javaagent.jar"

EXPOSE 3001
ENTRYPOINT ["java", "-jar", "app.jar"]