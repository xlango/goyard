package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"strings"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewSyncProducer(strings.Split("192.168.10.28:9092,192.168.10.29:9092,192.168.10.31:9092", ","), config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic:     "hello",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}
	var value string
	for {
		// 生产消息
		inputReader := bufio.NewReader(os.Stdin)
		value, err = inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		value = strings.Replace(value, "\n", "", -1)
		msg.Value = sarama.ByteEncoder(value)
		paritition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Send Message Fail")
		}
		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
	}
}
