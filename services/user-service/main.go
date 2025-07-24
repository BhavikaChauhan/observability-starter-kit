package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracerprovider"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func initTracer() func(context.Context) error {
	ctx := context.Background()

	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.Default()),
	)
	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func main() {
	shutdown := initTracer()
	defer shutdown(context.Background())

	tracer := otel.Tracer("user-service")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "handle-request")
		defer span.End()

		fmt.Fprintf(w, "Hello from user-service! 🎉")
		log.Println("Request handled")
		_ = ctx
	})

	log.Println("User Service running on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", otelhttp.NewHandler(handler, "root")))
}
