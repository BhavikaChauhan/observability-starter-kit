module user-service

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.49.0
    go.opentelemetry.io/otel v1.24.0
    go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.24.0
    go.opentelemetry.io/otel/sdk v1.24.0
)
