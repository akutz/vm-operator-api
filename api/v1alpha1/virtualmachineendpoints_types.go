// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// VirtualMachineEndpointConditionAddressesValid is the Type for a
	// VirtualMachineEndpoints resource's status condition.
	//
	// The condition's status is set to true only when the addresses from the
	// VirtualMachineEndpoints resource's subsets have been validated as
	// belonging to the specified VirtualMachine resource.
	VirtualMachineEndpointConditionAddressesValid = "AddressesValid"
)

// VirtualMachineEndpointAddress is a tuple that describes single IP address.
type VirtualMachineEndpointAddress struct {
	// IP is the IP of this endpoint.
	//
	// When omitted, the IP from the VirtualMachine's Status.VmIP field will be
	// used.
	//
	// Otherwise any IP specified is accepted as long as it is reported by the
	// VirtualMachine specified by the NodeName field.
	//
	// +optional
	IP string `json:"ip,omitempty"`

	// NodeName is the name of the node hosting this endpoint. This should be
	// the name of the VirtualMachine resource on Supervisor to which the
	// endpoint belongs. It is possible to look up the name of the
	// VirtualMachine resource by matching the node's provider ID with the
	// VirtualMachine's BIOS UUID.
	NodeName string `json:"nodeName"`
}

// VirtualMachineEndpointPort is a tuple that describes a single port.
type VirtualMachineEndpointPort struct {
	// Name of this port (corresponds to ServicePort.Name).
	// Optional if only one port is defined.
	// Must be a DNS_LABEL.
	//
	// +optional
	Name string `json:"name,omitempty"`

	// Port is the port number.
	Port int32 `json:"port"`

	// Protocol is the IP protocol for this port.
	//
	// +kubebuilder:default=TCP
	// +kubebuilder:validation:Enum=TCP;UDP;SCTP
	Protocol string `json:"protocol"`

	// AppProtocol is the application protocol for this port.
	// This field follows standard Kubernetes label syntax.
	// Un-prefixed names are reserved for IANA standard service names (as per
	// RFC-6335 and https://www.iana.org/assignments/service-names).
	// Non-standard protocols should use prefixed names such as
	// mycompany.com/my-custom-protocol.
	//
	// +optional
	AppProtocol *string `json:"appProtocol,omitempty"`
}

// VirtualMachineEndpointSubset is a group of addresses with a common set of
// ports. The expanded set of endpoints is the Cartesian product of Addresses x
// Ports.
type VirtualMachineEndpointSubset struct {
	Addresses []VirtualMachineEndpointAddress `json:"addresses"`
	Ports     []VirtualMachineEndpointPort    `json:"ports"`
}

// VirtualMachineEndpointsSpec defines the desired state of a
// VirtualMachineEndpoints resource.
type VirtualMachineEndpointsSpec struct {
	// Subsets is the union of all subsets.
	// +optional
	Subsets []VirtualMachineEndpointSubset `json:"subsets,omitempty"`
}

// VirtualMachineEndpointsStatus defines the observed state of a
// VirtualMachineEndpoints resource.
type VirtualMachineEndpointsStatus struct {
	// Conditions is a list of the latest, available observations of the
	// resource's current state.
	//
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=vmendpoints
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// VirtualMachineEndpoints is the schema for the virtualmachineendpoints API.
// VirtualMachineEndpoints is a collection of endpoints that implement the
// actual service.
type VirtualMachineEndpoints struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineEndpointsSpec   `json:"spec,omitempty"`
	Status VirtualMachineEndpointsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineEndpointsList contains a list of VirtualMachineEndpoints.
type VirtualMachineEndpointsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineEndpoints `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&VirtualMachineEndpoints{}, &VirtualMachineEndpointsList{})
}
