package util

import (
    "context"
    "log"
)

type Logger struct{ prefix string }

func New(prefix string) *Logger { return &Logger{prefix: prefix} }

// InfoCtx logs with request scope. Use this in request handlers so logs
// can be correlated by request id (extracted from ctx).
func (l *Logger) InfoCtx(ctx context.Context, msg string) {
    rid, _ := ctx.Value(reqIDKey).(string)
    if rid != "" {
        log.Printf("[%s][%s] %s", l.prefix, rid, msg)
        return
    }
    log.Printf("[%s] %s", l.prefix, msg)
}

// Info is kept for non-request code paths (startup, shutdown).
func (l *Logger) Info(msg string) {
    log.Printf("[%s] %s", l.prefix, msg)
}

type ctxKey int
const reqIDKey ctxKey = 1

func WithRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, reqIDKey, id)
}
