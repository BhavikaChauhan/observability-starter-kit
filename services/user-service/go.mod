module user-service

go 1.22

require (
<<<<<<< HEAD
	go.opentelemetry.io/otel v1.20.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.20.0
	go.opentelemetry.io/otel/sdk v1.20.0
)

require (
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
=======
	go.opentelemetry.io/otel v1.25.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.25.0
	go.opentelemetry.io/otel/sdk v1.25.0
)
>>>>>>> 6a0ee1b47636e05814d98c4ba8e3d1ea2bebba70
