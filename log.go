package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Singleton *zap.SugaredLogger
var PlainSingleton *zap.Logger

var level zapcore.Level
var internal *zap.SugaredLogger

func Instantiate(logLevel string, disableStacktrace bool) {
	if Singleton != nil || PlainSingleton != nil {
		panic("logger has been instantiated")
	}

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
	internal = PlainSingleton.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func Level() zapcore.Level {
	return level
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
	internal.Debug(args...)
}

func Info(args ...interface{}) {
	internal.Info(args...)
}

func Warn(args ...interface{}) {
	internal.Warn(args...)
}

func Error(args ...interface{}) {
	internal.Error(args...)
}

func DPanic(args ...interface{}) {
	internal.DPanic(args...)
}

func Panic(args ...interface{}) {
	internal.Panic(args...)
}

func Fatal(args ...interface{}) {
	internal.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	internal.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	internal.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	internal.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	internal.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	internal.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	internal.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	internal.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	internal.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	internal.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	internal.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	internal.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	internal.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	internal.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	internal.Fatalw(msg, keysAndValues...)
}

func Sync() error {
	return internal.Sync()
}
