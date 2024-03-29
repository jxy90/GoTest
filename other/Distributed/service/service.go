package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = net.JoinHostPort(host, port)

	go func() {
		log.fmt.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v start.press any key to stop \n", serviceName)
		var s string
		fmt.Scan(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
