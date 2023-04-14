package otelauto

// Package otelauto provides alternative and more handy constructors for the fundamental
// OpenTelemetry metric types inspired by github.com/prometheus/client_golang/prometheus/promauto.
// All constructors panic if the registration fails.
//
// The following example is a complete program to create a counter:
//
//	package main
//
//	import (
//		"ctx"
//
//		"github.com/mahboubii/otelauto"
//	)
//
//	var counter = Int64Counter("my.counter")
//
//	func main() {
//  	counter.Add(context.Background(), 1)
//	}
//

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
)

// Int64Counter returns a new instrument identified by name and configured
// with options. The instrument is used to synchronously record increasing
// int64 measurements during a computational operation.
func Int64Counter(name string, opts ...instrument.Int64Option) instrument.Int64Counter {
	return withDefault().Int64Counter(name, opts...)
}

// Int64UpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// int64 measurements during a computational operation.
func Int64UpDownCounter(name string, opts ...instrument.Int64Option) instrument.Int64UpDownCounter {
	return withDefault().Int64UpDownCounter(name, opts...)
}

// Int64Histogram returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// the distribution of int64 measurements during a computational operation.
func Int64Histogram(name string, opts ...instrument.Int64Option) instrument.Int64Histogram {
	return withDefault().Int64Histogram(name, opts...)
}

// Int64ObservableCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// increasing int64 measurements once per a measurement collection cycle.
func Int64ObservableCounter(name string, opts ...instrument.Int64ObserverOption) instrument.Int64ObservableCounter {
	return withDefault().Int64ObservableCounter(name, opts...)
}

// Int64ObservableUpDownCounter returns a new instrument identified by name
// and configured with options. The instrument is used to asynchronously
// record int64 measurements once per a measurement collection cycle.
func Int64ObservableUpDownCounter(name string, opts ...instrument.Int64ObserverOption) instrument.Int64ObservableUpDownCounter {
	return withDefault().Int64ObservableUpDownCounter(name, opts...)
}

// Int64ObservableGauge returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// instantaneous int64 measurements once per a measurement collection
// cycle.
func Int64ObservableGauge(name string, opts ...instrument.Int64ObserverOption) instrument.Int64ObservableGauge {
	return withDefault().Int64ObservableGauge(name, opts...)
}

// Float64Counter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// increasing float64 measurements during a computational operation.
func Float64Counter(name string, opts ...instrument.Float64Option) instrument.Float64Counter {
	return withDefault().Float64Counter(name, opts...)
}

// Float64UpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// float64 measurements during a computational operation.
func Float64UpDownCounter(name string, opts ...instrument.Float64Option) instrument.Float64UpDownCounter {
	return withDefault().Float64UpDownCounter(name, opts...)
}

// Float64Histogram returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// the distribution of float64 measurements during a computational
// operation.
func Float64Histogram(name string, opts ...instrument.Float64Option) instrument.Float64Histogram {
	return withDefault().Float64Histogram(name, opts...)
}

// Float64ObservableCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// increasing float64 measurements once per a measurement collection cycle.
func Float64ObservableCounter(name string, opts ...instrument.Float64ObserverOption) instrument.Float64ObservableCounter {
	return withDefault().Float64ObservableCounter(name, opts...)
}

// Float64ObservableUpDownCounter returns a new instrument identified by
// name and configured with options. The instrument is used to
// asynchronously record float64 measurements once per a measurement
// collection cycle.
func Float64ObservableUpDownCounter(name string, opts ...instrument.Float64ObserverOption) instrument.Float64ObservableUpDownCounter {
	return withDefault().Float64ObservableUpDownCounter(name, opts...)
}

// Float64ObservableGauge returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// instantaneous float64 measurements once per a measurement collection
// cycle.
func Float64ObservableGauge(name string, opts ...instrument.Float64ObserverOption) instrument.Float64ObservableGauge {
	return withDefault().Float64ObservableGauge(name, opts...)
}

// RegisterCallback registers f to be called during the collection of a
// measurement cycle.
//
// If Unregister of the returned Registration is called, f needs to be
// unregistered and not called during collection.
//
// The instruments f is registered with are the only instruments that f may
// observe values for.
//
// If no instruments are passed, f should not be registered nor called
// during collection.
func RegisterCallback(c metric.Callback, instruments ...instrument.Asynchronous) metric.Registration {
	return withDefault().RegisterCallback(c, instruments...)
}

type Factory struct {
	mp metric.Meter
}

