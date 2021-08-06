// +build !android

package gui

import (
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/borachi/webui"
	"github.com/zserge/lorca"
)

func Start() {
	ui, _ := lorca.New("", "", 480, 630)
	go webui.Run()
	ui.Eval(`window.location.href = "http://127.0.0.1:` + strconv.Itoa(webui.Port) + `/setup.html"`)
	<-ui.Done()
}
