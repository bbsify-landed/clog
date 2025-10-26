package clog_test

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"testing"

	"github.com/bbsify-landed/clog"
)

func TestSetDefault(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Set the default logger
	clog.SetDefault(logger)

	// Test that the default logger is used
	ctx := context.Background()
	clog.Info(ctx, "test message")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("Expected log output to contain 'test message', got: %s", output)
	}
}

func TestWithLogger(t *testing.T) {
	// Create two different loggers with different prefixes
	var buf1, buf2 bytes.Buffer
	logger1 := slog.New(slog.NewTextHandler(&buf1, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger2 := slog.New(slog.NewTextHandler(&buf2, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Set default logger
	clog.SetDefault(logger1)

	// Create context with specific logger
	ctx := context.Background()
	ctxWithLogger := clog.WithLogger(ctx, logger2)

	// Log with default logger
	clog.Info(ctx, "default logger message")

	// Log with context logger
	clog.Info(ctxWithLogger, "context logger message")

	// Check that messages went to correct loggers
	if !strings.Contains(buf1.String(), "default logger message") {
		t.Errorf("Expected buf1 to contain 'default logger message', got: %s", buf1.String())
	}

	if !strings.Contains(buf2.String(), "context logger message") {
		t.Errorf("Expected buf2 to contain 'context logger message', got: %s", buf2.String())
	}

	// Ensure messages didn't go to wrong loggers
	if strings.Contains(buf1.String(), "context logger message") {
		t.Error("buf1 should not contain 'context logger message'")
	}

	if strings.Contains(buf2.String(), "default logger message") {
		t.Error("buf2 should not contain 'default logger message'")
	}
}

func TestEnabled(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelWarn}))

	ctx := clog.WithLogger(context.Background(), logger)

	// Test that debug level is not enabled when logger is set to warn level
	if clog.Enabled(ctx, slog.LevelDebug) {
		t.Error("Expected debug level to be disabled")
	}

	// Test that warn level is enabled
	if !clog.Enabled(ctx, slog.LevelWarn) {
		t.Error("Expected warn level to be enabled")
	}

	// Test that error level is enabled
	if !clog.Enabled(ctx, slog.LevelError) {
		t.Error("Expected error level to be enabled")
	}
}

func TestLogLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := clog.WithLogger(context.Background(), logger)

	// Test all log levels
	clog.Debug(ctx, "debug message", "key", "value")
	clog.Info(ctx, "info message", "key", "value")
	clog.Warn(ctx, "warn message", "key", "value")
	clog.Error(ctx, "error message", "key", "value")

	output := buf.String()

	// Check that all messages are present
	expectedMessages := []string{"debug message", "info message", "warn message", "error message"}
	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Expected output to contain '%s', got: %s", msg, output)
		}
	}
}

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := clog.WithLogger(context.Background(), logger)

	// Test generic Log function
	clog.Log(ctx, slog.LevelInfo, "generic log message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, "generic log message") {
		t.Errorf("Expected output to contain 'generic log message', got: %s", output)
	}
}

func TestLogAttrs(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := clog.WithLogger(context.Background(), logger)

	// Test LogAttrs function
	attrs := []slog.Attr{
		slog.String("key1", "value1"),
		slog.Int("key2", 42),
	}
	clog.LogAttrs(ctx, slog.LevelInfo, "attrs message", attrs...)

	output := buf.String()
	if !strings.Contains(output, "attrs message") {
		t.Errorf("Expected output to contain 'attrs message', got: %s", output)
	}
	if !strings.Contains(output, "key1=value1") {
		t.Errorf("Expected output to contain 'key1=value1', got: %s", output)
	}
	if !strings.Contains(output, "key2=42") {
		t.Errorf("Expected output to contain 'key2=42', got: %s", output)
	}
}

func TestWithGroup(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := clog.WithLogger(context.Background(), logger)
	ctxWithGroup := clog.WithGroup(ctx, "testgroup")

	// Log with group
	clog.Info(ctxWithGroup, "grouped message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, "grouped message") {
		t.Errorf("Expected output to contain 'grouped message', got: %s", output)
	}
	// The exact format depends on the handler, but should contain the group
	if !strings.Contains(output, "testgroup") {
		t.Errorf("Expected output to contain 'testgroup', got: %s", output)
	}
}

func TestWithAttrs(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := clog.WithLogger(context.Background(), logger)
	attrs := []slog.Attr{
		slog.String("service", "test"),
		slog.Int("version", 1),
	}
	ctxWithAttrs := clog.WithAttrs(ctx, attrs...)

	// Log with pre-set attributes
	clog.Info(ctxWithAttrs, "message with attrs")

	output := buf.String()
	if !strings.Contains(output, "message with attrs") {
		t.Errorf("Expected output to contain 'message with attrs', got: %s", output)
	}
	if !strings.Contains(output, "service=test") {
		t.Errorf("Expected output to contain 'service=test', got: %s", output)
	}
	if !strings.Contains(output, "version=1") {
		t.Errorf("Expected output to contain 'version=1', got: %s", output)
	}
}

func TestWith(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := clog.WithLogger(context.Background(), logger)
	ctxWith := clog.With(ctx, "user", "john", "session", "abc123")

	// Log with pre-set key-value pairs
	clog.Info(ctxWith, "user action")

	output := buf.String()
	if !strings.Contains(output, "user action") {
		t.Errorf("Expected output to contain 'user action', got: %s", output)
	}
	if !strings.Contains(output, "user=john") {
		t.Errorf("Expected output to contain 'user=john', got: %s", output)
	}
	if !strings.Contains(output, "session=abc123") {
		t.Errorf("Expected output to contain 'session=abc123', got: %s", output)
	}
}

func TestContextPropagation(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Test that context values are preserved through multiple operations
	ctx := context.Background()
	ctx = clog.WithLogger(ctx, logger)
	ctx = clog.With(ctx, "request_id", "req-123")
	ctx = clog.WithGroup(ctx, "api")

	clog.Info(ctx, "processing request")

	output := buf.String()
	if !strings.Contains(output, "processing request") {
		t.Errorf("Expected output to contain 'processing request', got: %s", output)
	}
	if !strings.Contains(output, "request_id=req-123") {
		t.Errorf("Expected output to contain 'request_id=req-123', got: %s", output)
	}
}

// Example demonstrates using different loggers in different contexts
func Example() {
	// Create a buffer to capture output for this example
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	// Create contexts with different loggers
	ctx := context.Background()
	ctx = clog.WithLogger(ctx, logger)

	// Log to different contexts
	clog.Info(ctx, "Application started", "service", "api")

	// Create a request context with additional attributes
	requestCtx := clog.With(ctx, "request_id", "req-123", "user_id", "user-456")
	clog.Info(requestCtx, "Processing user request")

	// Create a database context with a group
	dbCtx := clog.WithGroup(requestCtx, "database")
	clog.Info(dbCtx, "Query executed", "table", "users", "duration_ms", 45)

	fmt.Println("Structured logging with context-aware clog package")
	// Output: Structured logging with context-aware clog package
}
