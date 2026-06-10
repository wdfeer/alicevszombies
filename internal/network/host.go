package network

import (
	"fmt"
	"net"
)

const hostPort uint16 = 7777

var portStr = fmt.Sprintf(":%d", hostPort)

func HostGame() {
	// TODO: queue changes to be made to the game state on the main thread
	go tcpServer()
	go udpServer()
}

func tcpServer() {
	ln, err := net.Listen("tcp", portStr)

	if err != nil {
		println("ERROR: failed to start the TCP listener!")
		println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			println("ERROR: failed to accept a TCP connection!")
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					println("ERROR: failed to read from TCP buffer!")
					continue
				}

				// TODO: figure out what to do with the messages
				c.Write(buf[:n])
			}
		}(conn)
	}
}

func udpServer() {
	addr, err := net.ResolveUDPAddr("udp", portStr)
	if err != nil {
		println("ERROR: failed to resolve UDP address!")
		println(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		println("ERROR: failed to start UDP listener!")
		println(err)
		return
	}

	buf := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			println("ERROR: failed to read from UDP buffer!")
			continue
		}

		// TODO: figure out what to do with the messages
		conn.WriteToUDP(buf[:n], clientAddr)
	}
}
