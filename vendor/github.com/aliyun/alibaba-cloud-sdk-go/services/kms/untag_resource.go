package kms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UntagResource invokes the kms.UntagResource API synchronously
// api document: https://help.aliyun.com/api/kms/untagresource.html
func (client *Client) UntagResource(request *UntagResourceRequest) (response *UntagResourceResponse, err error) {
	response = CreateUntagResourceResponse()
	err = client.DoAction(request, response)
	return
}

// UntagResourceWithChan invokes the kms.UntagResource API asynchronously
// api document: https://help.aliyun.com/api/kms/untagresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourceWithChan(request *UntagResourceRequest) (<-chan *UntagResourceResponse, <-chan error) {
	responseChan := make(chan *UntagResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UntagResource(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UntagResourceWithCallback invokes the kms.UntagResource API asynchronously
// api document: https://help.aliyun.com/api/kms/untagresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourceWithCallback(request *UntagResourceRequest, callback func(response *UntagResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UntagResourceResponse
		var err error
		defer close(result)
		response, err = client.UntagResource(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UntagResourceRequest is the request struct for api UntagResource
type UntagResourceRequest struct {
	*requests.RpcRequest
	TagKeys string `position:"Query" name:"TagKeys"`
	KeyId   string `position:"Query" name:"KeyId"`
}

// UntagResourceResponse is the response struct for api UntagResource
type UntagResourceResponse struct {
	*responses.BaseResponse
	KeyId string `json:"KeyId" xml:"KeyId"`
}

// CreateUntagResourceRequest creates a request to invoke UntagResource API
func CreateUntagResourceRequest() (request *UntagResourceRequest) {
	request = &UntagResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "UntagResource", "kms", "openAPI")
	return
}

// CreateUntagResourceResponse creates a response to parse from UntagResource response
func CreateUntagResourceResponse() (response *UntagResourceResponse) {
	response = &UntagResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
