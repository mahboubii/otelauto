# OTELAuto

[![ci](https://github.com/mahboubii/otelauto/actions/workflows/workflow.yaml/badge.svg?branch=main)](https://github.com/mahboubii/otelauto/actions/workflows/workflow.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mahboubii/otelauto)](https://goreportcard.com/report/github.com/mahboubii/otelauto)
[![Documentation](https://godoc.org/github.com/mahboubii/otelauto?status.svg)](https://pkg.go.dev/mod/github.com/mahboubii/otelauto)

An OpenTelemetry (OTel) metric instrumentation helper inspired by prometheus [promauto](github.com/prometheus/client_golang/prometheus/promauto).

## Install

```bash
$ go get github.com/mahboubii/otelauto
```

## Usage

This library implements all instrument types from the OpenTelemetry [Meter](https://pkg.go.dev/go.opentelemetry.io/otel/metric#Meter) interface.

```go
package main

import (
  "context"

  "github.com/mahboubii/otelauto"
)

var counter = otelauto.Int64Counter("my.counter")

func main() {
  counter.Add(context.Background(), 1)
}
```

Alternatively, you can provide custom meters:

```go

import (
  "context"

  "github.com/mahboubii/otelauto"
  "go.opentelemetry.io/otel/metric"
)

var counter = otelauto.With(metric.NewNoopMeter()).Int64Counter("my.counter")

func main() {
  counter.Add(context.Background(), 1)
}
```
