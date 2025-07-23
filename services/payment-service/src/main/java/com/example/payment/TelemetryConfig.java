package com.example.payment;

import io.opentelemetry.exporter.logging.LoggingSpanExporter;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;

import org.springframework.context.annotation.Configuration;

@Configuration
public class TelemetryConfig {
    private SdkTracerProvider sdkTracerProvider;

    @PostConstruct
    public void init() {
        sdkTracerProvider = SdkTracerProvider.builder()
            .addSpanProcessor(BatchSpanProcessor.builder(new LoggingSpanExporter()).build())
            .build();
    }

    @PreDestroy
    public void destroy() {
        sdkTracerProvider.shutdown();
    }
}

