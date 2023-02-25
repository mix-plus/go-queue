package main

import (
	"fmt"
	"github.com/mix-plus/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var c kq.KafkaConf
	conf.MustLoad("config.yaml", &c)

	q := kq.MustNewQueue(c, kq.WithHandle(func(key, value string) error {
		fmt.Printf("%s => %s\n", key, value)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
