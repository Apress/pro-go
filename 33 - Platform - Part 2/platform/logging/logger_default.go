package logging

import (
    "log"
    "fmt"
)

type DefaultLogger struct {
    minLevel LogLevel
    loggers map[LogLevel]*log.Logger
    triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
    return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
    if (l.minLevel <= level) {
        l.loggers[level].Output(2, message)
    }
}

func (l *DefaultLogger) Trace(msg string) {
    l.write(Trace, msg)
}

func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
    l.write(Trace, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Debug(msg string) {
    l.write(Debug, msg)
}

func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
    l.write(Debug, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Info(msg string) {
    l.write(Information, msg)
}

func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
    l.write(Information, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Warn(msg string) {
    l.write(Warning, msg)
}

func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
    l.write(Warning, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Panic(msg string) {
    l.write(Fatal, msg)
    if (l.triggerPanic) {
        panic(msg)
    }
}

func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
    formattedMsg := fmt.Sprintf(template, vals...)
    l.write(Fatal, formattedMsg)
    if (l.triggerPanic) {
        panic(formattedMsg)
    }
}
