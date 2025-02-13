//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfig) DeepCopyInto(out *IPXEBootConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfig.
func (in *IPXEBootConfig) DeepCopy() *IPXEBootConfig {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPXEBootConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfigList) DeepCopyInto(out *IPXEBootConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPXEBootConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfigList.
func (in *IPXEBootConfigList) DeepCopy() *IPXEBootConfigList {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPXEBootConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfigSpec) DeepCopyInto(out *IPXEBootConfigSpec) {
	*out = *in
	if in.IgnitionSecretRef != nil {
		in, out := &in.IgnitionSecretRef, &out.IgnitionSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfigSpec.
func (in *IPXEBootConfigSpec) DeepCopy() *IPXEBootConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfigStatus) DeepCopyInto(out *IPXEBootConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfigStatus.
func (in *IPXEBootConfigStatus) DeepCopy() *IPXEBootConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfigStatus)
	in.DeepCopyInto(out)
	return out
}
