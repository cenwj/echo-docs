package main

import (
	"echo-docs/conf"
	"echo-docs/db"
	"echo-docs/route"
	"log"
	"os"
	"os/signal"
	"syscall"
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
