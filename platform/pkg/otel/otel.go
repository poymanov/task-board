package otel

import (
	"context"
	"fmt"
	"time"

	"github.com/poymanov/codemania-task-board/platform/pkg/otel/tracer"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
	tracer_noop "go.opentelemetry.io/otel/trace/noop"
)

var shutdownTracing func(ctx context.Context) error //nolint:gochecknoglobals

func SilentModeInit() {
	otel.SetTracerProvider(tracer_noop.NewTracerProvider())
	tracer.Init(otel.Tracer(""))

	log.Info().Msg("otel: Tracer is disabled")
}

func Init(ctx context.Context, appName, endpoint, namespace, instanceID string) error {
	if endpoint == "" {
		SilentModeInit()

		return nil
	}

	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)

	otel.SetTextMapPropagator(prop)

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(endpoint), otlptracegrpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter, trace.WithBatchTimeout(time.Second)),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
			semconv.ServiceNamespaceKey.String(namespace),
			semconv.ServiceInstanceIDKey.String(instanceID),
		)),
	)

	shutdownTracing = traceProvider.Shutdown

	otel.SetTracerProvider(traceProvider)
	tracer.Init(otel.Tracer(""))

	return nil
}

func Close() {
	if shutdownTracing == nil {
		return
	}

	err := shutdownTracing(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("otel: failed to shutdown tracing")
	}

	log.Info().Msg("otel: closed")
}
