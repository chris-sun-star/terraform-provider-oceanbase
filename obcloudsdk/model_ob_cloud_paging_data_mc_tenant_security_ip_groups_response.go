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

// checks if the OBCloudPagingDataMcTenantSecurityIpGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OBCloudPagingDataMcTenantSecurityIpGroupsResponse{}

// OBCloudPagingDataMcTenantSecurityIpGroupsResponse struct for OBCloudPagingDataMcTenantSecurityIpGroupsResponse
type OBCloudPagingDataMcTenantSecurityIpGroupsResponse struct {
	DataList []McTenantSecurityIpGroupsResponse `json:"dataList,omitempty"`
	Total *int64 `json:"total,omitempty"`
	HasMore *bool `json:"hasMore,omitempty"`
}

// NewOBCloudPagingDataMcTenantSecurityIpGroupsResponse instantiates a new OBCloudPagingDataMcTenantSecurityIpGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOBCloudPagingDataMcTenantSecurityIpGroupsResponse() *OBCloudPagingDataMcTenantSecurityIpGroupsResponse {
	this := OBCloudPagingDataMcTenantSecurityIpGroupsResponse{}
	return &this
}

// NewOBCloudPagingDataMcTenantSecurityIpGroupsResponseWithDefaults instantiates a new OBCloudPagingDataMcTenantSecurityIpGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOBCloudPagingDataMcTenantSecurityIpGroupsResponseWithDefaults() *OBCloudPagingDataMcTenantSecurityIpGroupsResponse {
	this := OBCloudPagingDataMcTenantSecurityIpGroupsResponse{}
	return &this
}

// GetDataList returns the DataList field value if set, zero value otherwise.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) GetDataList() []McTenantSecurityIpGroupsResponse {
	if o == nil || IsNil(o.DataList) {
		var ret []McTenantSecurityIpGroupsResponse
		return ret
	}
	return o.DataList
}

// GetDataListOk returns a tuple with the DataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) GetDataListOk() ([]McTenantSecurityIpGroupsResponse, bool) {
	if o == nil || IsNil(o.DataList) {
		return nil, false
	}
	return o.DataList, true
}

// HasDataList returns a boolean if a field has been set.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) HasDataList() bool {
	if o != nil && !IsNil(o.DataList) {
		return true
	}

	return false
}

// SetDataList gets a reference to the given []McTenantSecurityIpGroupsResponse and assigns it to the DataList field.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) SetDataList(v []McTenantSecurityIpGroupsResponse) {
	o.DataList = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o OBCloudPagingDataMcTenantSecurityIpGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OBCloudPagingDataMcTenantSecurityIpGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataList) {
		toSerialize["dataList"] = o.DataList
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	return toSerialize, nil
}

type NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse struct {
	value *OBCloudPagingDataMcTenantSecurityIpGroupsResponse
	isSet bool
}

func (v NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse) Get() *OBCloudPagingDataMcTenantSecurityIpGroupsResponse {
	return v.value
}

func (v *NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse) Set(val *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse(val *OBCloudPagingDataMcTenantSecurityIpGroupsResponse) *NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse {
	return &NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse{value: val, isSet: true}
}

func (v NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOBCloudPagingDataMcTenantSecurityIpGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


