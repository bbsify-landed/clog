package clog

import (
	"context"
	"log/slog"
)

type contextKey struct{}

var defaultLogger *slog.Logger

func init() {
	defaultLogger = slog.Default()
}

// SetDefault sets the base logger for clog.
func SetDefault(logger *slog.Logger) {
	defaultLogger = logger
}

// WithLogger returns a context with the specified logger.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func loggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(contextKey{}).(*slog.Logger); ok {
		return logger
	}
	return defaultLogger
}

// Enabled reports whether the logger emits log records at the given level.
func Enabled(ctx context.Context, level slog.Level) bool {
	return loggerFromContext(ctx).Enabled(ctx, level)
}

// Debug logs at [slog.LevelDebug].
func Debug(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).DebugContext(ctx, msg, args...)
}

// Info logs at [slog.LevelInfo].
func Info(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).InfoContext(ctx, msg, args...)
}

// Warn logs at [slog.LevelWarn].
func Warn(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).WarnContext(ctx, msg, args...)
}

// Error logs at [slog.LevelError].
func Error(ctx context.Context, msg string, args ...any) {
	loggerFromContext(ctx).ErrorContext(ctx, msg, args...)
}

// Log emits a log record with the given level and message.
func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	loggerFromContext(ctx).Log(ctx, level, msg, args...)
}

// LogAttrs is a more efficient version of [Log] that accepts only [slog.Attr].
func LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	loggerFromContext(ctx).LogAttrs(ctx, level, msg, attrs...)
}

// WithGroup returns a context that starts a group. The keys of all
// attributes added to the [slog.Logger] will be qualified by the given name.
func WithGroup(ctx context.Context, group string) context.Context {
	logger := loggerFromContext(ctx).WithGroup(group)
	return WithLogger(ctx, logger)
}

// WithAttrs returns a context that includes the given attributes
// in each output operation.
func WithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
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
// as if by [slog.Logger.Log].
func With(ctx context.Context, args ...any) context.Context {
	logger := loggerFromContext(ctx).With(args...)
	return WithLogger(ctx, logger)
}
