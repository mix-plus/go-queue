package main

import (
	"encoding/json"
	"fmt"
	"github.com/mix-plus/go-queue/internal/bs"
	"github.com/mix-plus/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"payload"`
}

func main() {
	pusher := kq.NewPusher([]string{
		"127.0.0.1:9092",
		"127.0.0.1:9092",
		"127.0.0.1:9092",
	}, "kq")

	ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 3; round++ {
		<-ticker.C
		count := rand.Intn(100)
		m := Message{
			Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
			Value:   fmt.Sprintf("%d,%d", round, count),
			Payload: fmt.Sprintf("%d,%d", round, count),
		}
		body, err := json.Marshal(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(bs.BytesToString(body))
		if err := pusher.Push(bs.BytesToString(body)); err != nil {
			log.Fatal(err)
		}
	}

	cmdline.EnterToContinue()
}
