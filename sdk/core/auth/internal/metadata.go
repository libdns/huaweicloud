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

package internal

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/libdns/huaweicloud/sdk/core/sdkerr"
	"github.com/libdns/huaweicloud/sdk/core/utils"
)

type GetTemporaryCredentialFromMetadataResponse struct {
	Credential *Credential `json:"credential,omitempty"`
}

type Credential struct {
	ExpiresAt string `json:"expires_at"`

	Access string `json:"access"`

	Secret string `json:"secret"`

	Securitytoken string `json:"securitytoken"`
}

func GetCredentialFromMetadata() (*Credential, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get("http://169.254.169.254/openstack/latest/securitykey")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, &sdkerr.ServiceResponseError{
			StatusCode:   resp.StatusCode,
			ErrorMessage: string(body),
		}
	}

	respModel := &GetTemporaryCredentialFromMetadataResponse{}
	err = utils.Unmarshal(body, respModel)
	if err != nil {
		return nil, err
	}
	return respModel.Credential, nil
}
