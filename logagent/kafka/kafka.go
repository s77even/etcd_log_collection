package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

//向kafka中写入日志

type logData struct {
	topic string
	data string
}
var (
	client sarama.SyncProducer
	logDataChan chan *logData
)

//Init 初始化连接
func Init(adds []string,maxsize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //waitforall 等待所有leader和follower全部返回ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 轮询模式决定partition
	config.Producer.Return.Successes = true                   //success channel 返回成功交付
	//connect
	client, err = sarama.NewSyncProducer(adds, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	//init logDataChan
	logDataChan = make(chan *logData, maxsize)
	go SendToKafka()
	return
}

func SendToChan(topic, data string){
	msg := &logData{
		topic: topic,
		data: data,
	}
	logDataChan <- msg
}


func SendToKafka() {
	for{
		select {
		case ld:= <- logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			//send
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send message failed,", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond*50)
		}
	}

}