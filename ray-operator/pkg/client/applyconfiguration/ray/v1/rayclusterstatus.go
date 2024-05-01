// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RayClusterStatusApplyConfiguration represents an declarative configuration of the RayClusterStatus type for use
// with apply.
type RayClusterStatusApplyConfiguration struct {
	State                   *v1.ClusterState                 `json:"state,omitempty"`
	AvailableWorkerReplicas *int32                           `json:"availableWorkerReplicas,omitempty"`
	DesiredWorkerReplicas   *int32                           `json:"desiredWorkerReplicas,omitempty"`
	MinWorkerReplicas       *int32                           `json:"minWorkerReplicas,omitempty"`
	MaxWorkerReplicas       *int32                           `json:"maxWorkerReplicas,omitempty"`
	DesiredCPU              *resource.Quantity               `json:"desiredCPU,omitempty"`
	DesiredMemory           *resource.Quantity               `json:"desiredMemory,omitempty"`
	DesiredGPU              *resource.Quantity               `json:"desiredGPU,omitempty"`
	DesiredTPU              *resource.Quantity               `json:"desiredTPU,omitempty"`
	LastUpdateTime          *metav1.Time                     `json:"lastUpdateTime,omitempty"`
	StateTransitionTimes    map[v1.ClusterState]*metav1.Time `json:"stateTransitionTimes,omitempty"`
	Endpoints               map[string]string                `json:"endpoints,omitempty"`
	Head                    *HeadInfoApplyConfiguration      `json:"head,omitempty"`
	Reason                  *string                          `json:"reason,omitempty"`
	ObservedGeneration      *int64                           `json:"observedGeneration,omitempty"`
}

// RayClusterStatusApplyConfiguration constructs an declarative configuration of the RayClusterStatus type for use with
// apply.
func RayClusterStatus() *RayClusterStatusApplyConfiguration {
	return &RayClusterStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithState(value v1.ClusterState) *RayClusterStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithAvailableWorkerReplicas sets the AvailableWorkerReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableWorkerReplicas field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithAvailableWorkerReplicas(value int32) *RayClusterStatusApplyConfiguration {
	b.AvailableWorkerReplicas = &value
	return b
}

// WithDesiredWorkerReplicas sets the DesiredWorkerReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredWorkerReplicas field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithDesiredWorkerReplicas(value int32) *RayClusterStatusApplyConfiguration {
	b.DesiredWorkerReplicas = &value
	return b
}

// WithMinWorkerReplicas sets the MinWorkerReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinWorkerReplicas field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithMinWorkerReplicas(value int32) *RayClusterStatusApplyConfiguration {
	b.MinWorkerReplicas = &value
	return b
}

// WithMaxWorkerReplicas sets the MaxWorkerReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxWorkerReplicas field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithMaxWorkerReplicas(value int32) *RayClusterStatusApplyConfiguration {
	b.MaxWorkerReplicas = &value
	return b
}

// WithDesiredCPU sets the DesiredCPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredCPU field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithDesiredCPU(value resource.Quantity) *RayClusterStatusApplyConfiguration {
	b.DesiredCPU = &value
	return b
}

// WithDesiredMemory sets the DesiredMemory field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredMemory field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithDesiredMemory(value resource.Quantity) *RayClusterStatusApplyConfiguration {
	b.DesiredMemory = &value
	return b
}

// WithDesiredGPU sets the DesiredGPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredGPU field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithDesiredGPU(value resource.Quantity) *RayClusterStatusApplyConfiguration {
	b.DesiredGPU = &value
	return b
}

// WithDesiredTPU sets the DesiredTPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredTPU field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithDesiredTPU(value resource.Quantity) *RayClusterStatusApplyConfiguration {
	b.DesiredTPU = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithLastUpdateTime(value metav1.Time) *RayClusterStatusApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithStateTransitionTimes puts the entries into the StateTransitionTimes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the StateTransitionTimes field,
// overwriting an existing map entries in StateTransitionTimes field with the same key.
func (b *RayClusterStatusApplyConfiguration) WithStateTransitionTimes(entries map[v1.ClusterState]*metav1.Time) *RayClusterStatusApplyConfiguration {
	if b.StateTransitionTimes == nil && len(entries) > 0 {
		b.StateTransitionTimes = make(map[v1.ClusterState]*metav1.Time, len(entries))
	}
	for k, v := range entries {
		b.StateTransitionTimes[k] = v
	}
	return b
}

// WithEndpoints puts the entries into the Endpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Endpoints field,
// overwriting an existing map entries in Endpoints field with the same key.
func (b *RayClusterStatusApplyConfiguration) WithEndpoints(entries map[string]string) *RayClusterStatusApplyConfiguration {
	if b.Endpoints == nil && len(entries) > 0 {
		b.Endpoints = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Endpoints[k] = v
	}
	return b
}

// WithHead sets the Head field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Head field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithHead(value *HeadInfoApplyConfiguration) *RayClusterStatusApplyConfiguration {
	b.Head = value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithReason(value string) *RayClusterStatusApplyConfiguration {
	b.Reason = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *RayClusterStatusApplyConfiguration) WithObservedGeneration(value int64) *RayClusterStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}
