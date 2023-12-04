package main

import (
	"capture/internal/app"
	"capture/internal/conf"
	"capture/internal/logger"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	received chan os.Signal = nil
)

func init() {
	// Handle signals:
	//
	// - SIGUSR1
	// - SIGUSR2
	// - SIGINT
	received = make(chan os.Signal, 1)
	initSigHandler()

	// Initial logger, changes when configuration is loaded
	app.Log = logger.NewSyslog("info", app.Name, true)
	log.SetOutput(app.Log.Writer())

	// Load configuration
	// Can override values using flags
	app.Cfg = conf.New(app.Name)

	// Create Logger as it is defined in configuration
	app.Log = logger.NewFromType(
		app.Cfg.LogType,
		app.Cfg.LogFile,
		app.Name,
		app.Cfg.LogLevel,
		app.Cfg.LogUseTag,
		app.Cfg.LogUseColor,
	)
	log.SetOutput(app.Log.Writer())

	// Started
	app.Log.Info(app.Name + " v" + app.Version + " started successfully")
}

func initSigHandler() {
	signal.Notify(received, os.Interrupt, syscall.SIGUSR1, syscall.SIGUSR2)

	go func(r chan os.Signal) {
		for {
			s := <-r

			switch s {
			// SIGUSR1	 10
			case syscall.SIGUSR1:
				app.PrintStats()
			// SIGUSR2	 12
			case syscall.SIGUSR2:
				app.WriteStats()
			default:
				fmt.Println("Received CTRL-C. Exiting.")
				app.Log.Alert("Received CTRL-C. Exiting.")
				os.Exit(1)
			}
		}
	}(received)
}
