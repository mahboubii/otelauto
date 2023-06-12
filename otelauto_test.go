package otelauto

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/metric/noop"
)

func TestDefault(t *testing.T) {
	assert.NotPanics(t, func() {
		counter := Int64Counter("my.counter")
		counter.Add(context.Background(), 1)
	})
}

func TestWith(t *testing.T) {
	assert.NotPanics(t, func() {
		counter := With(noop.NewMeterProvider().Meter("")).Int64Counter("my.counter")
		counter.Add(context.Background(), 1)
	})
}

func TestWithNil(t *testing.T) {
	assert.Panics(t, func() {
		counter := With(nil).Int64Counter("my.counter")
		counter.Add(context.Background(), 1)
	})
}
