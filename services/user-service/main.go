package main

import (
	"log"
	"net/http"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/sdk"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"context"
)

func initTracer() (*trace.TracerProvider, error) {
	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("user-service"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	log.Println("Starting User Service on port 8000...")

	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = tp.Shutdown(context.Background()) }()

	mux := http.NewServeMux()
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from user-service!"))
	}), "root-handler"))

	log.Fatal(http.ListenAndServe(":8000", mux))
}
