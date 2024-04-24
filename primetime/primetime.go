package primetime

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"strings"
)

type Request struct {
    Method string
    Number float64
}

func isPrime(number float64) bool {
    numberInt := int(number)
    if number != float64(numberInt) || numberInt < 2 {
	return false
    }
    root := int(math.Sqrt(number))
    for i := 2; i <= root; i++ {
	if numberInt % i == 0 {
	    return false
	}
    }
    return true
}

func getResult(request Request) string {
    if request.Method == "" {
	return "{}\n"
    }
    return fmt.Sprintf("{\"method\": \"isPrime\", \"prime\": %t}\n", isPrime(request.Number))
}

func serializeData(data []byte) Request {
    var m map[string]interface{}
    defaultRequest := Request{"", 0}
    
    err := json.Unmarshal(data, &m)

    if err != nil {
	return defaultRequest
    }

    method, ok := m["method"]
    
    if !ok || method != "isPrime" {
	return defaultRequest
    }

    number := m["number"]
    numberType := fmt.Sprintf("%T", number)
    
    if numberType != "float64" {
	return defaultRequest
    }

    return Request{method.(string), number.(float64)}
}

func handleConnection(connection net.Conn) {
    defer connection.Close()

    reader := bufio.NewReader(connection)

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                fmt.Println("Error reading: ", err)
            }
            break
        }
        line = strings.TrimSuffix(line, "\n")

        fmt.Printf("Received %d bytes: %s\n", len(line), line)

        request := serializeData([]byte(line))

        result := getResult(request)
        
        fmt.Println(result)

        _, err = connection.Write([]byte(result))
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
