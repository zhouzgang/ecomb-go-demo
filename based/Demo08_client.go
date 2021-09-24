package main

import (
	"fmt"
	"net"
	"time"
)

const RECV_BUF_LEN_CLIENT = 1024

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN_CLIENT)
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write buffer error:", err.Error())
			break
		}
		fmt.Println(msg)

		n, err = conn.Read(buf)
		if err != nil {
			println("Read buffer error:", err.Error())
			break
		}
		fmt.Println("return: ", string(buf[0:n]))

		time.Sleep(time.Second)
	}
}