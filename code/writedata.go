package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: writedata /path/to/file")
		panic("You're missing the file name")
	}
	filename := os.Args[1]
	hostname, err := os.Hostname()

	if err != nil {
		panic("Cannot determine the hostname.")
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	for i := 0; i < 50; i++ {
		data := fmt.Sprintf("Host: %s - Loop Number: %d\n", hostname, i)

		_, err := f.WriteString(data)
		if err != nil {
			panic("Cannot write to file")
		}
		time.Sleep(time.Second)
	}
}

// Compile with:
// env GOARCH=386 GOOS=linux go build writedata.go
