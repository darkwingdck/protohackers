package main

import (
    "fmt"
    "os"
    "protohackers/smoketest"
    "protohackers/primetime"
    "protohackers/meanstoanend"
)

func main()  {
    if len(os.Args) != 2 {
	fmt.Println("Usage: ./main [task number]")
	return
    }
    taskNumber := os.Args[1]
    switch taskNumber{
    case "0":
	smoketest.Run()
    case "1":
	primetime.Run()
    case "2":
	meanstoanend.Run()
    default:
	fmt.Println("Task not found")
    }
}
