package smoketest

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 4096)

	for {
		length, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading: ", err)
			return
		}

		fmt.Printf("Recieved: %s\n", buffer[:length])

		_, err = connection.Write(buffer[:length])
		if err != nil {
			fmt.Println("Error writing: ", err)
		}
	}
}

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	listener, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		fmt.Println("Error listening: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port " + port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			continue
		}

		go handleConnection(connection)
	}
}