// Int64Counter returns a new instrument identified by name and configured
// with options. The instrument is used to synchronously record increasing
// int64 measurements during a computational operation.
func (f Factory) Int64Counter(name string, opts ...instrument.Int64Option) instrument.Int64Counter {
	i, err := f.mp.Int64Counter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Int64UpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// int64 measurements during a computational operation.
func (f Factory) Int64UpDownCounter(name string, opts ...instrument.Int64Option) instrument.Int64UpDownCounter {
	i, err := f.mp.Int64UpDownCounter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Int64Histogram returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// the distribution of int64 measurements during a computational operation.
func (f Factory) Int64Histogram(name string, opts ...instrument.Int64Option) instrument.Int64Histogram {
	i, err := f.mp.Int64Histogram(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Int64ObservableCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// increasing int64 measurements once per a measurement collection cycle.
func (f Factory) Int64ObservableCounter(name string, opts ...instrument.Int64ObserverOption) instrument.Int64ObservableCounter {
	i, err := f.mp.Int64ObservableCounter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Int64ObservableUpDownCounter returns a new instrument identified by name
// and configured with options. The instrument is used to asynchronously
// record int64 measurements once per a measurement collection cycle.
func (f Factory) Int64ObservableUpDownCounter(name string, opts ...instrument.Int64ObserverOption) instrument.Int64ObservableUpDownCounter {
	i, err := f.mp.Int64ObservableUpDownCounter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Int64ObservableGauge returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// instantaneous int64 measurements once per a measurement collection
// cycle.
func (f Factory) Int64ObservableGauge(name string, opts ...instrument.Int64ObserverOption) instrument.Int64ObservableGauge {
	i, err := f.mp.Int64ObservableGauge(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Float64Counter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// increasing float64 measurements during a computational operation.
func (f Factory) Float64Counter(name string, opts ...instrument.Float64Option) instrument.Float64Counter {
	i, err := f.mp.Float64Counter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Float64UpDownCounter returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// float64 measurements during a computational operation.
func (f Factory) Float64UpDownCounter(name string, opts ...instrument.Float64Option) instrument.Float64UpDownCounter {
	i, err := f.mp.Float64UpDownCounter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Float64Histogram returns a new instrument identified by name and
// configured with options. The instrument is used to synchronously record
// the distribution of float64 measurements during a computational
// operation.
func (f Factory) Float64Histogram(name string, opts ...instrument.Float64Option) instrument.Float64Histogram {
	i, err := f.mp.Float64Histogram(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Float64ObservableCounter returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// increasing float64 measurements once per a measurement collection cycle.
func (f Factory) Float64ObservableCounter(name string, opts ...instrument.Float64ObserverOption) instrument.Float64ObservableCounter {
	i, err := f.mp.Float64ObservableCounter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Float64ObservableUpDownCounter returns a new instrument identified by
// name and configured with options. The instrument is used to
// asynchronously record float64 measurements once per a measurement
// collection cycle.
func (f Factory) Float64ObservableUpDownCounter(name string, opts ...instrument.Float64ObserverOption) instrument.Float64ObservableUpDownCounter {
	i, err := f.mp.Float64ObservableUpDownCounter(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// Float64ObservableGauge returns a new instrument identified by name and
// configured with options. The instrument is used to asynchronously record
// instantaneous float64 measurements once per a measurement collection
// cycle.
func (f Factory) Float64ObservableGauge(name string, opts ...instrument.Float64ObserverOption) instrument.Float64ObservableGauge {
	i, err := f.mp.Float64ObservableGauge(name, opts...)
	if err != nil {
		panic(err)
	}

	return i
}

// RegisterCallback registers f to be called during the collection of a
// measurement cycle.
//
// If Unregister of the returned Registration is called, f needs to be
// unregistered and not called during collection.
//
// The instruments f is registered with are the only instruments that f may
// observe values for.
//
// If no instruments are passed, f should not be registered nor called
// during collection.
func (f Factory) RegisterCallback(c metric.Callback, instruments ...instrument.Asynchronous) metric.Registration {
	r, err := f.mp.RegisterCallback(c, instruments...)
	if err != nil {
		panic(err)
	}

	return r
}

// With creates a Factory using the provided Meter for registration of the
// created Collectors. If the provided Meter is nil, it will panic.
func With(mp metric.Meter) Factory {
	return Factory{mp}
}

func withDefault() Factory {
	return With(global.MeterProvider().Meter("github.com/mahboubii/otelauto"))
}
