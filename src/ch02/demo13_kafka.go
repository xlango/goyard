package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	// kafka 服务器地址,以及端口号,这里可以指定多个地址,使用逗号分隔开即可.
	kafka  = "192.168.10.28:9092,192.168.10.29:9092,192.168.10.31:9092"
	wg     sync.WaitGroup
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {

	sarama.Logger = logger
	// 连接kafka消息服务器
	consumer, err := sarama.NewConsumer(strings.Split(kafka, ","), nil)
	if err != nil {
		logger.Println("Failed to start consumer: %s", err)
	}

	// consumer.Partitions 用户获取Topic上所有的Partitions. 消息服务器上已经创建了test这个topic,所以,在这里指定参数为test.
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		logger.Println("Failed to get the list of partitions: ", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Println("message is :", msg)
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}

	wg.Wait()

	logger.Println("Done consuming topic hello")
	consumer.Close()
}
