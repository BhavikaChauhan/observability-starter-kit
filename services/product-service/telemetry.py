from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor, ConsoleSpanExporter

def init_telemetry():
    tracer_provider = TracerProvider()
    trace.set_tracer_provider(tracer_provider)
    exporter = ConsoleSpanExporter()
    tracer_provider.add_span_processor(BatchSpanProcessor(exporter))

def tracer():
    return trace.get_tracer(__name__)