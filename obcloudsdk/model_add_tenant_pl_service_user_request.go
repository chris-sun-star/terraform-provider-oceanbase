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

// checks if the AddTenantPlServiceUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddTenantPlServiceUserRequest{}

// AddTenantPlServiceUserRequest struct for AddTenantPlServiceUserRequest
type AddTenantPlServiceUserRequest struct {
	UserAccount *string `json:"userAccount,omitempty"`
}

// NewAddTenantPlServiceUserRequest instantiates a new AddTenantPlServiceUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTenantPlServiceUserRequest() *AddTenantPlServiceUserRequest {
	this := AddTenantPlServiceUserRequest{}
	return &this
}

// NewAddTenantPlServiceUserRequestWithDefaults instantiates a new AddTenantPlServiceUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTenantPlServiceUserRequestWithDefaults() *AddTenantPlServiceUserRequest {
	this := AddTenantPlServiceUserRequest{}
	return &this
}

// GetUserAccount returns the UserAccount field value if set, zero value otherwise.
func (o *AddTenantPlServiceUserRequest) GetUserAccount() string {
	if o == nil || IsNil(o.UserAccount) {
		var ret string
		return ret
	}
	return *o.UserAccount
}

// GetUserAccountOk returns a tuple with the UserAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTenantPlServiceUserRequest) GetUserAccountOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccount) {
		return nil, false
	}
	return o.UserAccount, true
}

// HasUserAccount returns a boolean if a field has been set.
func (o *AddTenantPlServiceUserRequest) HasUserAccount() bool {
	if o != nil && !IsNil(o.UserAccount) {
		return true
	}

	return false
}

// SetUserAccount gets a reference to the given string and assigns it to the UserAccount field.
func (o *AddTenantPlServiceUserRequest) SetUserAccount(v string) {
	o.UserAccount = &v
}

func (o AddTenantPlServiceUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddTenantPlServiceUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAccount) {
		toSerialize["userAccount"] = o.UserAccount
	}
	return toSerialize, nil
}

type NullableAddTenantPlServiceUserRequest struct {
	value *AddTenantPlServiceUserRequest
	isSet bool
}

func (v NullableAddTenantPlServiceUserRequest) Get() *AddTenantPlServiceUserRequest {
	return v.value
}

func (v *NullableAddTenantPlServiceUserRequest) Set(val *AddTenantPlServiceUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTenantPlServiceUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTenantPlServiceUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTenantPlServiceUserRequest(val *AddTenantPlServiceUserRequest) *NullableAddTenantPlServiceUserRequest {
	return &NullableAddTenantPlServiceUserRequest{value: val, isSet: true}
}

func (v NullableAddTenantPlServiceUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTenantPlServiceUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


