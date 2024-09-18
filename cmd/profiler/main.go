package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}

	start := time.Now()
	res, err := client.Get("http://localhost:8000/messages/count")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	duration := time.Since(start)
	fmt.Printf("Duration: %v, Status code: %d", duration, res.StatusCode)
}
