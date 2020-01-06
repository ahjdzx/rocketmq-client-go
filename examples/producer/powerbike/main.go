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
	"os"

	"github.com/apache/rocketmq-client-go"
	"github.com/apache/rocketmq-client-go/primitive"
	"github.com/apache/rocketmq-client-go/producer"
	"github.com/apache/rocketmq-client-go/rlog"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"10.111.75.96:9876", "10.111.75.95:9876"}),
		producer.WithRetry(2),
	)
	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to new producer: %s", err), nil)
		return
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	// tags := []string{"TagA", "TagB", "TagC"}
	for i := 0; ; i++ {
		// tag := tags[i%3]

		var tag string
		var msg *primitive.Message
		if i%2 == 0 {
			tag = "powerbike_gw_lock.mock"
			msg = primitive.NewMessage("ev_tcp_lock_message_exchange", []byte(`{"commandCode":1,"id":"123"}`))
			msg.WithTag(tag)
		} else {
			tag = "powerbike_gw_common.mock"
			msg = primitive.NewMessage("ev_tcp_common_message_exchange", []byte(`{"commandCode":7,"id":"123"}`))
			msg.WithTag(tag)
		}

		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}

	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
