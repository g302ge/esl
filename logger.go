package esl

import "log"

var Logger Log = new(defaultLogger)

// Log interface for record
// suggest using the logrus lib
// e.g. esl.Logger = logrus.New()
type Log interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

// default implements the STDOUT output
type defaultLogger struct {
	log.Logger
}

func (l *defaultLogger) Debug(args ...interface{}) {
	l.Print(args...)
}

func (l *defaultLogger) Debugf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *defaultLogger) Error(args ...interface{}) {
	l.Print(args...)
}

func (l *defaultLogger) Errorf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *defaultLogger) Warn(args ...interface{}) {
	l.Print(args...)
}

func (l *defaultLogger) Warnf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

// helper functions used in current namespace to log some useful informations

func debug(args ...interface{}) {
	Logger.Debug(args...)
}

func debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

func errorm(args ...interface{}) {
	Logger.Error(args...)
}

func errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

func warn(args ...interface{}) {
	Logger.Warn(args...)
}

func warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}
