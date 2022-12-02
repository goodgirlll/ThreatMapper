/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ControlsAgentControls struct for ControlsAgentControls
type ControlsAgentControls struct {
	Beatrate *int32 `json:"beatrate,omitempty"`
	Commands []map[string]interface{} `json:"commands,omitempty"`
}

// NewControlsAgentControls instantiates a new ControlsAgentControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsAgentControls() *ControlsAgentControls {
	this := ControlsAgentControls{}
	return &this
}

// NewControlsAgentControlsWithDefaults instantiates a new ControlsAgentControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsAgentControlsWithDefaults() *ControlsAgentControls {
	this := ControlsAgentControls{}
	return &this
}

// GetBeatrate returns the Beatrate field value if set, zero value otherwise.
func (o *ControlsAgentControls) GetBeatrate() int32 {
	if o == nil || isNil(o.Beatrate) {
		var ret int32
		return ret
	}
	return *o.Beatrate
}

// GetBeatrateOk returns a tuple with the Beatrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsAgentControls) GetBeatrateOk() (*int32, bool) {
	if o == nil || isNil(o.Beatrate) {
    return nil, false
	}
	return o.Beatrate, true
}

// HasBeatrate returns a boolean if a field has been set.
func (o *ControlsAgentControls) HasBeatrate() bool {
	if o != nil && !isNil(o.Beatrate) {
		return true
	}

	return false
}

// SetBeatrate gets a reference to the given int32 and assigns it to the Beatrate field.
func (o *ControlsAgentControls) SetBeatrate(v int32) {
	o.Beatrate = &v
}

// GetCommands returns the Commands field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlsAgentControls) GetCommands() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlsAgentControls) GetCommandsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Commands) {
    return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *ControlsAgentControls) HasCommands() bool {
	if o != nil && isNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []map[string]interface{} and assigns it to the Commands field.
func (o *ControlsAgentControls) SetCommands(v []map[string]interface{}) {
	o.Commands = v
}

func (o ControlsAgentControls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Beatrate) {
		toSerialize["beatrate"] = o.Beatrate
	}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableControlsAgentControls struct {
	value *ControlsAgentControls
	isSet bool
}

func (v NullableControlsAgentControls) Get() *ControlsAgentControls {
	return v.value
}

func (v *NullableControlsAgentControls) Set(val *ControlsAgentControls) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsAgentControls) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsAgentControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsAgentControls(val *ControlsAgentControls) *NullableControlsAgentControls {
	return &NullableControlsAgentControls{value: val, isSet: true}
}

func (v NullableControlsAgentControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsAgentControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

