// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// VirtualMachineReplicaSetReplicaFailure is added in a replica set when
	// one of its VMs fails to be created due to insufficient quota, limit
	// ranges, security policy, host selection, or deleted due to the host being
	// down or finalizers are failing.
	VirtualMachineReplicaSetReplicaFailure = "ReplicaFailure"
)

// VirtualMachineTemplateSpec describes the data a VM should have when created
// from a template.
type VirtualMachineTemplateSpec struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VirtualMachineSpec `json:"spec,omitempty"`
}

// VirtualMachineReplicaSetSpec is the specification of a
// VirtualMachineReplicaSet.
type VirtualMachineReplicaSetSpec struct {
	// Replicas is the number of desired replicas.
	// This is a pointer to distinguish between explicit zero and unspecified.
	//
	// +optional
	// +kubebuilder:default=1
	Replicas *int32 `json:"replicas,omitempty"`

	// MinReadySeconds is the minimum number of seconds for which a newly
	// created VM should be ready for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	//
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// Selector is a label query over VMs that should match the replica count.
	// Label keys and values that must match in order to be controlled by this
	// replica set.
	//
	// It must match the VM template's labels.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *metav1.LabelSelector `json:"selector"`

	// Template is the object that describes the VM that will be created if
	// insufficient replicas are detected.
	//
	// +optional
	Template VirtualMachineTemplateSpec `json:"template,omitempty"`
}

// VirtualMachineReplicaSetStatus represents the observed state of a
// VirtualMachineReplicaSet resource.
type VirtualMachineReplicaSetStatus struct {
	// Replicas is the most recently observed number of replicas.
	Replicas int32 `json:"replicas"`

	// FullyLabeledReplicas is the number of VMs that have labels matching the
	// labels of the VM template of the replica set.
	//
	// +optional
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty"`

	// ReadyReplicas is the number of VMs targeted by this replica set with a
	// Ready Condition.
	//
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// AvailableReplicas is the number of available replicas (ready for at
	// least minReadySeconds) for this replica set.
	//
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed
	// VirtualMachineReplicaSet.
	//
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions represents the latest available observations of a replica
	// set's current state.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vmreplicaset
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Replicas",type="integer",priority=1,JSONPath=".status.replicas"
// +kubebuilder:printcolumn:name="Ready-Replicas",type="integer",priority=1,JSONPath=".status.readyReplicas"
// +kubebuilder:printcolumn:name="Available-Replicas",type="integer",JSONPath=".status.availableReplicas"

// VirtualMachineReplicaSet is the schema for the virtualmachinereplicasets API
// and represents the desired state and observed status of a
// virtualmachinereplicasets resource.
type VirtualMachineReplicaSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineReplicaSetSpec   `json:"spec,omitempty"`
	Status VirtualMachineReplicaSetStatus `json:"status,omitempty"`
}

func (vmrs VirtualMachineReplicaSet) NamespacedName() string {
	return vmrs.Namespace + "/" + vmrs.Name
}

// +kubebuilder:object:root=true

// VirtualMachineReplicaSetList contains a list of VirtualMachineReplicaSet.
type VirtualMachineReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineReplicaSet `json:"items"`
}

func init() {
	RegisterTypeWithScheme(
		&VirtualMachineReplicaSet{},
		&VirtualMachineReplicaSetList{},
	)
}
