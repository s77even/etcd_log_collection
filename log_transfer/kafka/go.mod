module oldboystudy.com/log_transfer/kafka

go 1.15

require (
	github.com/Shopify/sarama v1.27.1
	oldboystudy.com/log_transfer/es v0.0.0-00010101000000-000000000000
)

replace oldboystudy.com/log_transfer/es => ../es
