package logger

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/kazegusuri/grpc-panic-handler"
	"go.uber.org/zap"
	"google.golang.org/grpc/grpclog"
)

func InstantiateGRPC() {
	if grpcSingleton != nil {
		panic("grpc logger has been instantiated")
	}

	grpcSingleton = &GRPCLogger{
		log: PlainSingleton.WithOptions(zap.AddCallerSkip(2)).Sugar().Named("rpc"),
	}
	grpclog.SetLoggerV2(grpcSingleton)
	panichandler.InstallPanicHandler(grpcSingleton.LogPanicHandler)
}

var grpcSingleton *GRPCLogger

type GRPCLogger struct {
	log *zap.SugaredLogger
}

func (g *GRPCLogger) LogPanicHandler(r interface{}) {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("Recovered from panic: %#v (%v)", r, r))
	for i := 0; true; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %d: %v:%v\n", i, file, line))
	}
	g.log.Error(buffer.String())
}

func (g *GRPCLogger) Info(args ...interface{}) {
	g.log.Info(args...)
}

func (g *GRPCLogger) Infoln(args ...interface{}) {
	g.log.Info(args...)
}

func (g *GRPCLogger) Infof(format string, args ...interface{}) {
	g.log.Infof(format, args...)
}

func (g *GRPCLogger) Warning(args ...interface{}) {
	g.log.Warn(args...)
}

func (g *GRPCLogger) Warningln(args ...interface{}) {
	g.log.Warn(args...)
}

func (g *GRPCLogger) Warningf(format string, args ...interface{}) {
	g.log.Warnf(format, args...)
}

func (g *GRPCLogger) Error(args ...interface{}) {
	g.log.Error(args...)
}

func (g *GRPCLogger) Errorln(args ...interface{}) {
	g.log.Error(args)
}

func (g *GRPCLogger) Errorf(format string, args ...interface{}) {
	g.log.Errorf(format, args...)
}

func (g *GRPCLogger) Fatal(args ...interface{}) {
	g.log.Fatal(args...)
}

func (g *GRPCLogger) Fatalln(args ...interface{}) {
	g.log.Fatal(args...)
}

func (g *GRPCLogger) Fatalf(format string, args ...interface{}) {
	g.log.Fatalf(format, args...)
}

func (g *GRPCLogger) V(l int) bool {
	return l <= 2 // verbosity level is 2
}
