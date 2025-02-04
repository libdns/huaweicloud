// Copyright 2022 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package invoker

import (
	"time"

	"github.com/libdns/huaweicloud/sdk/core"
	"github.com/libdns/huaweicloud/sdk/core/auth"
	"github.com/libdns/huaweicloud/sdk/core/def"
	"github.com/libdns/huaweicloud/sdk/core/exchange"
	"github.com/libdns/huaweicloud/sdk/core/invoker/retry"
)

type RetryChecker func(interface{}, error) bool

type BaseInvoker struct {
	Exchange *exchange.SdkExchange

	client  *core.HcHttpClient
	request interface{}
	meta    *def.HttpRequestDef
	headers map[string]string

	retryTimes      int
	retryChecker    RetryChecker
	backoffStrategy retry.Strategy
}

func NewBaseInvoker(client *core.HcHttpClient, request interface{}, meta *def.HttpRequestDef) *BaseInvoker {
	exch := &exchange.SdkExchange{
		ApiReference: &exchange.ApiReference{
			Method: meta.Method,
			Path:   meta.Path,
		},
		Attributes: make(map[string]interface{}),
	}

	return &BaseInvoker{
		Exchange: exch,
		client:   client,
		request:  request,
		meta:     meta,
		headers:  make(map[string]string),
	}
}

func (b *BaseInvoker) ReplaceCredentialWhen(fun func(auth.ICredential) auth.ICredential) *BaseInvoker {
	b.client.WithCredential(fun(b.client.GetCredential()))
	return b
}

func (b *BaseInvoker) AddHeaders(headers map[string]string) *BaseInvoker {
	if headers != nil {
		for k, v := range headers {
			b.headers[k] = v
		}
	}
	return b
}

func (b *BaseInvoker) WithRetry(retryTimes int, checker RetryChecker, backoffStrategy retry.Strategy) *BaseInvoker {
	b.retryTimes = retryTimes
	b.retryChecker = checker
	b.backoffStrategy = backoffStrategy
	return b
}

func (b *BaseInvoker) Invoke() (interface{}, error) {
	if b.retryTimes == 0 || b.retryChecker == nil {
		return b.client.SyncInvokeWithExtraHeaders(b.request, b.meta, b.Exchange, b.headers)
	}

	var execTimes int
	var resp interface{}
	var err error
	for {
		resp, err = b.client.SyncInvokeWithExtraHeaders(b.request, b.meta, b.Exchange, b.headers)
		execTimes++

		if execTimes > b.retryTimes || !b.retryChecker(resp, err) {
			break
		}

		delay := b.backoffStrategy.ComputeDelayBeforeNextRetry(int32(execTimes))
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
	return resp, err
}
