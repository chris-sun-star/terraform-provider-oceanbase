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

// checks if the ModifyUserStatusParamDo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyUserStatusParamDo{}

// ModifyUserStatusParamDo struct for ModifyUserStatusParamDo
type ModifyUserStatusParamDo struct {
	RequestId *string `json:"requestId,omitempty"`
	CallerType *string `json:"callerType,omitempty"`
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	SecurityToken *string `json:"securityToken,omitempty"`
	CallerUid *string `json:"callerUid,omitempty"`
	CallerBid *string `json:"callerBid,omitempty"`
	StsTokenCallerUid *string `json:"stsTokenCallerUid,omitempty"`
	StsTokenCallerBid *string `json:"stsTokenCallerBid,omitempty"`
	AcceptLanguage *string `json:"acceptLanguage,omitempty"`
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	MergedCallerBid *string `json:"mergedCallerBid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	UserName *string `json:"userName,omitempty"`
	UserStatus *string `json:"userStatus,omitempty"`
	ClusterGroupName *string `json:"clusterGroupName,omitempty"`
	MergedCallerUid *string `json:"mergedCallerUid,omitempty"`
	UserID *string `json:"user_ID,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewModifyUserStatusParamDo instantiates a new ModifyUserStatusParamDo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyUserStatusParamDo() *ModifyUserStatusParamDo {
	this := ModifyUserStatusParamDo{}
	return &this
}

// NewModifyUserStatusParamDoWithDefaults instantiates a new ModifyUserStatusParamDo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyUserStatusParamDoWithDefaults() *ModifyUserStatusParamDo {
	this := ModifyUserStatusParamDo{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyUserStatusParamDo) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCallerType returns the CallerType field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetCallerType() string {
	if o == nil || IsNil(o.CallerType) {
		var ret string
		return ret
	}
	return *o.CallerType
}

// GetCallerTypeOk returns a tuple with the CallerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetCallerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CallerType) {
		return nil, false
	}
	return o.CallerType, true
}

// HasCallerType returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasCallerType() bool {
	if o != nil && !IsNil(o.CallerType) {
		return true
	}

	return false
}

// SetCallerType gets a reference to the given string and assigns it to the CallerType field.
func (o *ModifyUserStatusParamDo) SetCallerType(v string) {
	o.CallerType = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *ModifyUserStatusParamDo) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetSecurityToken() string {
	if o == nil || IsNil(o.SecurityToken) {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetSecurityTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityToken) {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasSecurityToken() bool {
	if o != nil && !IsNil(o.SecurityToken) {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *ModifyUserStatusParamDo) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetCallerUid returns the CallerUid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetCallerUid() string {
	if o == nil || IsNil(o.CallerUid) {
		var ret string
		return ret
	}
	return *o.CallerUid
}

// GetCallerUidOk returns a tuple with the CallerUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetCallerUidOk() (*string, bool) {
	if o == nil || IsNil(o.CallerUid) {
		return nil, false
	}
	return o.CallerUid, true
}

// HasCallerUid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasCallerUid() bool {
	if o != nil && !IsNil(o.CallerUid) {
		return true
	}

	return false
}

// SetCallerUid gets a reference to the given string and assigns it to the CallerUid field.
func (o *ModifyUserStatusParamDo) SetCallerUid(v string) {
	o.CallerUid = &v
}

// GetCallerBid returns the CallerBid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetCallerBid() string {
	if o == nil || IsNil(o.CallerBid) {
		var ret string
		return ret
	}
	return *o.CallerBid
}

// GetCallerBidOk returns a tuple with the CallerBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetCallerBidOk() (*string, bool) {
	if o == nil || IsNil(o.CallerBid) {
		return nil, false
	}
	return o.CallerBid, true
}

// HasCallerBid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasCallerBid() bool {
	if o != nil && !IsNil(o.CallerBid) {
		return true
	}

	return false
}

// SetCallerBid gets a reference to the given string and assigns it to the CallerBid field.
func (o *ModifyUserStatusParamDo) SetCallerBid(v string) {
	o.CallerBid = &v
}

