# clog

A context-aware structured logging package for Go, built on top of `slog`.

clog re-exports the entire `slog` API so you can use it as a drop-in replacement — no need to import `log/slog` separately.

## Installation

```bash
go get github.com/bbsify-landed/clog
```

## Usage

### Basic logging

```go
ctx := context.Background()
clog.Info(ctx, "hello", "user", "alice")
```

### Setting up a default logger

```go
logger := clog.New(clog.NewJSONHandler(os.Stdout, &clog.HandlerOptions{
    Level: clog.LevelDebug,
}))
clog.SetDefault(logger)

ctx := context.Background()
clog.Info(ctx, "server started", "addr", ":8080")
```

### Adding request-scoped attributes

```go
ctx = clog.With(ctx, "request_id", "abc-123", "user_id", "user-456")
clog.Info(ctx, "processing request")
// Output includes: request_id=abc-123 user_id=user-456
```

### Grouping attributes

```go
dbCtx := clog.WithGroup(ctx, "db")
clog.Info(dbCtx, "query executed", "table", "users", "duration_ms", 12)
// Output includes: db.table=users db.duration_ms=12
```

### Using typed attributes

```go
ctx = clog.WithAttrs(ctx,
    clog.String("service", "billing"),
    clog.Int("version", 3),
)
clog.Info(ctx, "invoice created")
```

### Per-context loggers

```go
debugLogger := clog.New(clog.NewTextHandler(os.Stderr, &clog.HandlerOptions{
    Level: clog.LevelDebug,
}))
ctx = clog.WithLogger(ctx, debugLogger)
clog.Debug(ctx, "verbose output only for this context")
```

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.
