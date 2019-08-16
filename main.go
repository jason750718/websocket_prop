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
		fmt.Printf("New client connected\n")
		return
	}

	serv.OnMessage = func(fd int64, data []byte) {
		fmt.Printf("%d_Client send:%s\n", fd, string(data))
		serv.Send(fd, data)
	}

	serv.OnHeartBeat = func(fd int64) {
		s := fmt.Sprintf("%d_Client heartbeat", fd)
		fmt.Print(s + "\n")
		serv.Send(fd, []byte(s))
	}

	serv.OnTimeout = func(fd int64) {
		s := fmt.Sprintf("%d_Client timeout", fd)
		fmt.Print(s + "\n")
		serv.Boardcast(fd, []byte(s))
	}

	serv.OnClose = func(fd int64) {
		//deal close
	}

	serv.Run()
}
