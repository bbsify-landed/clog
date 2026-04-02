# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com),
and this project adheres to [Semantic Versioning](https://semver.org).

## [v1.1.0] - 2026-04-01

clog now re-exports the entire `slog` ecosystem. You no longer need to import `log/slog` alongside clog.

### Added

- Type aliases for all core slog types: `Level`, `LevelVar`, `Leveler`, `Logger`, `Attr`, `Handler`, `HandlerOptions`, `JSONHandler`, `TextHandler`, `Value`, `Record`, `Source`, `Kind`, `LogValuer`
- Log level constants: `LevelDebug`, `LevelInfo`, `LevelWarn`, `LevelError`
- Value kind constants: `KindAny`, `KindBool`, `KindDuration`, `KindFloat64`, `KindInt64`, `KindString`, `KindTime`, `KindUint64`, `KindGroup`, `KindLogValuer`
- Standard key constants: `TimeKey`, `LevelKey`, `MessageKey`, `SourceKey`
- Logger constructors: `New`, `Default`, `NewLogLogger`, `SetLogLoggerLevel`
- Handler constructors: `NewJSONHandler`, `NewTextHandler`
- Attr constructors: `Any`, `Bool`, `Duration`, `Float64`, `Group`, `GroupAttrs`, `Int`, `Int64`, `String`, `Time`, `Uint64`
- Value constructors: `AnyValue`, `BoolValue`, `DurationValue`, `Float64Value`, `GroupValue`, `Int64Value`, `IntValue`, `StringValue`, `TimeValue`, `Uint64Value`
- `DiscardHandler` variable for discarding log output

### Changed

- `SetDefault` now also calls `slog.SetDefault`, keeping clog and slog in sync
- Doc comments reference clog types instead of `slog` types

## [v1.0.1] - 2025-12-15

### Changed

- Example uses `SetDefault` instead of `WithLogger` for initial setup
- Example output is now verified by `go test`

## [v1.0.0] - 2025-12-01

### Added

- Context-aware logging functions: `Debug`, `Info`, `Warn`, `Error`, `Log`, `LogAttrs`
- Context enrichment: `WithLogger`, `WithGroup`, `WithAttrs`, `With`
- `Enabled` to check if a log level is active
- `SetDefault` to set the package-wide default logger

[v1.1.0]: https://github.com/bbsify-landed/clog/compare/v1.0.1...v1.1.0
[v1.0.1]: https://github.com/bbsify-landed/clog/compare/v1.0.0...v1.0.1
[v1.0.0]: https://github.com/bbsify-landed/clog/releases/tag/v1.0.0
