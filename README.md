# clog

A context-aware structured logging wrapper around Go's `slog` package.

## Installation

```bash
go get github.com/bbsify-landed/clog
```

## Usage

```go
ctx := context.Background()
clog.Info(ctx, "hello", "user", "alice")
```
