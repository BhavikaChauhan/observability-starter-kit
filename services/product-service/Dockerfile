FROM python:3.10-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install opentelemetry-distro opentelemetry-exporter-otlp opentelemetry-instrumentation-flask

COPY . .

EXPOSE 3002

CMD ["opentelemetry-instrument", "python", "app.py"]