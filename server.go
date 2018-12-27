package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cenwj/echo-docs/conf"
	"github.com/cenwj/echo-docs/db"
	"github.com/cenwj/echo-docs/route"
)

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			conf.ReloadConfig()
			log.Println("Reload config")
		}
	}()

	e := route.Init()

	db.Init()

	db.RedisConn(0)

	e.Logger.Fatal(e.Start(":" + conf.Config().Port))
}
