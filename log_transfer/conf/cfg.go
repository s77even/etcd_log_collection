//conf中存放了从配置文件中ini解析的结构体及数据类型
package conf

type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	EsCfg `ini:"es"`
}

type KafkaCfg struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type EsCfg struct {
	Address string `ini:"address"`
	Chan_size int `ini:"chan_size"'`
	Goroutine_nums int `ini:"goroutine_nums"`
}