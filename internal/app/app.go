package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

const Name = "chat-service"

type application struct {
	Ctx        context.Context
	cancelFunc context.CancelFunc
}

var (
	A *application
)

func init() {
	A = &application{}
}

func WithGracefulShutdown() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	A.Ctx, A.cancelFunc = context.WithCancel(context.Background())
}

func Wait() {
	defer A.cancelFunc()
	<-A.Ctx.Done()
}
