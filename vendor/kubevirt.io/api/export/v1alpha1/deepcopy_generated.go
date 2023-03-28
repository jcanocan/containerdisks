//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The KubeVirt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExport) DeepCopyInto(out *VirtualMachineExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VirtualMachineExportStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExport.
func (in *VirtualMachineExport) DeepCopy() *VirtualMachineExport {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportLink) DeepCopyInto(out *VirtualMachineExportLink) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VirtualMachineExportVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportLink.
func (in *VirtualMachineExportLink) DeepCopy() *VirtualMachineExportLink {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportLinks) DeepCopyInto(out *VirtualMachineExportLinks) {
	*out = *in
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = new(VirtualMachineExportLink)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(VirtualMachineExportLink)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportLinks.
func (in *VirtualMachineExportLinks) DeepCopy() *VirtualMachineExportLinks {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportLinks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportList) DeepCopyInto(out *VirtualMachineExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportList.
func (in *VirtualMachineExportList) DeepCopy() *VirtualMachineExportList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportSpec) DeepCopyInto(out *VirtualMachineExportSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(string)
		**out = **in
	}
	if in.TTLDuration != nil {
		in, out := &in.TTLDuration, &out.TTLDuration
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportSpec.
func (in *VirtualMachineExportSpec) DeepCopy() *VirtualMachineExportSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportStatus) DeepCopyInto(out *VirtualMachineExportStatus) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = new(VirtualMachineExportLinks)
		(*in).DeepCopyInto(*out)
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(string)
		**out = **in
	}
	if in.TTLExpirationTime != nil {
		in, out := &in.TTLExpirationTime, &out.TTLExpirationTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportStatus.
func (in *VirtualMachineExportStatus) DeepCopy() *VirtualMachineExportStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportVolume) DeepCopyInto(out *VirtualMachineExportVolume) {
	*out = *in
	if in.Formats != nil {
		in, out := &in.Formats, &out.Formats
		*out = make([]VirtualMachineExportVolumeFormat, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportVolume.
func (in *VirtualMachineExportVolume) DeepCopy() *VirtualMachineExportVolume {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineExportVolumeFormat) DeepCopyInto(out *VirtualMachineExportVolumeFormat) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineExportVolumeFormat.
func (in *VirtualMachineExportVolumeFormat) DeepCopy() *VirtualMachineExportVolumeFormat {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineExportVolumeFormat)
	in.DeepCopyInto(out)
	return out
}