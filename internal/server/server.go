package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	cfg                *config.Config
	srv                *http.Server
	ctx, cancel        = context.WithCancel(context.Background())
	defaultQuitTimeout = 15 * time.Second
	cleaner            *cleaner2.CleanHandle
)

func parse() {
	cfg = config.Cfg

	logger.Init(cfg.Log)
	mdb.Init(cfg.Mysql, cfg.Gorm)
	mongodb.Init(cfg.Mongo)
	redis.Init(cfg.Redis)
	minio.Init(cfg.Minio)
	// some init
	// dao.Init(ctx, cfg)

}

func runJobs() {
	cleaner = cleaner2.NewCleaner(ctx, cfg.TmpDir,
		time.Duration(cfg.TmpCleanInterval)*time.Second, time.Duration(cfg.TmpCleanTimeLimit)*time.Second)
	cleaner.Run()
}

func run() {
	app := gin.New()
	router.SetMiddleware(app)
	rt := pkgRouter.NewRouter(app)
	rt.AddRegisters(appRouter.Prefix, appRouter.Registers)
	rt.SetPublicHost(config.ModifyHttpHost(config.Cfg.Env.AppServerPublicHost))
	rt.Register()
	router.SetRouter(rt)
	srv = &http.Server{
		Addr:         cfg.Listen,
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		_ = srv.ListenAndServe()
	}()
	logger.Infof("listening on %s", cfg.Listen)
	<-ctx.Done()
}

func stop() error {
	_ctx, _cancel := context.WithCancel(ctx)
	done := make(chan bool)
	go func() {
		_ = srv.Shutdown(_ctx)
		close(done)
	}()
	defer _cancel()
	select {
	case <-done:
		return nil
	case <-time.NewTimer(defaultQuitTimeout).C:
		return errors.New("shutdown time out")
	}
}

// nolint:revive
func Run() {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("[main] recover painc:%s. stack:%s", r, string(buf))
		}
	}()
	parse()
	runJobs()
	run()
}

// nolint:revive
func Stop() {
	_ = stop()
	cancel()
}
