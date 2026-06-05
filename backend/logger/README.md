# Logger

Structured logging powered by `go.uber.org/zap`.

Initialized once at startup in `main.go`. All server messages use the zap sugared logger.

## Behavior

| `ENV` value | Encoding | Format |
|-------------|----------|--------|
| `production` | JSON | `zap.NewProduction()` |
| anything else | Console | Human-readable with ISO8601 timestamps |

## Usage

```go
import "backend/logger"

logger.InitLogger()
sugar := logger.GetLogger().Sugar()

sugar.Info("message")
sugar.Infof("formatted %s", "message")
sugar.Warnf("warning: %v", err)
sugar.Fatalf("fatal: %v", err)  // calls os.Exit(1)
```
