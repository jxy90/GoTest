package main

import (
	"context"
	"fmt"
	"github.com/jxy90/GoTest/other/distributed/log"
	"github.com/jxy90/GoTest/other/distributed/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(context.Background(), "Log Service", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
