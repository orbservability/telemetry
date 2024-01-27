# Telemetry

Common configuration and uses for logs, metrics, and traces.

## Usage

### Server

- gRPC

  ```go
  import (
    "google.golang.org/grpc"

    "github.com/orbservability/telemetry/pkg/logs"
    "github.com/orbservability/telemetry/pkg/metrics"
    "github.com/orbservability/telemetry/pkg/traces"
  )

  grpcServer := grpc.NewServer(
    grpc.ChainUnaryInterceptor(traces.UnaryServerInterceptor, logs.UnaryServerInterceptor, metrics.UnaryServerInterceptor),
    grpc.ChainStreamInterceptor(traces.StreamServerInterceptor, logs.StreamServerInterceptor, metrics.StreamServerInterceptor),
  )
  ```

### Client

- gRPC

  ```go
  import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    "github.com/orbservability/telemetry/pkg/logs"
    "github.com/orbservability/telemetry/pkg/metrics"
    "github.com/orbservability/telemetry/pkg/traces"
  )

  conn, err := grpc.Dial(
    "https://api.orbservability.com",
    grpc.WithChainUnaryInterceptor(traces.UnaryClientInterceptor, logs.UnaryClientInterceptor, metrics.UnaryClientInterceptor),
    grpc.WithChainStreamInterceptor(traces.StreamClientInterceptor, logs.StreamClientInterceptor, metrics.StreamClientInterceptor),
    grpc.WithTransportCredentials(insecure.NewCredentials()),
  )
  ```
