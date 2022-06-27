// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespace,shortName=vmclassinstance
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// VirtualMachineClassInstance is the Schema for the
// virtualmachineclassinstances API.
//
// A VirtualMachineClassInstance is a deep-copy of a VirtualMachineClass
// resource's spec and status. An instance of this resource is created per VM
// that references a given VM Class. This ensures that VMs are not affected by
// changes to VM Classes.
type VirtualMachineClassInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineClassSpec   `json:"spec,omitempty"`
	Status VirtualMachineClassStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineClassInstanceList contains a list of VirtualMachineClassInstance.
type VirtualMachineClassInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineClassInstance `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&VirtualMachineClassInstance{}, &VirtualMachineClassInstanceList{})
}
