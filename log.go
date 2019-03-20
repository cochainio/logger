package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Singleton *zap.SugaredLogger
var PlainSingleton *zap.Logger

func Instantiate(logLevel string, disableStacktrace bool) {
	if Singleton != nil || PlainSingleton != nil {
		panic("logger has been instantiated")
	}

	var level zapcore.Level
	err := level.Set(logLevel)
	if err != nil {
		panic(err.Error())
	}

	var conf zap.Config
	if level == zapcore.DebugLevel {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}
	conf.Level.SetLevel(level)
	conf.DisableStacktrace = disableStacktrace

	PlainSingleton, err = conf.Build()
	if err != nil {
		panic(err.Error())
	}
	Singleton = PlainSingleton.Sugar()
}

type SugaredLogger struct {
	*zap.SugaredLogger
}

func Named(name string) *SugaredLogger {
	return &SugaredLogger{Singleton.Named(name)}
}

func (l *SugaredLogger) Event(event string) *zap.SugaredLogger {
	return l.With("event", event)
}

func (l *SugaredLogger) Context(context string) *zap.SugaredLogger {
	return l.With("context", context)
}

func Debug(args ...interface{}) {
	Singleton.Debug(args...)
}

func Info(args ...interface{}) {
	Singleton.Info(args...)
}

func Warn(args ...interface{}) {
	Singleton.Warn(args...)
}

func Error(args ...interface{}) {
	Singleton.Error(args...)
}

func DPanic(args ...interface{}) {
	Singleton.DPanic(args...)
}

func Panic(args ...interface{}) {
	Singleton.Panic(args...)
}

func Fatal(args ...interface{}) {
	Singleton.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	Singleton.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	Singleton.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	Singleton.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	Singleton.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	Singleton.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	Singleton.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Singleton.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Singleton.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Singleton.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Singleton.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Singleton.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	Singleton.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	Singleton.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Singleton.Fatalw(msg, keysAndValues...)
}

func Sync() error {
	return Singleton.Sync()
}
