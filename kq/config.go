package kq

import (
	"github.com/zeromicro/go-zero/core/service"
)

const (
	firstOffset = "first"
	lastOffset  = "last"
)

type KafkaConf struct {
	service.ServiceConf
	Brokers     []string
	Group       string
	Topic       string
	Offset      string `json:",options=first|last,default=last"`
	Conns       int    `json:",default=1"`
	Consumers   int    `json:",default=8"`
	Processors  int    `json:",default=8"`
	MinBytes    int    `json:",default=10240"`    // 10K
	MaxBytes    int    `json:",default=10485760"` // 10M
	Username    string `json:",optional"`
	Password    string `json:",optional"`
	ForceCommit bool   `json:",default=true"`
	// Maximum amount of time to wait for new data to come when fetching batches
	MaxWait int `json:",default=10"`
	// ReadBatchTimeout amount of time to wait to fetch message from kafka messages batch.
	ReadBatchTimeout int `json:",default=10"`
	// Limit of how many attempts will be made before delivering the error.
	MaxAttempts int `json:",default=3"`
}