// GetStsTokenCallerUid returns the StsTokenCallerUid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetStsTokenCallerUid() string {
	if o == nil || IsNil(o.StsTokenCallerUid) {
		var ret string
		return ret
	}
	return *o.StsTokenCallerUid
}

// GetStsTokenCallerUidOk returns a tuple with the StsTokenCallerUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetStsTokenCallerUidOk() (*string, bool) {
	if o == nil || IsNil(o.StsTokenCallerUid) {
		return nil, false
	}
	return o.StsTokenCallerUid, true
}

// HasStsTokenCallerUid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasStsTokenCallerUid() bool {
	if o != nil && !IsNil(o.StsTokenCallerUid) {
		return true
	}

	return false
}

// SetStsTokenCallerUid gets a reference to the given string and assigns it to the StsTokenCallerUid field.
func (o *ModifyUserStatusParamDo) SetStsTokenCallerUid(v string) {
	o.StsTokenCallerUid = &v
}

// GetStsTokenCallerBid returns the StsTokenCallerBid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetStsTokenCallerBid() string {
	if o == nil || IsNil(o.StsTokenCallerBid) {
		var ret string
		return ret
	}
	return *o.StsTokenCallerBid
}

// GetStsTokenCallerBidOk returns a tuple with the StsTokenCallerBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetStsTokenCallerBidOk() (*string, bool) {
	if o == nil || IsNil(o.StsTokenCallerBid) {
		return nil, false
	}
	return o.StsTokenCallerBid, true
}

// HasStsTokenCallerBid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasStsTokenCallerBid() bool {
	if o != nil && !IsNil(o.StsTokenCallerBid) {
		return true
	}

	return false
}

// SetStsTokenCallerBid gets a reference to the given string and assigns it to the StsTokenCallerBid field.
func (o *ModifyUserStatusParamDo) SetStsTokenCallerBid(v string) {
	o.StsTokenCallerBid = &v
}

// GetAcceptLanguage returns the AcceptLanguage field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetAcceptLanguage() string {
	if o == nil || IsNil(o.AcceptLanguage) {
		var ret string
		return ret
	}
	return *o.AcceptLanguage
}

// GetAcceptLanguageOk returns a tuple with the AcceptLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetAcceptLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptLanguage) {
		return nil, false
	}
	return o.AcceptLanguage, true
}

// HasAcceptLanguage returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasAcceptLanguage() bool {
	if o != nil && !IsNil(o.AcceptLanguage) {
		return true
	}

	return false
}

// SetAcceptLanguage gets a reference to the given string and assigns it to the AcceptLanguage field.
func (o *ModifyUserStatusParamDo) SetAcceptLanguage(v string) {
	o.AcceptLanguage = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ModifyUserStatusParamDo) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ModifyUserStatusParamDo) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetMergedCallerBid returns the MergedCallerBid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetMergedCallerBid() string {
	if o == nil || IsNil(o.MergedCallerBid) {
		var ret string
		return ret
	}
	return *o.MergedCallerBid
}

// GetMergedCallerBidOk returns a tuple with the MergedCallerBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetMergedCallerBidOk() (*string, bool) {
	if o == nil || IsNil(o.MergedCallerBid) {
		return nil, false
	}
	return o.MergedCallerBid, true
}

// HasMergedCallerBid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasMergedCallerBid() bool {
	if o != nil && !IsNil(o.MergedCallerBid) {
		return true
	}

	return false
}

// SetMergedCallerBid gets a reference to the given string and assigns it to the MergedCallerBid field.
func (o *ModifyUserStatusParamDo) SetMergedCallerBid(v string) {
	o.MergedCallerBid = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ModifyUserStatusParamDo) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ModifyUserStatusParamDo) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ModifyUserStatusParamDo) SetUserName(v string) {
	o.UserName = &v
}

// GetUserStatus returns the UserStatus field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetUserStatus() string {
	if o == nil || IsNil(o.UserStatus) {
		var ret string
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetUserStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UserStatus) {
		return nil, false
	}
	return o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasUserStatus() bool {
	if o != nil && !IsNil(o.UserStatus) {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given string and assigns it to the UserStatus field.
func (o *ModifyUserStatusParamDo) SetUserStatus(v string) {
	o.UserStatus = &v
}

// GetClusterGroupName returns the ClusterGroupName field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetClusterGroupName() string {
	if o == nil || IsNil(o.ClusterGroupName) {
		var ret string
		return ret
	}
	return *o.ClusterGroupName
}

// GetClusterGroupNameOk returns a tuple with the ClusterGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetClusterGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterGroupName) {
		return nil, false
	}
	return o.ClusterGroupName, true
}

