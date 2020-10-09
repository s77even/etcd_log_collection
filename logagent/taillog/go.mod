module oldboystudy.com/logagent/taillog

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/hpcloud/tail v1.0.0
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	oldboystudy.com/logagent/kafka v0.0.0
	oldboystudy.com/logagent/etcd v0.0.0
)

replace oldboystudy.com/logagent/kafka => ../kafka
replace oldboystudy.com/logagent/etcd => ../etcd