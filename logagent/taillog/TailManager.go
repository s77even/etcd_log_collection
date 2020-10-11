package taillog

import (
	"fmt"
	"oldboystudy.com/logagent/etcd"
	"time"
)

// 日志收集任务的管理者
type Manager struct {
	logEntry []*etcd.LogEntry //所有日志收集项的配置 包括path和topic
	taskMap map[string]*TailTask// 日志收集任务的映射 方便查找到对应任务
	newConfChan chan []*etcd.LogEntry //监控配置改变的通道
}

var TskMgr *Manager
//Init ...
func Init(logEntryConf []*etcd.LogEntry){
	TskMgr = &Manager{
		logEntry: logEntryConf,
		taskMap: make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry) ,  //无缓冲区通道
	}
	//收集日志发往kafka
	for _ , logEntry := range logEntryConf {
		tailObj :=NewTailTask(logEntry.Path,logEntry.Topic)
		mkey := fmt.Sprintf("%s_%s",logEntry.Path,logEntry.Topic)
		TskMgr.taskMap[mkey]=tailObj
	}

	go TskMgr.run()
}

// run 对日志收集项的配置文件的改变做出对应操作
func (t *Manager)run(){
	for {
		select {
		case newConf := <- t.newConfChan:
			//新增
			for _,conf := range newConf{
				key := fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
				_ , ok := t.taskMap[key]
				if ok {
					continue
				}else{
					tailobj := NewTailTask(conf.Path,conf.Topic)
					t.taskMap[key]=tailobj
				}
				}
				//删除 在logEntry中 不在newConf中
			for _ , c1 := range t.logEntry{
				isDelete := true
				for _ , c2 := range newConf{
					if c2.Topic==c1.Topic && c2.Path==c1.Path {
						isDelete = false
						continue
					}
				}
				if isDelete {
					key := fmt.Sprintf("%s_%s",c1.Path,c1.Topic)
					t.taskMap[key].cancelFunc()
				}
			}


			fmt.Println("new config come..." ,newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 对外暴露newConfChan
func NewConfChan()  chan<- []*etcd.LogEntry {
	return  TskMgr.newConfChan
}
