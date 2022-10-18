// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineDefaults is the schema for the virtualmachinedefaults API and
// represents the desired state and observed status of a virtualmachinedefaults
// resource.
//
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vmdefaults
// +kubebuilder:storageversion
type VirtualMachineDefaults struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VirtualMachineSpec `json:"spec,omitempty"`
}

func (vm VirtualMachineDefaults) NamespacedName() string {
	return vm.Namespace + "/" + vm.Name
}

// VirtualMachineDefaultsList contains a list of VirtualMachineDefaults.
//
// +kubebuilder:object:root=true
type VirtualMachineDefaultsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineDefaults `json:"items"`
}

func init() {
	RegisterTypeWithScheme(
		&VirtualMachineDefaults{},
		&VirtualMachineDefaultsList{},
	)
}
