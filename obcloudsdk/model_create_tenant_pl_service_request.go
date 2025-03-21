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

// checks if the CreateTenantPlServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTenantPlServiceRequest{}

// CreateTenantPlServiceRequest struct for CreateTenantPlServiceRequest
type CreateTenantPlServiceRequest struct {
	UserAccount *string `json:"userAccount,omitempty"`
}

// NewCreateTenantPlServiceRequest instantiates a new CreateTenantPlServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTenantPlServiceRequest() *CreateTenantPlServiceRequest {
	this := CreateTenantPlServiceRequest{}
	return &this
}

// NewCreateTenantPlServiceRequestWithDefaults instantiates a new CreateTenantPlServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTenantPlServiceRequestWithDefaults() *CreateTenantPlServiceRequest {
	this := CreateTenantPlServiceRequest{}
	return &this
}

// GetUserAccount returns the UserAccount field value if set, zero value otherwise.
func (o *CreateTenantPlServiceRequest) GetUserAccount() string {
	if o == nil || IsNil(o.UserAccount) {
		var ret string
		return ret
	}
	return *o.UserAccount
}

// GetUserAccountOk returns a tuple with the UserAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantPlServiceRequest) GetUserAccountOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccount) {
		return nil, false
	}
	return o.UserAccount, true
}

// HasUserAccount returns a boolean if a field has been set.
func (o *CreateTenantPlServiceRequest) HasUserAccount() bool {
	if o != nil && !IsNil(o.UserAccount) {
		return true
	}

	return false
}

// SetUserAccount gets a reference to the given string and assigns it to the UserAccount field.
func (o *CreateTenantPlServiceRequest) SetUserAccount(v string) {
	o.UserAccount = &v
}

func (o CreateTenantPlServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTenantPlServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAccount) {
		toSerialize["userAccount"] = o.UserAccount
	}
	return toSerialize, nil
}

type NullableCreateTenantPlServiceRequest struct {
	value *CreateTenantPlServiceRequest
	isSet bool
}

func (v NullableCreateTenantPlServiceRequest) Get() *CreateTenantPlServiceRequest {
	return v.value
}

func (v *NullableCreateTenantPlServiceRequest) Set(val *CreateTenantPlServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTenantPlServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTenantPlServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTenantPlServiceRequest(val *CreateTenantPlServiceRequest) *NullableCreateTenantPlServiceRequest {
	return &NullableCreateTenantPlServiceRequest{value: val, isSet: true}
}

func (v NullableCreateTenantPlServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTenantPlServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


