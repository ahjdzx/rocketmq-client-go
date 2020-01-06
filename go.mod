module github.com/apache/rocketmq-client-go

require (
	github.com/emirpasic/gods v1.12.0
	github.com/golang/mock v1.3.1
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.1
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/gjson v1.2.1
	github.com/tidwall/match v1.0.1 // indirect
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65 // indirect
	go.uber.org/atomic v1.5.1
	stathat.com/c/consistent v1.0.0
	upspin.io v0.0.0-20191202193759-e5f200a5136b
)

replace stathat.com/c/consistent v1.0.0 => github.com/stathat/consistent v1.0.0

go 1.13
