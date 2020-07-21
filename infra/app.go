package infra

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	errC chan error
	Ctx  context.Context

	GRPCAddr string
	httpAddr string

	GRPCServer *grpc.Server

	ServeMux    *runtime.ServeMux
	DialOptions []grpc.DialOption
}

func NewApp() *App {
	// 解配置文件
	parse()

	grpcAddr := fmt.Sprintf(":%v", Config.Server.Port.GRPC)
	httpAddr := fmt.Sprintf(":%v", Config.Server.Port.HTTP)

	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(RecoveryInterceptor, LoggingInterceptor, ValidateInterceptor))

	serveMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	//ServeMux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	return &App{
		errC:     make(chan error),
		httpAddr: httpAddr,

		GRPCAddr:    grpcAddr,
		Ctx:         context.Background(),
		GRPCServer:  grpcServer,
		ServeMux:    serveMux,
		DialOptions: opts,
	}
}

func (app *App) Run() error {
	app.start()

	err := <-app.errC

	app.stop()

	return err
}

func (app *App) start() {
	glog.Info("start()")

	var g errgroup.Group

	// grpc
	g.Go(func() error {
		glog.Infof("Listening and serving GRPC on: %v", app.GRPCAddr)
		listener, err := net.Listen("tcp", app.GRPCAddr)
		if err != nil {
			return err
		}
		return app.GRPCServer.Serve(listener)
	})

	// http
	g.Go(func() error {
		// Start HTTP server (and proxy calls to gRPC server endpoint)
		glog.Infof("Listening and serving HTTP on: %v", app.httpAddr)
		return http.ListenAndServe(app.httpAddr, app.ServeMux)
	})

	go func() {
		app.errC <- g.Wait()
	}()

	go func() {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		sig := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)

		err := fmt.Errorf("signal: %v", <-sig)
		glog.Info(err)

		app.errC <- err
	}()
}

func (app *App) stop() {
	glog.Info("stop()")
	app.GRPCServer.Stop()
}
