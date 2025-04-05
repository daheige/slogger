package logger

import (
	"log/slog"
	"testing"
)

/*
=== RUN   TestSlogOutputJson
{
    "time": "2025-04-05T21:56:06.001346+08:00",
    "level": "INFO",
    "source": {
        "function": "github.com/daheige/logger.TestSlogOutputJson",
        "file": "/Users/amanda/web/go/logger/slog_test.go",
        "line": 11
    },
    "msg": "hello",
    "name": "slog",
    "number": 10
}
-- PASS: TestSlogOutputJson (0.00s)
PASS
*/
func TestSlogOutputJson(t *testing.T) {
	// Init(WithEnableJSON())
	Init(WithEnableJSON(), WithAddSource(true))
	slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
	slog.Info("hello", "a", 1, "name", "go")
	slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
}

/**
=== RUN   TestSlogOutputText
time=2025-04-05T21:55:35.231+08:00 level=INFO msg=hello name=slog number=10
time=2025-04-05T21:55:35.231+08:00 level=INFO msg=hello a=1 name=go
time=2025-04-05T21:55:35.231+08:00 level=INFO msg=hello a=1 name=go user=coco
time=2025-04-05T21:55:35.231+08:00 level=WARN msg="foo waring" a=1 name=go user=coco
--- PASS: TestSlogOutputText (0.00s)
PASS
*/
func TestSlogOutputText(t *testing.T) {
	// output to text
	Init()
	// Init(WithAddSource(true))
	slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
	slog.Info("hello", "a", 1, "name", "go")
	slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
	slog.Warn("foo waring", "a", 1, "name", "go", slog.String("user", "coco"))
}

func TestSlogOutputFile(t *testing.T) {
	// log to file with json
	Init(
		WithEnableJSON(), WithAddSource(true), WithOutputToFile(true),
		WithFileName("app.log"),
	)

	slog.Debug("debug abc", "a", 1, "b", "234abc") // no log to output

	slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
	slog.Info("hello", "a", 1, "name", "go")
	slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
	slog.Warn("foo waring", "a", 1, "name", "go", slog.String("user", "coco"))
}
