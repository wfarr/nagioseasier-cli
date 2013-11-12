package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/var/lib/nagios/rw/nagios.qh")

	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	go func(r io.Reader) {
		buf := make([]byte, 4096)

		for {
			n, err := r.Read(buf[:])
			if err != nil {
				return
			}

			fmt.Println("got:", string(buf[0:n]))
		}
	} (conn)

	_, err = conn.Write([]byte("#help"))

	if err != nil {
		panic(err.Error())
	}
}