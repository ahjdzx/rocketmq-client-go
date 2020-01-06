/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/rocketmq-client-go"
	"github.com/apache/rocketmq-client-go/consumer"
	"github.com/apache/rocketmq-client-go/primitive"
	"github.com/apache/rocketmq-client-go/rlog"
)

func main() {
	// hostname, _ := os.Hostname()
	consumers := []rocketmq.PushConsumer{}

	//	for i := 0; i < 5; i++ {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("easybike_gateway_mock_consumer"),
		consumer.WithNameServer([]string{"10.111.75.96:9876", "10.111.75.95:9876"}),
		// consumer.WithInstance(fmt.Sprintf("%s_%d", hostname, i)),
	)
	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to new pullConsumer: %s", err), nil)
		return
	}

	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "TagA || TagC",
	}

	selector = consumer.MessageSelector{}

	err = c.Subscribe("tcp_new_message_exchange", selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := c.Start(); err != nil {
		log.Println(err)
		return
	}

	consumers = append(consumers, c)
	//	}

	time.Sleep(time.Second * 10)

	for _, c := range consumers {
		if err := c.Shutdown(); err != nil {
			log.Println(err)
		}
	}
}
