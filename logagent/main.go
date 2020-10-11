package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"oldboystudy.com/logagent/conf"
	"oldboystudy.com/logagent/etcd"
	"oldboystudy.com/logagent/kafka"
	"oldboystudy.com/logagent/taillog"
	"sync"
	"time"
)

// 构造一个结构体指针
var cfg = new(conf.AppConfig)



func main(){
	//加载配置文件conf
	err := ini.MapTo(cfg,"./conf/config.ini")
	if err != nil {
		fmt.Println("load ini file failed,",err)
		return
	}
	//初始化kafka连接
	err =kafka.Init([]string{cfg.KafkaConf.Address},cfg.KafkaConf.ChanMaxsize)
	if err!=nil{
		fmt.Println("Init kafka failed,err:",err)
		return
	}
	fmt.Println("inti kafka success")
	//初始化etcd连接
	err = etcd.Init(cfg.EtcdConf.Address,time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("Init etcd failed , err :" ,err)
		return
	}
	fmt.Println("inti etcd success")
	//从etcd中获取到配置信息
	logEntryConf ,err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Println("Getconf from etcd failed , err :" ,err)
		return
	}
	fmt.Println("Getconf from etcd success",logEntryConf )

	// 循环每一个日志收集箱 创建一个tailobj
	taillog.Init(logEntryConf)

	var wg sync.WaitGroup
	newConfChan := taillog.NewConfChan()
	//派一个哨兵监控日志配置文件的变化
	wg.Add(1)
	go etcd.ConfWatch(cfg.EtcdConf.Key,newConfChan)
	wg.Wait()
}