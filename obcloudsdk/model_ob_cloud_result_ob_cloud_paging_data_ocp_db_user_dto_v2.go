/*
OceanBase Cloud API

API Documentation for OceanBase Cloud

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package obcloudsdk

import (
	"encoding/json"
)

// checks if the OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OBCloudResultOBCloudPagingDataOcpDbUserDtoV2{}

// OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 struct for OBCloudResultOBCloudPagingDataOcpDbUserDtoV2
type OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 struct {
	Success *bool `json:"success,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	Data *OBCloudPagingDataOcpDbUserDtoV2 `json:"data,omitempty"`
	Cost *int64 `json:"cost,omitempty"`
	Server *string `json:"server,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
	Extra map[string]map[string]interface{} `json:"extra,omitempty"`
}

// NewOBCloudResultOBCloudPagingDataOcpDbUserDtoV2 instantiates a new OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOBCloudResultOBCloudPagingDataOcpDbUserDtoV2() *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 {
	this := OBCloudResultOBCloudPagingDataOcpDbUserDtoV2{}
	return &this
}

// NewOBCloudResultOBCloudPagingDataOcpDbUserDtoV2WithDefaults instantiates a new OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOBCloudResultOBCloudPagingDataOcpDbUserDtoV2WithDefaults() *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 {
	this := OBCloudResultOBCloudPagingDataOcpDbUserDtoV2{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetSuccess(v bool) {
	o.Success = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetData() OBCloudPagingDataOcpDbUserDtoV2 {
	if o == nil || IsNil(o.Data) {
		var ret OBCloudPagingDataOcpDbUserDtoV2
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetDataOk() (*OBCloudPagingDataOcpDbUserDtoV2, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OBCloudPagingDataOcpDbUserDtoV2 and assigns it to the Data field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetData(v OBCloudPagingDataOcpDbUserDtoV2) {
	o.Data = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetCost() int64 {
	if o == nil || IsNil(o.Cost) {
		var ret int64
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetCostOk() (*int64, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int64 and assigns it to the Cost field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetCost(v int64) {
	o.Cost = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetServer(v string) {
	o.Server = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetExtra() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Extra) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) GetExtraOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Extra) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]map[string]interface{} and assigns it to the Extra field.
func (o *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) SetExtra(v map[string]map[string]interface{}) {
	o.Extra = v
}

func (o OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	return toSerialize, nil
}

type NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2 struct {
	value *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2
	isSet bool
}

func (v NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2) Get() *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 {
	return v.value
}

func (v *NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2) Set(val *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) {
	v.value = val
	v.isSet = true
}

func (v NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2) IsSet() bool {
	return v.isSet
}

func (v *NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2(val *OBCloudResultOBCloudPagingDataOcpDbUserDtoV2) *NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2 {
	return &NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2{value: val, isSet: true}
}

func (v NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOBCloudResultOBCloudPagingDataOcpDbUserDtoV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