// HasClusterGroupName returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasClusterGroupName() bool {
	if o != nil && !IsNil(o.ClusterGroupName) {
		return true
	}

	return false
}

// SetClusterGroupName gets a reference to the given string and assigns it to the ClusterGroupName field.
func (o *ModifyUserStatusParamDo) SetClusterGroupName(v string) {
	o.ClusterGroupName = &v
}

// GetMergedCallerUid returns the MergedCallerUid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetMergedCallerUid() string {
	if o == nil || IsNil(o.MergedCallerUid) {
		var ret string
		return ret
	}
	return *o.MergedCallerUid
}

// GetMergedCallerUidOk returns a tuple with the MergedCallerUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetMergedCallerUidOk() (*string, bool) {
	if o == nil || IsNil(o.MergedCallerUid) {
		return nil, false
	}
	return o.MergedCallerUid, true
}

// HasMergedCallerUid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasMergedCallerUid() bool {
	if o != nil && !IsNil(o.MergedCallerUid) {
		return true
	}

	return false
}

// SetMergedCallerUid gets a reference to the given string and assigns it to the MergedCallerUid field.
func (o *ModifyUserStatusParamDo) SetMergedCallerUid(v string) {
	o.MergedCallerUid = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetUserID() string {
	if o == nil || IsNil(o.UserID) {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.UserID) {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasUserID() bool {
	if o != nil && !IsNil(o.UserID) {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *ModifyUserStatusParamDo) SetUserID(v string) {
	o.UserID = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ModifyUserStatusParamDo) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserStatusParamDo) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ModifyUserStatusParamDo) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ModifyUserStatusParamDo) SetUid(v string) {
	o.Uid = &v
}

func (o ModifyUserStatusParamDo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyUserStatusParamDo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.CallerType) {
		toSerialize["callerType"] = o.CallerType
	}
	if !IsNil(o.AccessKeyId) {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if !IsNil(o.SecurityToken) {
		toSerialize["securityToken"] = o.SecurityToken
	}
	if !IsNil(o.CallerUid) {
		toSerialize["callerUid"] = o.CallerUid
	}
	if !IsNil(o.CallerBid) {
		toSerialize["callerBid"] = o.CallerBid
	}
	if !IsNil(o.StsTokenCallerUid) {
		toSerialize["stsTokenCallerUid"] = o.StsTokenCallerUid
	}
	if !IsNil(o.StsTokenCallerBid) {
		toSerialize["stsTokenCallerBid"] = o.StsTokenCallerBid
	}
	if !IsNil(o.AcceptLanguage) {
		toSerialize["acceptLanguage"] = o.AcceptLanguage
	}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.MergedCallerBid) {
		toSerialize["mergedCallerBid"] = o.MergedCallerBid
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.UserStatus) {
		toSerialize["userStatus"] = o.UserStatus
	}
	if !IsNil(o.ClusterGroupName) {
		toSerialize["clusterGroupName"] = o.ClusterGroupName
	}
	if !IsNil(o.MergedCallerUid) {
		toSerialize["mergedCallerUid"] = o.MergedCallerUid
	}
	if !IsNil(o.UserID) {
		toSerialize["user_ID"] = o.UserID
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableModifyUserStatusParamDo struct {
	value *ModifyUserStatusParamDo
	isSet bool
}

func (v NullableModifyUserStatusParamDo) Get() *ModifyUserStatusParamDo {
	return v.value
}

func (v *NullableModifyUserStatusParamDo) Set(val *ModifyUserStatusParamDo) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyUserStatusParamDo) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyUserStatusParamDo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyUserStatusParamDo(val *ModifyUserStatusParamDo) *NullableModifyUserStatusParamDo {
	return &NullableModifyUserStatusParamDo{value: val, isSet: true}
}

func (v NullableModifyUserStatusParamDo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyUserStatusParamDo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


