package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"oldboystudy.com/logagent/kafka"
)

//从日志文件中搜集日志

type TailTask struct {
	path string
	topic string
	instance *tail.Tail
	ctx context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path , topic string) (tailObj *TailTask){
	ctx , cancle := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path : path,
		topic: topic,
		ctx : ctx,
		cancelFunc: cancle,
	}
	tailObj.init()
	return
}

func (t *TailTask)init(){
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //开始位置
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance , err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed,", err)
	}
	go t.run()
}

func (t *TailTask)run(){
	for {
		select {
		case <- t.ctx.Done():
			fmt.Printf("TailTask %s _ %s done..." ,t.path , t.topic )
			return
		case line :=  <- t.instance.Lines:
			kafka.SendToChan(t.topic, line.Text )
			//fmt.Println("SendToChan success")
		}
	}
}



func (t *TailTask)ReadChan() <-chan *tail.Line{
	return t.instance.Lines
}
