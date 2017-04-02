package main

import "net"
import "fmt"
import "bufio"
import "strconv"


func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	i := 0

	for {
		text := strconv.Itoa(i)
		i++
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}