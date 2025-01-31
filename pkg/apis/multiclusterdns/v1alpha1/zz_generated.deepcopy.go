// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDNS) DeepCopyInto(out *ClusterDNS) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDNS.
func (in *ClusterDNS) DeepCopy() *ClusterDNS {
	if in == nil {
		return nil
	}
	out := new(ClusterDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServiceDNSRecord) DeepCopyInto(out *MultiClusterServiceDNSRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServiceDNSRecord.
func (in *MultiClusterServiceDNSRecord) DeepCopy() *MultiClusterServiceDNSRecord {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServiceDNSRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterServiceDNSRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServiceDNSRecordList) DeepCopyInto(out *MultiClusterServiceDNSRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterServiceDNSRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServiceDNSRecordList.
func (in *MultiClusterServiceDNSRecordList) DeepCopy() *MultiClusterServiceDNSRecordList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServiceDNSRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterServiceDNSRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServiceDNSRecordSpec) DeepCopyInto(out *MultiClusterServiceDNSRecordSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServiceDNSRecordSpec.
func (in *MultiClusterServiceDNSRecordSpec) DeepCopy() *MultiClusterServiceDNSRecordSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServiceDNSRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServiceDNSRecordStatus) DeepCopyInto(out *MultiClusterServiceDNSRecordStatus) {
	*out = *in
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]ClusterDNS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServiceDNSRecordStatus.
func (in *MultiClusterServiceDNSRecordStatus) DeepCopy() *MultiClusterServiceDNSRecordStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServiceDNSRecordStatus)
	in.DeepCopyInto(out)
	return out
}
