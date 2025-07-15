// +build telemetry
package main

import (
  "context"
  "log"
  "go.opentelemetry.io/otel"
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
  sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTelemetry() func() {
  exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
  if err != nil {
    log.Fatal(err)
  }
  tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
  otel.SetTracerProvider(tp)
  return func() { _ = tp.Shutdown(context.Background()) }
}