package conf

//日志配置文件 映射结构体
type AppConfig struct {
	KafkaConf`ini:"kafka"`
	EtcdConf`ini:"etcd"`
}
// 日志文件 不同结的内容
// kafka config
type KafkaConf struct {
	Address string	`ini:"address"`
	ChanMaxsize int `ini:"chan_max_size"`
}

// etcd config
type EtcdConf struct {
	Address string	`ini:"address"`
	Key string `ini:"collect_log_key"`
	Timeout int `ini:"timeout"`
}


//TailLogConf unused
type TailLogConf struct {
	Filename string `ini:"filename"`
}