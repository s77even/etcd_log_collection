package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)
type LogData struct {
	Topic string `json :"topic"`
	Data string `json :“data“`
}
var (
	client *elastic.Client
	ch chan *LogData
)

func Init(addr string , chan_size , goroutine_nums int)(err error){
	if !strings.HasPrefix(addr,"http://"){
		addr= "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		// Handle error
		fmt.Println("creat es new client failed , ", err)
		return
	}
	//init chan
	ch  = make(chan *LogData , chan_size)
	for i:=0; i<goroutine_nums; i++{
		go sendToEs()
	}

	return
}

func SendToEsChan(msg *LogData){
	ch <- msg
}

func sendToEs(){
	for{
		select {
		case msg := <- ch:
			put1 , err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println( put1.Index , put1.Type)
		default:
			time.Sleep(100*time.Millisecond)
		}
	}
}
