package clog

import (
	"context"
	"io"
	"log"
	"log/slog"
	"time"
)

// Types
type (
	Level          = slog.Level
	LevelVar       = slog.LevelVar
	Leveler        = slog.Leveler
	Logger         = slog.Logger
	Attr           = slog.Attr
	Handler        = slog.Handler
	HandlerOptions = slog.HandlerOptions
	JSONHandler    = slog.JSONHandler
	TextHandler    = slog.TextHandler
	Value          = slog.Value
	Record         = slog.Record
	Source         = slog.Source
	Kind           = slog.Kind
	LogValuer      = slog.LogValuer
)

// Constants
const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError

	KindAny       = slog.KindAny
	KindBool      = slog.KindBool
	KindDuration  = slog.KindDuration
	KindFloat64   = slog.KindFloat64
	KindInt64     = slog.KindInt64
	KindString    = slog.KindString
	KindTime      = slog.KindTime
	KindUint64    = slog.KindUint64
	KindGroup     = slog.KindGroup
	KindLogValuer = slog.KindLogValuer

	TimeKey    = slog.TimeKey
	LevelKey   = slog.LevelKey
	MessageKey = slog.MessageKey
	SourceKey  = slog.SourceKey
)

// Variables
var (
	DiscardHandler = slog.DiscardHandler
)

type contextKey struct{}

var defaultLogger *Logger

func init() {
	defaultLogger = slog.Default()
}

// New creates a new Logger with the given Handler.
func New(h Handler) *Logger {
	return slog.New(h)
}

// Default returns the default Logger.
func Default() *Logger {
	return slog.Default()
}

// SetDefault sets the base logger for clog.
func SetDefault(logger *Logger) {
	defaultLogger = logger
	slog.SetDefault(logger)
}

// NewLogLogger returns a new log.Logger that writes to h at the given level.
func NewLogLogger(h Handler, level Level) *log.Logger {
	return slog.NewLogLogger(h, level)
}

// SetLogLoggerLevel sets the level for the bridge between the log and slog packages.
func SetLogLoggerLevel(level Level) Level {
	return slog.SetLogLoggerLevel(level)
}

// WithLogger returns a context with the specified logger.
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func loggerFromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(contextKey{}).(*Logger); ok {
		return logger
	}
	return defaultLogger
}

// Enabled reports whether the logger emits log records at the given level.
func Enabled(ctx context.Context, level Level) bool {
	return loggerFromContext(ctx).Enabled(ctx, level)
}

// Debug logs at [LevelDebug].
func Debug(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).DebugContext(ctx, msg, args...)
}

// Info logs at [LevelInfo].
func Info(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).InfoContext(ctx, msg, args...)
}

// Warn logs at [LevelWarn].
func Warn(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).WarnContext(ctx, msg, args...)
}

// Error logs at [LevelError].
func Error(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).ErrorContext(ctx, msg, args...)
}

// Log emits a log record with the given level and message.
func Log(ctx context.Context, level Level, msg string, args ...any) {
	loggerFromContext(ctx).Log(ctx, level, msg, args...)
}

// LogAttrs is a more efficient version of [Log] that accepts only [Attr].
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	loggerFromContext(ctx).LogAttrs(ctx, level, msg, attrs...)
}

// WithGroup returns a context that starts a group. The keys of all
// attributes added to the [Logger] will be qualified by the given name.
func WithGroup(ctx context.Context, group string) context.Context {
	logger := loggerFromContext(ctx).WithGroup(group)
	return WithLogger(ctx, logger)
}

// WithAttrs returns a context that includes the given attributes
// in each output operation.
func WithAttrs(ctx context.Context, attrs ...Attr) context.Context {
	logger := loggerFromContext(ctx)
	args := make([]any, 0, len(attrs)*2)
	for _, attr := range attrs {
		args = append(args, attr.Key, attr.Value.Any())
	}
	enrichedLogger := logger.With(args...)
	return WithLogger(ctx, enrichedLogger)
}

// With returns a context that includes the given attributes
// in each output operation. Arguments are converted to attributes
// as if by [Logger.Log].
func With(ctx context.Context, args ...any) context.Context {
	logger := loggerFromContext(ctx).With(args...)
	return WithLogger(ctx, logger)
}

// Attr constructors
func Any(key string, value any) Attr                { return slog.Any(key, value) }
func Bool(key string, v bool) Attr                  { return slog.Bool(key, v) }
func Duration(key string, v time.Duration) Attr     { return slog.Duration(key, v) }
func Float64(key string, v float64) Attr            { return slog.Float64(key, v) }
func Group(key string, args ...any) Attr            { return slog.Group(key, args...) }
func GroupAttrs(key string, attrs ...Attr) Attr     { return slog.GroupAttrs(key, attrs...) }
func Int(key string, value int) Attr                { return slog.Int(key, value) }
func Int64(key string, value int64) Attr            { return slog.Int64(key, value) }
func String(key, value string) Attr                 { return slog.String(key, value) }
func Time(key string, v time.Time) Attr             { return slog.Time(key, v) }
func Uint64(key string, v uint64) Attr              { return slog.Uint64(key, v) }

// Handler constructors
func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler {
	return slog.NewJSONHandler(w, opts)
}
func NewTextHandler(w io.Writer, opts *HandlerOptions) *TextHandler {
	return slog.NewTextHandler(w, opts)
}

// Record constructors
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record {
	return slog.NewRecord(t, level, msg, pc)
}

// Value constructors
func AnyValue(v any) Value              { return slog.AnyValue(v) }
func BoolValue(v bool) Value            { return slog.BoolValue(v) }
func DurationValue(v time.Duration) Value { return slog.DurationValue(v) }
func Float64Value(v float64) Value       { return slog.Float64Value(v) }
func GroupValue(as ...Attr) Value       { return slog.GroupValue(as...) }
func Int64Value(v int64) Value          { return slog.Int64Value(v) }
func IntValue(v int) Value              { return slog.IntValue(v) }
func StringValue(value string) Value    { return slog.StringValue(value) }
func TimeValue(v time.Time) Value       { return slog.TimeValue(v) }
func Uint64Value(v uint64) Value        { return slog.Uint64Value(v) }
