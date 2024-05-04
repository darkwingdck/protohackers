package meanstoanend

import (
    "bufio"
    "fmt"
    "io"
    "net"
    "os"
)

type Request struct {
    Method string  `json:"method"`
    Number *float64 `json:"number"`
}

type Response struct {
    Method string `json:"method"`
    Prime  bool   `json:"prime"` 
}

func handleConnection(connection net.Conn) {
    defer connection.Close()

    reader := bufio.NewReader(connection)

    for {
	line, err := reader.ReadString('\n')
	if line == "\n" {
	    continue
	}
	if err != nil {
	    if err != io.EOF {
		fmt.Println("Error reading: ", err)
	    }
	    break
	}
	// line = strings.TrimSuffix(line, "\n")

	fmt.Printf("Received %d bytes: %s\n", len(line), line)

	fmt.Println("Result: ", fmt.Sprint(101))
	_, err = connection.Write([]byte("101"))
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
    listener, err := net.Listen("tcp", "0.0.0.0:"+port)

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
