package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{}

	start := time.Now()
	addr := fmt.Sprintf("http://%s/messages/count", os.Getenv("SERVER_ADDRESS"))
	res, err := client.Get(addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	duration := time.Since(start)
	fmt.Printf("Duration: %v, Status code: %d", duration, res.StatusCode)
}
