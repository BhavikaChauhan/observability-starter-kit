package main

import (
    "context"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
    exporter, err := otlptracehttp.New(ctx)
    if err != nil {
        return nil, err
    }

    res, err := resource.New(ctx,
        resource.WithAttributes(
            semconv.ServiceNameKey.String("user-service"),
        ),
    )
    if err != nil {
        return nil, err
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(res),
    )
    otel.SetTracerProvider(tp)
    return tp, nil
}

func main() {
    ctx := context.Background()

    tp, err := initTracer(ctx)
    if err != nil {
        log.Fatalf("failed to initialize tracer: %v", err)
    }
    defer func() {
        _ = tp.Shutdown(ctx)
    }()

    r := gin.Default()
    r.Use(otelgin.Middleware("user-service"))

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"users": []string{"Alice", "Bob"}})
    })

    r.Run(":8080")
}
