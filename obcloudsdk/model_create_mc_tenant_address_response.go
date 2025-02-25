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

// checks if the CreateMcTenantAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMcTenantAddressResponse{}

// CreateMcTenantAddressResponse struct for CreateMcTenantAddressResponse
type CreateMcTenantAddressResponse struct {
	AddressId *string `json:"addressId,omitempty"`
}

// NewCreateMcTenantAddressResponse instantiates a new CreateMcTenantAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMcTenantAddressResponse() *CreateMcTenantAddressResponse {
	this := CreateMcTenantAddressResponse{}
	return &this
}

// NewCreateMcTenantAddressResponseWithDefaults instantiates a new CreateMcTenantAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMcTenantAddressResponseWithDefaults() *CreateMcTenantAddressResponse {
	this := CreateMcTenantAddressResponse{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *CreateMcTenantAddressResponse) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMcTenantAddressResponse) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *CreateMcTenantAddressResponse) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *CreateMcTenantAddressResponse) SetAddressId(v string) {
	o.AddressId = &v
}

func (o CreateMcTenantAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMcTenantAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["addressId"] = o.AddressId
	}
	return toSerialize, nil
}

type NullableCreateMcTenantAddressResponse struct {
	value *CreateMcTenantAddressResponse
	isSet bool
}

func (v NullableCreateMcTenantAddressResponse) Get() *CreateMcTenantAddressResponse {
	return v.value
}

func (v *NullableCreateMcTenantAddressResponse) Set(val *CreateMcTenantAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMcTenantAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMcTenantAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMcTenantAddressResponse(val *CreateMcTenantAddressResponse) *NullableCreateMcTenantAddressResponse {
	return &NullableCreateMcTenantAddressResponse{value: val, isSet: true}
}

func (v NullableCreateMcTenantAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMcTenantAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


