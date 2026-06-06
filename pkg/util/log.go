package util

import "log"

type Logger struct{ prefix string }

func New(prefix string) *Logger { return &Logger{prefix: prefix} }

func (l *Logger) Info(msg string) {
    log.Printf("[%s] %s", l.prefix, msg)
}
