package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

//向kafka中写入日志
//logData 中的数据 topic表示要写入的标题 data是日志数据
type logData struct {
	topic string
	data string
}
//声明全局变量 cilent 和 日志文件写入的通道
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

// SendToChan 将topic data封装为结构体 传入通道
func SendToChan(topic, data string){
	msg := &logData{
		topic: topic,
		data: data,
	}
	logDataChan <- msg
}

//SendToKafka 从logDataChan中取出数据 将数据发往kafka
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