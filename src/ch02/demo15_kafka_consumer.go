package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
)

var (
	wg1 sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("192.168.10.28:9092,192.168.10.29:9092,192.168.10.31:9092", ","), nil)
	if err != nil {
		panic(err)
	}
	partitionList, err := consumer.Partitions("hello")
	if err != nil {
		panic(err)
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("hello", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg1.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg1.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg1.Wait()
		consumer.Close()
	}
}
