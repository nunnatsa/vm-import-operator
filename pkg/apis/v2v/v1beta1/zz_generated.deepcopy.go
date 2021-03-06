// +build !ignore_autogenerated

/*
Copyright 2020 The vm import Authors.

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
// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeItem) DeepCopyInto(out *DataVolumeItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeItem.
func (in *DataVolumeItem) DeepCopy() *DataVolumeItem {
	if in == nil {
		return nil
	}
	out := new(DataVolumeItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkResourceMappingItem) DeepCopyInto(out *NetworkResourceMappingItem) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Target.DeepCopyInto(&out.Target)
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkResourceMappingItem.
func (in *NetworkResourceMappingItem) DeepCopy() *NetworkResourceMappingItem {
	if in == nil {
		return nil
	}
	out := new(NetworkResourceMappingItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectIdentifier) DeepCopyInto(out *ObjectIdentifier) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectIdentifier.
func (in *ObjectIdentifier) DeepCopy() *ObjectIdentifier {
	if in == nil {
		return nil
	}
	out := new(ObjectIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvirtMappings) DeepCopyInto(out *OvirtMappings) {
	*out = *in
	if in.NetworkMappings != nil {
		in, out := &in.NetworkMappings, &out.NetworkMappings
		*out = new([]NetworkResourceMappingItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]NetworkResourceMappingItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.StorageMappings != nil {
		in, out := &in.StorageMappings, &out.StorageMappings
		*out = new([]StorageResourceMappingItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]StorageResourceMappingItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.DiskMappings != nil {
		in, out := &in.DiskMappings, &out.DiskMappings
		*out = new([]StorageResourceMappingItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]StorageResourceMappingItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvirtMappings.
func (in *OvirtMappings) DeepCopy() *OvirtMappings {
	if in == nil {
		return nil
	}
	out := new(OvirtMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMapping) DeepCopyInto(out *ResourceMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMapping.
func (in *ResourceMapping) DeepCopy() *ResourceMapping {
	if in == nil {
		return nil
	}
	out := new(ResourceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMappingList) DeepCopyInto(out *ResourceMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMappingList.
func (in *ResourceMappingList) DeepCopy() *ResourceMappingList {
	if in == nil {
		return nil
	}
	out := new(ResourceMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMappingSpec) DeepCopyInto(out *ResourceMappingSpec) {
	*out = *in
	if in.OvirtMappings != nil {
		in, out := &in.OvirtMappings, &out.OvirtMappings
		*out = new(OvirtMappings)
		(*in).DeepCopyInto(*out)
	}
	if in.VmwareMappings != nil {
		in, out := &in.VmwareMappings, &out.VmwareMappings
		*out = new(VmwareMappings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMappingSpec.
func (in *ResourceMappingSpec) DeepCopy() *ResourceMappingSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMappingStatus) DeepCopyInto(out *ResourceMappingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMappingStatus.
func (in *ResourceMappingStatus) DeepCopy() *ResourceMappingStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResourceMappingItem) DeepCopyInto(out *StorageResourceMappingItem) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Target.DeepCopyInto(&out.Target)
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(v1.PersistentVolumeMode)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResourceMappingItem.
func (in *StorageResourceMappingItem) DeepCopy() *StorageResourceMappingItem {
	if in == nil {
		return nil
	}
	out := new(StorageResourceMappingItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImportConfig) DeepCopyInto(out *VMImportConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImportConfig.
func (in *VMImportConfig) DeepCopy() *VMImportConfig {
	if in == nil {
		return nil
	}
	out := new(VMImportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMImportConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImportConfigList) DeepCopyInto(out *VMImportConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VMImportConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImportConfigList.
func (in *VMImportConfigList) DeepCopy() *VMImportConfigList {
	if in == nil {
		return nil
	}
	out := new(VMImportConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMImportConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImportConfigSpec) DeepCopyInto(out *VMImportConfigSpec) {
	*out = *in
	in.Infra.DeepCopyInto(&out.Infra)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImportConfigSpec.
func (in *VMImportConfigSpec) DeepCopy() *VMImportConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VMImportConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImportConfigStatus) DeepCopyInto(out *VMImportConfigStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImportConfigStatus.
func (in *VMImportConfigStatus) DeepCopy() *VMImportConfigStatus {
	if in == nil {
		return nil
	}
	out := new(VMImportConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImport) DeepCopyInto(out *VirtualMachineImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImport.
func (in *VirtualMachineImport) DeepCopy() *VirtualMachineImport {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportCondition) DeepCopyInto(out *VirtualMachineImportCondition) {
	*out = *in
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.LastHeartbeatTime != nil {
		in, out := &in.LastHeartbeatTime, &out.LastHeartbeatTime
		*out = (*in).DeepCopy()
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportCondition.
func (in *VirtualMachineImportCondition) DeepCopy() *VirtualMachineImportCondition {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportList) DeepCopyInto(out *VirtualMachineImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportList.
func (in *VirtualMachineImportList) DeepCopy() *VirtualMachineImportList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportOvirtSourceSpec) DeepCopyInto(out *VirtualMachineImportOvirtSourceSpec) {
	*out = *in
	in.VM.DeepCopyInto(&out.VM)
	if in.Mappings != nil {
		in, out := &in.Mappings, &out.Mappings
		*out = new(OvirtMappings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportOvirtSourceSpec.
func (in *VirtualMachineImportOvirtSourceSpec) DeepCopy() *VirtualMachineImportOvirtSourceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportOvirtSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportOvirtSourceVMClusterSpec) DeepCopyInto(out *VirtualMachineImportOvirtSourceVMClusterSpec) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportOvirtSourceVMClusterSpec.
func (in *VirtualMachineImportOvirtSourceVMClusterSpec) DeepCopy() *VirtualMachineImportOvirtSourceVMClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportOvirtSourceVMClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportOvirtSourceVMSpec) DeepCopyInto(out *VirtualMachineImportOvirtSourceVMSpec) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(VirtualMachineImportOvirtSourceVMClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportOvirtSourceVMSpec.
func (in *VirtualMachineImportOvirtSourceVMSpec) DeepCopy() *VirtualMachineImportOvirtSourceVMSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportOvirtSourceVMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportSourceSpec) DeepCopyInto(out *VirtualMachineImportSourceSpec) {
	*out = *in
	if in.Ovirt != nil {
		in, out := &in.Ovirt, &out.Ovirt
		*out = new(VirtualMachineImportOvirtSourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Vmware != nil {
		in, out := &in.Vmware, &out.Vmware
		*out = new(VirtualMachineImportVmwareSourceSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportSourceSpec.
func (in *VirtualMachineImportSourceSpec) DeepCopy() *VirtualMachineImportSourceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportSpec) DeepCopyInto(out *VirtualMachineImportSpec) {
	*out = *in
	in.ProviderCredentialsSecret.DeepCopyInto(&out.ProviderCredentialsSecret)
	if in.ResourceMapping != nil {
		in, out := &in.ResourceMapping, &out.ResourceMapping
		*out = new(ObjectIdentifier)
		(*in).DeepCopyInto(*out)
	}
	in.Source.DeepCopyInto(&out.Source)
	if in.TargetVMName != nil {
		in, out := &in.TargetVMName, &out.TargetVMName
		*out = new(string)
		**out = **in
	}
	if in.StartVM != nil {
		in, out := &in.StartVM, &out.StartVM
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportSpec.
func (in *VirtualMachineImportSpec) DeepCopy() *VirtualMachineImportSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportStatus) DeepCopyInto(out *VirtualMachineImportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VirtualMachineImportCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataVolumes != nil {
		in, out := &in.DataVolumes, &out.DataVolumes
		*out = make([]DataVolumeItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportStatus.
func (in *VirtualMachineImportStatus) DeepCopy() *VirtualMachineImportStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportVmwareSourceSpec) DeepCopyInto(out *VirtualMachineImportVmwareSourceSpec) {
	*out = *in
	in.VM.DeepCopyInto(&out.VM)
	if in.Mappings != nil {
		in, out := &in.Mappings, &out.Mappings
		*out = new(VmwareMappings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportVmwareSourceSpec.
func (in *VirtualMachineImportVmwareSourceSpec) DeepCopy() *VirtualMachineImportVmwareSourceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportVmwareSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineImportVmwareSourceVMSpec) DeepCopyInto(out *VirtualMachineImportVmwareSourceVMSpec) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineImportVmwareSourceVMSpec.
func (in *VirtualMachineImportVmwareSourceVMSpec) DeepCopy() *VirtualMachineImportVmwareSourceVMSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineImportVmwareSourceVMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmwareMappings) DeepCopyInto(out *VmwareMappings) {
	*out = *in
	if in.NetworkMappings != nil {
		in, out := &in.NetworkMappings, &out.NetworkMappings
		*out = new([]NetworkResourceMappingItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]NetworkResourceMappingItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.StorageMappings != nil {
		in, out := &in.StorageMappings, &out.StorageMappings
		*out = new([]StorageResourceMappingItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]StorageResourceMappingItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.DiskMappings != nil {
		in, out := &in.DiskMappings, &out.DiskMappings
		*out = new([]StorageResourceMappingItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]StorageResourceMappingItem, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmwareMappings.
func (in *VmwareMappings) DeepCopy() *VmwareMappings {
	if in == nil {
		return nil
	}
	out := new(VmwareMappings)
	in.DeepCopyInto(out)
	return out
}
