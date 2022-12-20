package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	conn := newServer()

	go start(conn)

	select {}
}

func newServer() *net.UDPConn {
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 10001, Zone: ""})

	return conn
}

func start(conn *net.UDPConn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, _, _ := conn.ReadFromUDP(buf)
		var d map[string]interface{}

		json.Unmarshal(buf[0:n], &d)

		fmt.Println(d)
	}
}
