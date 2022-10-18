// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LocalObjectRef describes a reference to another object in the same
// namespace as the referrer.
type LocalObjectRef struct {
	// APIVersion defines the versioned schema of this representation of an
	// object. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion"`

	// Kind is a string value representing the REST resource this object
	// represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	// In CamelCase.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind"`

	// Name refers to a unique resource in the current namespace.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
}

// PartialObjectRef describes a reference to another object in the same
// namespace as the referrer. The reference can be just a name but may also
// include the referred resource's APIVersion and Kind.
type PartialObjectRef struct {
	metav1.TypeMeta `json:",inline"`

	// Name refers to a unique resource in the current namespace.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
}

// NameAndKindRef describes a reference to another object in the same
// namespace as the referrer. The reference can be just a name but may also
// include the referred resource's Kind.
type NameAndKindRef struct {
	// Kind is a string value representing the kind of resource to which this
	// object refers.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind,omitempty"`

	// Name refers to a unique resource in the current namespace.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
}

// ValueOrSecretKeySelector describes a value from either a SecretKeySelector
// or value directly in this object.
type ValueOrSecretKeySelector struct {
	// From is specified to reference a value from a Secret resource.
	//
	// Please note this field is mutually exclusive with the Value field.
	//
	// +optional
	From *corev1.SecretKeySelector `json:"from,omitempty"`

	// Value is used to directly specify a value.
	//
	// Please note this field is mutually exclusive with the From field.
	//
	// +optional
	Value *string `json:"value,omitempty"`
}

// KeyValuePair is useful when wanting to realize a map as a list of key/value
// pairs.
type KeyValuePair struct {
	// Key is the key part of the key/value pair.
	Key string `json:"key"`
	// Value is the optional value part of the key/value pair.
	// +optional
	Value string `json:"value,omitempty"`
}

// NameValuePair is useful when wanting to realize a map as a list of name/value
// pairs.
type NameValuePair struct {
	// Name is the name part of the name/value pair.
	Name string `json:"name"`
	// Value is the optional value part of the name/value pair.
	// +optional
	Value string `json:"value,omitempty"`
}
