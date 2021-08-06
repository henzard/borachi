// +build android

package gui

import (
	"os"
	"os/signal"
	"syscall"

	"git.mrcyjanek.net/mrcyjanek/borachi/webui"
)

func Start() {
	go webui.Run()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
