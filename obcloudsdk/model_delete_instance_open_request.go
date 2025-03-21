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

// checks if the DeleteInstanceOpenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteInstanceOpenRequest{}

// DeleteInstanceOpenRequest struct for DeleteInstanceOpenRequest
type DeleteInstanceOpenRequest struct {
	RequestId *string `json:"requestId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
}

// NewDeleteInstanceOpenRequest instantiates a new DeleteInstanceOpenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInstanceOpenRequest() *DeleteInstanceOpenRequest {
	this := DeleteInstanceOpenRequest{}
	return &this
}

// NewDeleteInstanceOpenRequestWithDefaults instantiates a new DeleteInstanceOpenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInstanceOpenRequestWithDefaults() *DeleteInstanceOpenRequest {
	this := DeleteInstanceOpenRequest{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DeleteInstanceOpenRequest) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInstanceOpenRequest) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DeleteInstanceOpenRequest) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DeleteInstanceOpenRequest) SetRequestId(v string) {
	o.RequestId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *DeleteInstanceOpenRequest) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteInstanceOpenRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *DeleteInstanceOpenRequest) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *DeleteInstanceOpenRequest) SetInstanceId(v string) {
	o.InstanceId = &v
}

func (o DeleteInstanceOpenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteInstanceOpenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	return toSerialize, nil
}

type NullableDeleteInstanceOpenRequest struct {
	value *DeleteInstanceOpenRequest
	isSet bool
}

func (v NullableDeleteInstanceOpenRequest) Get() *DeleteInstanceOpenRequest {
	return v.value
}

func (v *NullableDeleteInstanceOpenRequest) Set(val *DeleteInstanceOpenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInstanceOpenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInstanceOpenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInstanceOpenRequest(val *DeleteInstanceOpenRequest) *NullableDeleteInstanceOpenRequest {
	return &NullableDeleteInstanceOpenRequest{value: val, isSet: true}
}

func (v NullableDeleteInstanceOpenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInstanceOpenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


