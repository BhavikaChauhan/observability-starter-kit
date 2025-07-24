FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Install auto-instrumentation tool
RUN pip install opentelemetry-distro

EXPOSE 3002

# Entry point with OpenTelemetry instrumentation
CMD ["opentelemetry-instrument", "python", "app.py"]
