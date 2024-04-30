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
    Method string  `json:"method"`
    Number *float64 `json:"number"`
}

type Response struct {
    Method string `json:"method"`
    Prime  bool   `json:"prime"` 
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

func getResult(request *Request) ([]byte, error) {
    defaultResult := []byte("{}\n")

    isNumberNil := request.Number == nil
    if isNumberNil {
	return defaultResult, fmt.Errorf("Error: number is nil")
    }

    isMethodValid := request.Method == "isPrime"
    isNumberValid := fmt.Sprintf("%T", *request.Number) == "float64"


    if !isMethodValid || !isNumberValid {
	return defaultResult, fmt.Errorf("Error: request not valid")
    }

    response := Response{
	Method: request.Method,
	Prime: isPrime(*request.Number),
    }

    result, _ := json.Marshal(response)
    result = append(result, '\n')
    return result, nil
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
	line = strings.TrimSuffix(line, "\n")

	fmt.Printf("Received %d bytes: %s\n", len(line), line)

	request := Request{}
	err = json.Unmarshal([]byte(line), &request)
	if err != nil {
	    fmt.Println("Error Unmarshal: ", err)
	    request = Request{Method: "Error", Number: nil}
	}

	result, _ := getResult(&request)
	fmt.Println("Result: ", string(result))
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
