package main

import (
	"fmt"
	"net/http"

	"websocket"
)

func main() {
	serv := websocket.NewServer(9008, "/", 5, 1, false, "", "")

	serv.OnOpen = func(r *http.Request) (isOpen bool, code int) {
		isOpen = true
		code = 101
		return
	}

	serv.OnMessage = func(fd int64, data []byte) {
		fmt.Printf("%d_Cliend send:%s\n", fd, string(data))
		serv.Send(fd, data)
	}

	serv.OnHeartBeat = func(fd int64) {
		fmt.Printf("%d_Cliend heartbeat\n", fd)
		serv.Send(fd, []byte("heartbeat"))
	}

	serv.OnClose = func(fd int64) {
		//deal close
	}

	serv.Run()
}
