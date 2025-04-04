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

package retry

import (
	"github.com/libdns/huaweicloud/sdk/core/utils"
)

// DecorRelatedJitter 去相关抖动退避 delay = min(最大等待时间, random(基础等待时间, 基础等待时间 * 3);
type DecorRelatedJitter struct {
	*strategyBase
}

func NewDecorRelatedJitter() *DecorRelatedJitter {
	return &DecorRelatedJitter{strategyBase: newStrategyBase()}
}

func (d *DecorRelatedJitter) ComputeDelayBeforeNextRetry(int32) int32 {
	return utils.Min32(MaxDelay, utils.RandInt32(d.GetBaseDelay(), d.GetBaseDelay()*3))
}
