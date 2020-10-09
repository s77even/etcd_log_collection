package conf

type AppConfig struct {
	KafkaConf`ini:"kafka"`
	EtcdConf`ini:"etcd"`
}
type KafkaConf struct {
	Address string	`ini:"address"`
	ChanMaxsize int `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string	`ini:"address"`
	Key string `ini:"collect_log_key"`
	Timeout int `ini:"timeout"`
}

//TailLogConf unused
type TailLogConf struct {
	Filename string `ini:"filename"`
}