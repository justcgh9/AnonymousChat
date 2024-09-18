package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	client := http.Client{}

	start := time.Now()
	addr := fmt.Sprintf("http://%s/messages/count", os.Getenv("SERVER_ADDRESS"))

	count, err := strconv.Atoi(os.Getenv("PROFILE_REQUEST_COUNT"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var mx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(count)

	successCount := 0

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			res, err := client.Get(addr)
			if err != nil {
				fmt.Println(err.Error())
			}

			if res.StatusCode == http.StatusOK {
				mx.Lock()
				successCount++
				mx.Unlock()
			}
		}()
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("Duration: %v, Success count: %v, Total count: %v", duration, successCount, count)
}
