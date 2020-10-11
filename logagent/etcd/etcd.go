package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path string `json:"path"` // 日志存放的路径
	Topic string `json:"topic"` // 日子要发往的topic
}

func Init(addr string, timeout time.Duration)(err error){
	cli , err = clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Println("connect to etcd failed ," ,err)
		return
	}
	return
}

//GetConf 根据key获取value 获取配置文件项
func GetConf(key string)(LogEntryValue []*LogEntry , err error){
	//get
	ctx , cancle := context.WithTimeout(context.Background(),time.Second)
	getResp, err := cli.Get(ctx, key)
	cancle()
	if err != nil {
		fmt.Println("get failed, ", err)
		return
	}
	for _ ,ev := range  getResp.Kvs {
		err = json.Unmarshal(ev.Value, &LogEntryValue)
		if err != nil {
			fmt.Println("unmarshal etcd key-value failed , ", err)
			return
		}
	}
	return
}

//ConfWatch 利用etcd的watch机制 监控配置文件的变化
func ConfWatch(key string, newConfCh chan<- []*LogEntry){
	ch := cli.Watch(context.Background(),key)
	for wresp := range ch{
		for _ , evt := range  wresp.Events{
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete{ // 如果是删除操作 在此什么都不做 通道传入空的newconf
				err :=json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("json unmarshal failed , err :", err)
					continue
				}
			}
			newConfCh <- newConf

		}
	}
}
