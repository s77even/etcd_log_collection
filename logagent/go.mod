module oldboystudy.com/logagent

go 1.15

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/google/uuid v1.1.2 // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.32.0 // indirect
	gopkg.in/ini.v1 v1.61.0
	oldboystudy.com/logagent/etcd v0.0.0
	oldboystudy.com/logagent/kafka v0.0.0
	oldboystudy.com/logagent/taillog v0.0.0
)

replace oldboystudy.com/logagent/kafka => ./kafka

replace oldboystudy.com/logagent/taillog => ./taillog

replace oldboystudy.com/logagent/etcd => ./etcd

replace google.golang.org/grpc v1.32.0 => google.golang.org/grpc v1.26.0


