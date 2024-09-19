package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {

	addr := url.URL{Scheme: "ws", Host: "localhost:8000", Path: "/ws/chat"}

	count, err := strconv.Atoi(os.Getenv("PROFILE_REQUEST_COUNT"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var wsConnPool = make([]*websocket.Conn, 0, 8)
	var wsConnMutexes = make([]*sync.Mutex, 0, 8)
	for range 8 {
		wsConn, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
		if err != nil {
			log.Fatal("Error connecting to WebSocket server:", err)
		}

		wsConnPool = append(wsConnPool, wsConn)
		wsConnMutexes = append(wsConnMutexes, &sync.Mutex{})
	}
	defer func() {
		for _, wsConn := range wsConnPool {
			wsConn.Close()
		}
	}()

	var mx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(count)

	successCount := 0
	start := time.Now()

	for i := 0; i < count; i++ {

		connIndex := i % len(wsConnPool)
		conn := wsConnPool[connIndex]
		connMutex := wsConnMutexes[connIndex]

		go func() {
			defer wg.Done()

			message := []byte("Benchmark message " + time.Now().String())
			connMutex.Lock()
			err = conn.WriteMessage(websocket.TextMessage, message)
			connMutex.Unlock()
			if err != nil {
				log.Println("Write error:", err)
				return
			}

			connMutex.Lock()
			_, _, err = conn.ReadMessage()
			connMutex.Unlock()
			if err != nil {
				log.Println("Read error:", err)
				return
			}

			mx.Lock()
			successCount++
			mx.Unlock()
		}()
	}

	wg.Wait()

	duration := time.Since(start)
	// time.Sleep(5 * time.Second)
	fmt.Printf("Duration: %v, Success count: %v, Total count: %v", duration, successCount, count)
}
