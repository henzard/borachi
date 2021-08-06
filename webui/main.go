package webui

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	bore "github.com/jkuri/bore/client"
	"golang.org/x/net/html"
)

//go:embed template.html
var template string

//go:embed setup.html
var setup []byte

var Port = 3414

var client bore.BoreClient
var serr string

func Run() {
	http.HandleFunc("/setup.html", func(rw http.ResponseWriter, r *http.Request) { rw.Write(setup) })
	http.HandleFunc("/run", run)
	http.HandleFunc("/info", info)
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(Port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func message(title string, content string) []byte {
	return []byte(fmt.Sprintf(template, title, content, "%"))
}

func run(rw http.ResponseWriter, r *http.Request) {
	rs := strings.Split(r.PostFormValue("RemoteServer"), ":")
	if len(rs) != 2 {
		rw.Write(message("Starting...", `Remote server should be "host:port" <a href="/setup.html">go back</a>`))
		return
	}
	rport, err := strconv.Atoi(rs[1])
	if len(rs) != 2 {
		rw.Write(message("Starting...", err.Error()+` <a href="/setup.html">go back</a>`))
		return
	}

	ls := strings.Split(r.PostFormValue("LocalServer"), ":")
	if len(ls) != 2 {
		rw.Write(message("Starting...", `Local server should be "host:port" <a href="/setup.html">go back</a>`))
		return
	}
	lport, err := strconv.Atoi(ls[1])
	if len(rs) != 2 {
		rw.Write(message("Starting...", err.Error()+`<a href="/setup.html">go back</a>`))
		return
	}
	bindPort, err := strconv.Atoi(r.PostFormValue("BindPort"))
	if err != nil {
		rw.Write(message("Starting...", err.Error()+`<a href="/setup.html">go back</a>`))
		return
	}
	client = bore.NewBoreClient(bore.Config{
		RemoteServer: rs[0],
		RemotePort:   rport,
		LocalServer:  ls[0],
		LocalPort:    lport,
		BindPort:     bindPort,
		ID:           r.PostFormValue("ID"),
		KeepAlive:    r.PostFormValue("KeepAlive") == "true",
	})
	rw.Write(message("Starting...", `Please wait, your NAT limitation is about to be gone... <meta http-equiv="Refresh" content="0; url='/info'" />`))
	go runclient(r.PostFormValue("AutoReconnect"))
}

func info(rw http.ResponseWriter, r *http.Request) {
	if len(strings.Split(html.EscapeString(client.RemoteData.DirectTCP), ":")) < 2 {
		rw.Write(message("Info", `We are still loading.. please wait. <meta http-equiv="Refresh" content="1; url='/info'" />`))
		return
	}
	var tcp string
	if len(strings.Split(html.EscapeString(client.RemoteData.DirectTCP), ":")) == 4 {
		tcp = strings.Join(strings.Split(html.EscapeString(client.RemoteData.DirectTCP), ":")[0:2], ":") + ":" + strings.Split(html.EscapeString(client.RemoteData.DirectTCP), ":")[3]
	} else {
		tcp = html.EscapeString(client.RemoteData.DirectTCP)
	}
	rw.Write(message("Info", `
<table>
	<tr>
		<th>Local</th>
		<td>`+html.EscapeString(client.LocalEndpoint.String())+`</td>
	</tr>
	<tr>
		<th>Server</th>
		<td>`+html.EscapeString(client.ServerEndpoint.String())+`</td>
	</tr>
	<tr>
		<th>TCP</th>
		<td>`+tcp+`</td>
	</tr>
	<tr>
		<th>HTTP</th>
		<td>`+strings.Join(strings.Split(html.EscapeString(client.RemoteData.HTTPurl), ":")[0:2], ":")+`</td>
	</tr>
	<tr>
		<th>HTTPS</th>
		<td>`+strings.Join(strings.Split(html.EscapeString(client.RemoteData.HTTPSurl), ":")[0:2], ":")+`</td>
	</tr>
</table>
<b>`+serr+`</b>
<meta http-equiv="Refresh" content="5; url='/info'" />`))

}

func runclient(reconnect string) {
connect:
	if err := client.Run(); err != nil {
		if reconnect != "true" {
			return
		}
		serr = time.Now().String() + ": connection failed due: " + err.Error() + " reconnecting in 5s..."
		log.Println(serr)
		time.Sleep(time.Second * 5)
		goto connect
	}
}
