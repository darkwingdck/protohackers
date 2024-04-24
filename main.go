package main

import (
    "fmt"
    "os"
    "protohackers/smoketest"
    "protohackers/primetime"
)

func main()  {
    if len(os.Args) != 2 {
	fmt.Println("Usage: ./main [task number]")
	return
    }
    taskName := os.Args[1]
    switch taskName {
    case "1":
	smoketest.Run()
    case "2":
	primetime.Run()
    default:
	fmt.Println("Task not found")
    }
}
