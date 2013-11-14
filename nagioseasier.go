package main

import (
	"fmt"
	"os"
	"net"
	"sort"
	"strings"

	"github.com/stevedomin/termtable"
)

func main() {
	// establish connection to our socket, for both reads and writes
	conn, err := net.Dial("unix", "/var/lib/nagios/rw/nagios.qh")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// suss out what command we actually wish to run
	if len(os.Args) == 0 {
		send_command(conn, "help")
	}

	switch os.Args[1] {
	case "help", "status", "acknowledge", "unacknowledge", "disable_notifications", "enable_notifications", "downtime", "problems":
		send_command(conn, strings.Join(os.Args[1:], " "))
	case "ack":
		send_command(conn, "acknowledge" + strings.Join(os.Args[2:], " "))
	case "unack":
		send_command(conn, "unacknowledge" + strings.Join(os.Args[2:], " "))
	case "mute":
		send_command(conn, "disable_notifications" + strings.Join(os.Args[2:], " "))
	case "unmute":
		send_command(conn, "enable_notifications" + strings.Join(os.Args[2:], " "))
	default:
		send_command(conn, "help")
	}

	output := read_results(conn)

	lines := strings.Split(output, "\n")
	sort.Sort(sort.StringSlice(lines))

	if len(lines[1:]) > 0 {
		t := termtable.NewTable(nil, &termtable.TableOptions{Padding: 5, UseSeparator: true})
		t.SetHeader([]string{"Service", "Status", "Details"})

		for _, line := range lines[1:] {
			parts := strings.Split(line, ";")
			var truncated string

			if len(parts[2]) > 50 {
				truncated = parts[2][0:50]
			} else {
				truncated = parts[2]
			}

			t.AddRow([]string{parts[0], parts[1], truncated})
		}

		fmt.Println(t.Render())
	} else {
		fmt.Println(lines[1:])
	}
}

func send_command(conn net.Conn, cmd string) {
	_, err := conn.Write([]byte(fmt.Sprintf("#nagioseasier %s\000", cmd)))
	if err != nil {
		panic(err.Error())
	}

}

func read_results(conn net.Conn) (output string) {
	buf := make([]byte, 4096)
	for ;; {
		n, err := conn.Read(buf[:])

		if err != nil {
			return output
		}

		if n == 0 {
			fmt.Println("Connection closed by socket")
			return output
		}

		output = output + string(buf[0:n])
	}
}