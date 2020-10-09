package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"oldboystudy.com/log_transfer/conf"
	"oldboystudy.com/log_transfer/es"
	"oldboystudy.com/log_transfer/kafka"
)

func main(){
	//load config file (.ini)
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("ini config failed , ", err)
		return
	}
	//fmt.Println(cfg)

	//init es
	err =es.Init(cfg.EsCfg.Address,cfg.Chan_size,cfg.Goroutine_nums)
	if err != nil {
		fmt.Println("init er failed ," , err)
		return
	}
	fmt.Println("init es success...")
	// init kafka
	err =kafka.Init([]string{cfg.KafkaCfg.Address},cfg.Topic)
	if err != nil {
		fmt.Println("kafka init failed , " , err)
		return
	}
	select {}
}


