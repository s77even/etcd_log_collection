module oldboystudy.com/log_transfer

go 1.15

require (
	gopkg.in/ini.v1 v1.61.0
	oldboystudy.com/log_transfer/es v0.0.0-00010101000000-000000000000
	oldboystudy.com/log_transfer/kafka v0.0.0-00010101000000-000000000000
)

replace oldboystudy.com/log_transfer/kafka => ./kafka

replace oldboystudy.com/log_transfer/es => ./es
