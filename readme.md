# logger
- Based on the standard library log/slog encapsulation, in order to better customize slog configuration. 
- The usage method is the same as the official Slog library of Go.
- Log level: DEBUG,INFO,WARN,ERROR

# usage
import logger package
```go
import(
	"github.com/daheige/slogger"
)
```

## output to stdout with text
````go
slogger.Init()
// slogger.Init(WithAddSource(true))
slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
slog.Info("hello", "a", 1, "name", "go")
slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
slog.Warn("foo waring", "a", 1, "name", "go", slog.String("user", "coco"))
````

## output to stdout with json
```go
// slogger.Init(WithEnableJSON())
slogger.Init(WithEnableJSON(), WithAddSource(true))
slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
slog.Info("hello", "a", 1, "name", "go")
slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
```

## output to file with text
```go
slogger.Init(
    WithAddSource(true), WithOutputToFile(true),
    WithFileName("app.log"),
)

slog.Debug("debug abc", "a", 1, "b", "234abc") // no log to output

slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
slog.Info("hello", "a", 1, "name", "go")
slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
slog.Warn("foo waring", "a", 1, "name", "go", slog.String("user", "coco"))
```

## output to file with json
```go
slogger.Init(
    WithEnableJSON(), WithAddSource(true), WithOutputToFile(true),
    WithFileName("app.log"),
)

slog.Debug("debug abc", "a", 1, "b", "234abc") // no log to output

slog.Info("hello", slog.String("name", "slog"), slog.Int("number", 10))
slog.Info("hello", "a", 1, "name", "go")
slog.Info("hello", "a", 1, "name", "go", slog.String("user", "coco"))
slog.Warn("foo waring", "a", 1, "name", "go", slog.String("user", "coco"))
```
