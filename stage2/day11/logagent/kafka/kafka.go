package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

var kafkaClient sarama.SyncProducer

var logDataChan chan *logData

// Init 初始化kafka
func Init(addr []string) (err error) {
	config := sarama.NewConfig()
	// tailf包使⽤
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	kafkaClient, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	logDataChan = make(chan *logData, 100000)
	go sendDataToKafka()
	return
}

func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

func sendDataToKafka() {
	for {
		select {
		case logData := <-logDataChan:
			msg := &sarama.ProducerMessage{
				Topic: logData.topic,
				Value: sarama.StringEncoder(logData.data),
			}
			message, offset, err := kafkaClient.SendMessage(msg)
			if err != nil {
				panic(err)
			}
			fmt.Printf("pid:%v offset:%v\n", message, offset)
		default:
			time.Sleep(time.Second)
		}
	}
}
