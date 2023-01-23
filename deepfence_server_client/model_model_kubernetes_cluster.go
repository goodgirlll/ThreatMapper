/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelKubernetesCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelKubernetesCluster{}

// ModelKubernetesCluster struct for ModelKubernetesCluster
type ModelKubernetesCluster struct {
	CloudMetadata map[string]interface{} `json:"cloud_metadata"`
	Containers []ModelHost `json:"containers"`
	Metrics ModelComputeMetrics `json:"metrics"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
}

// NewModelKubernetesCluster instantiates a new ModelKubernetesCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelKubernetesCluster(cloudMetadata map[string]interface{}, containers []ModelHost, metrics ModelComputeMetrics, nodeId string, nodeName string) *ModelKubernetesCluster {
	this := ModelKubernetesCluster{}
	this.CloudMetadata = cloudMetadata
	this.Containers = containers
	this.Metrics = metrics
	this.NodeId = nodeId
	this.NodeName = nodeName
	return &this
}

// NewModelKubernetesClusterWithDefaults instantiates a new ModelKubernetesCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelKubernetesClusterWithDefaults() *ModelKubernetesCluster {
	this := ModelKubernetesCluster{}
	return &this
}

// GetCloudMetadata returns the CloudMetadata field value
func (o *ModelKubernetesCluster) GetCloudMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CloudMetadata
}

// GetCloudMetadataOk returns a tuple with the CloudMetadata field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetCloudMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CloudMetadata, true
}

// SetCloudMetadata sets field value
func (o *ModelKubernetesCluster) SetCloudMetadata(v map[string]interface{}) {
	o.CloudMetadata = v
}

// GetContainers returns the Containers field value
// If the value is explicit nil, the zero value for []ModelHost will be returned
func (o *ModelKubernetesCluster) GetContainers() []ModelHost {
	if o == nil {
		var ret []ModelHost
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelKubernetesCluster) GetContainersOk() ([]ModelHost, bool) {
	if o == nil || isNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ModelKubernetesCluster) SetContainers(v []ModelHost) {
	o.Containers = v
}

// GetMetrics returns the Metrics field value
func (o *ModelKubernetesCluster) GetMetrics() ModelComputeMetrics {
	if o == nil {
		var ret ModelComputeMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetMetricsOk() (*ModelComputeMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *ModelKubernetesCluster) SetMetrics(v ModelComputeMetrics) {
	o.Metrics = v
}

// GetNodeId returns the NodeId field value
func (o *ModelKubernetesCluster) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelKubernetesCluster) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelKubernetesCluster) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelKubernetesCluster) SetNodeName(v string) {
	o.NodeName = v
}

func (o ModelKubernetesCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelKubernetesCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_metadata"] = o.CloudMetadata
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	toSerialize["metrics"] = o.Metrics
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	return toSerialize, nil
}

type NullableModelKubernetesCluster struct {
	value *ModelKubernetesCluster
	isSet bool
}

func (v NullableModelKubernetesCluster) Get() *ModelKubernetesCluster {
	return v.value
}

func (v *NullableModelKubernetesCluster) Set(val *ModelKubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableModelKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableModelKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelKubernetesCluster(val *ModelKubernetesCluster) *NullableModelKubernetesCluster {
	return &NullableModelKubernetesCluster{value: val, isSet: true}
}

func (v NullableModelKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

