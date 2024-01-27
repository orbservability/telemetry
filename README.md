# Telemetry

Common configuration and uses for logs, metrics, and traces.

## Installation

```go
require (
  "github.com/orbservability/telemetry v0.0.1"
)
```

## Usage

### Global

Global configuration is automatically provided via `init()` functions, provided that the package is imported.

### Server

Provides common tooling for the following use cases:

- gRPC

  ```go
  import (
    "github.com/orbservability/telemetry/pkg/logs"
    "github.com/orbservability/telemetry/pkg/metrics"
    "github.com/orbservability/telemetry/pkg/traces"
    "google.golang.org/grpc"
  )

  grpcServer := grpc.NewServer(
    grpc.ChainUnaryInterceptor(traces.UnaryServerInterceptor, logs.UnaryServerInterceptor, metrics.UnaryServerInterceptor),
    grpc.ChainStreamInterceptor(traces.StreamServerInterceptor, logs.StreamServerInterceptor, metrics.StreamServerInterceptor),
    // ...
  )
  ```

### Client

Provides common tooling for the following use cases:

- gRPC

  ```go
  import (
    "github.com/orbservability/telemetry/pkg/logs"
    "github.com/orbservability/telemetry/pkg/metrics"
    "github.com/orbservability/telemetry/pkg/traces"
    "google.golang.org/grpc"
  )

  conn, err := grpc.Dial(
    "https://api.orbservability.com",
    grpc.WithChainUnaryInterceptor(traces.UnaryClientInterceptor, logs.UnaryClientInterceptor, metrics.UnaryClientInterceptor),
    grpc.WithChainStreamInterceptor(traces.StreamClientInterceptor, logs.StreamClientInterceptor, metrics.StreamClientInterceptor),
    // ...
  )
  ```
