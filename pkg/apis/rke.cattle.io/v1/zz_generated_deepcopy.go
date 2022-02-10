// +build !ignore_autogenerated

/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	genericcondition "github.com/rancher/wrangler/pkg/genericcondition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUpgradeStrategy) DeepCopyInto(out *ClusterUpgradeStrategy) {
	*out = *in
	in.ControlPlaneDrainOptions.DeepCopyInto(&out.ControlPlaneDrainOptions)
	in.WorkerDrainOptions.DeepCopyInto(&out.WorkerDrainOptions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUpgradeStrategy.
func (in *ClusterUpgradeStrategy) DeepCopy() *ClusterUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(ClusterUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMachine) DeepCopyInto(out *CustomMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMachine.
func (in *CustomMachine) DeepCopy() *CustomMachine {
	if in == nil {
		return nil
	}
	out := new(CustomMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMachineList) DeepCopyInto(out *CustomMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMachineList.
func (in *CustomMachineList) DeepCopy() *CustomMachineList {
	if in == nil {
		return nil
	}
	out := new(CustomMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMachineSpec) DeepCopyInto(out *CustomMachineSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMachineSpec.
func (in *CustomMachineSpec) DeepCopy() *CustomMachineSpec {
	if in == nil {
		return nil
	}
	out := new(CustomMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMachineStatus) DeepCopyInto(out *CustomMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1beta1.MachineAddress, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMachineStatus.
func (in *CustomMachineStatus) DeepCopy() *CustomMachineStatus {
	if in == nil {
		return nil
	}
	out := new(CustomMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrainHook) DeepCopyInto(out *DrainHook) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrainHook.
func (in *DrainHook) DeepCopy() *DrainHook {
	if in == nil {
		return nil
	}
	out := new(DrainHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrainOptions) DeepCopyInto(out *DrainOptions) {
	*out = *in
	if in.IgnoreDaemonSets != nil {
		in, out := &in.IgnoreDaemonSets, &out.IgnoreDaemonSets
		*out = new(bool)
		**out = **in
	}
	if in.PreDrainHooks != nil {
		in, out := &in.PreDrainHooks, &out.PreDrainHooks
		*out = make([]DrainHook, len(*in))
		copy(*out, *in)
	}
	if in.PostDrainHooks != nil {
		in, out := &in.PostDrainHooks, &out.PostDrainHooks
		*out = make([]DrainHook, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrainOptions.
func (in *DrainOptions) DeepCopy() *DrainOptions {
	if in == nil {
		return nil
	}
	out := new(DrainOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCD) DeepCopyInto(out *ETCD) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(ETCDSnapshotS3)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCD.
func (in *ETCD) DeepCopy() *ETCD {
	if in == nil {
		return nil
	}
	out := new(ETCD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshot) DeepCopyInto(out *ETCDSnapshot) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(ETCDSnapshotS3)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshot.
func (in *ETCDSnapshot) DeepCopy() *ETCDSnapshot {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotCreate) DeepCopyInto(out *ETCDSnapshotCreate) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(ETCDSnapshotS3)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotCreate.
func (in *ETCDSnapshotCreate) DeepCopy() *ETCDSnapshotCreate {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotCreate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotRestore) DeepCopyInto(out *ETCDSnapshotRestore) {
	*out = *in
	in.ETCDSnapshot.DeepCopyInto(&out.ETCDSnapshot)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotRestore.
func (in *ETCDSnapshotRestore) DeepCopy() *ETCDSnapshotRestore {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDSnapshotS3) DeepCopyInto(out *ETCDSnapshotS3) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDSnapshotS3.
func (in *ETCDSnapshotS3) DeepCopy() *ETCDSnapshotS3 {
	if in == nil {
		return nil
	}
	out := new(ETCDSnapshotS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericMap.
func (in *GenericMap) DeepCopy() *GenericMap {
	if in == nil {
		return nil
	}
	out := new(GenericMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterAuthEndpoint) DeepCopyInto(out *LocalClusterAuthEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterAuthEndpoint.
func (in *LocalClusterAuthEndpoint) DeepCopy() *LocalClusterAuthEndpoint {
	if in == nil {
		return nil
	}
	out := new(LocalClusterAuthEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mirror) DeepCopyInto(out *Mirror) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Rewrites != nil {
		in, out := &in.Rewrites, &out.Rewrites
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mirror.
func (in *Mirror) DeepCopy() *Mirror {
	if in == nil {
		return nil
	}
	out := new(Mirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrap) DeepCopyInto(out *RKEBootstrap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrap.
func (in *RKEBootstrap) DeepCopy() *RKEBootstrap {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapList) DeepCopyInto(out *RKEBootstrapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKEBootstrap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapList.
func (in *RKEBootstrapList) DeepCopy() *RKEBootstrapList {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapSpec) DeepCopyInto(out *RKEBootstrapSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapSpec.
func (in *RKEBootstrapSpec) DeepCopy() *RKEBootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapStatus) DeepCopyInto(out *RKEBootstrapStatus) {
	*out = *in
	if in.DataSecretName != nil {
		in, out := &in.DataSecretName, &out.DataSecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapStatus.
func (in *RKEBootstrapStatus) DeepCopy() *RKEBootstrapStatus {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapTemplate) DeepCopyInto(out *RKEBootstrapTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapTemplate.
func (in *RKEBootstrapTemplate) DeepCopy() *RKEBootstrapTemplate {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrapTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapTemplateList) DeepCopyInto(out *RKEBootstrapTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKEBootstrapTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapTemplateList.
func (in *RKEBootstrapTemplateList) DeepCopy() *RKEBootstrapTemplateList {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEBootstrapTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEBootstrapTemplateSpec) DeepCopyInto(out *RKEBootstrapTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEBootstrapTemplateSpec.
func (in *RKEBootstrapTemplateSpec) DeepCopy() *RKEBootstrapTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(RKEBootstrapTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKECluster) DeepCopyInto(out *RKECluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKECluster.
func (in *RKECluster) DeepCopy() *RKECluster {
	if in == nil {
		return nil
	}
	out := new(RKECluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKECluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterList) DeepCopyInto(out *RKEClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKECluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterList.
func (in *RKEClusterList) DeepCopy() *RKEClusterList {
	if in == nil {
		return nil
	}
	out := new(RKEClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterSpec) DeepCopyInto(out *RKEClusterSpec) {
	*out = *in
	if in.ControlPlaneEndpoint != nil {
		in, out := &in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint
		*out = new(Endpoint)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterSpec.
func (in *RKEClusterSpec) DeepCopy() *RKEClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RKEClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterSpecCommon) DeepCopyInto(out *RKEClusterSpecCommon) {
	*out = *in
	in.UpgradeStrategy.DeepCopyInto(&out.UpgradeStrategy)
	in.ChartValues.DeepCopyInto(&out.ChartValues)
	in.MachineGlobalConfig.DeepCopyInto(&out.MachineGlobalConfig)
	if in.MachineSelectorConfig != nil {
		in, out := &in.MachineSelectorConfig, &out.MachineSelectorConfig
		*out = make([]RKESystemConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Registries != nil {
		in, out := &in.Registries, &out.Registries
		*out = new(Registry)
		(*in).DeepCopyInto(*out)
	}
	if in.ETCD != nil {
		in, out := &in.ETCD, &out.ETCD
		*out = new(ETCD)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterSpecCommon.
func (in *RKEClusterSpecCommon) DeepCopy() *RKEClusterSpecCommon {
	if in == nil {
		return nil
	}
	out := new(RKEClusterSpecCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEClusterStatus) DeepCopyInto(out *RKEClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEClusterStatus.
func (in *RKEClusterStatus) DeepCopy() *RKEClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RKEClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKECommonNodeConfig) DeepCopyInto(out *RKECommonNodeConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKECommonNodeConfig.
func (in *RKECommonNodeConfig) DeepCopy() *RKECommonNodeConfig {
	if in == nil {
		return nil
	}
	out := new(RKECommonNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlane) DeepCopyInto(out *RKEControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlane.
func (in *RKEControlPlane) DeepCopy() *RKEControlPlane {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlaneList) DeepCopyInto(out *RKEControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RKEControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlaneList.
func (in *RKEControlPlaneList) DeepCopy() *RKEControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RKEControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlaneSpec) DeepCopyInto(out *RKEControlPlaneSpec) {
	*out = *in
	in.RKEClusterSpecCommon.DeepCopyInto(&out.RKEClusterSpecCommon)
	if in.AgentEnvVars != nil {
		in, out := &in.AgentEnvVars, &out.AgentEnvVars
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	out.LocalClusterAuthEndpoint = in.LocalClusterAuthEndpoint
	if in.ETCDSnapshotCreate != nil {
		in, out := &in.ETCDSnapshotCreate, &out.ETCDSnapshotCreate
		*out = new(ETCDSnapshotCreate)
		(*in).DeepCopyInto(*out)
	}
	if in.ETCDSnapshotRestore != nil {
		in, out := &in.ETCDSnapshotRestore, &out.ETCDSnapshotRestore
		*out = new(ETCDSnapshotRestore)
		(*in).DeepCopyInto(*out)
	}
	if in.RotateCertificates != nil {
		in, out := &in.RotateCertificates, &out.RotateCertificates
		*out = new(RotateCertificates)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlaneSpec.
func (in *RKEControlPlaneSpec) DeepCopy() *RKEControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEControlPlaneStatus) DeepCopyInto(out *RKEControlPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	if in.ETCDSnapshotRestore != nil {
		in, out := &in.ETCDSnapshotRestore, &out.ETCDSnapshotRestore
		*out = new(ETCDSnapshotRestore)
		(*in).DeepCopyInto(*out)
	}
	if in.ETCDSnapshotCreate != nil {
		in, out := &in.ETCDSnapshotCreate, &out.ETCDSnapshotCreate
		*out = new(ETCDSnapshotCreate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEControlPlaneStatus.
func (in *RKEControlPlaneStatus) DeepCopy() *RKEControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(RKEControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEMachineStatus) DeepCopyInto(out *RKEMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1beta1.MachineAddress, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEMachineStatus.
func (in *RKEMachineStatus) DeepCopy() *RKEMachineStatus {
	if in == nil {
		return nil
	}
	out := new(RKEMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKESystemConfig) DeepCopyInto(out *RKESystemConfig) {
	*out = *in
	if in.MachineLabelSelector != nil {
		in, out := &in.MachineLabelSelector, &out.MachineLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKESystemConfig.
func (in *RKESystemConfig) DeepCopy() *RKESystemConfig {
	if in == nil {
		return nil
	}
	out := new(RKESystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.Mirrors != nil {
		in, out := &in.Mirrors, &out.Mirrors
		*out = make(map[string]Mirror, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make(map[string]RegistryConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfig) DeepCopyInto(out *RegistryConfig) {
	*out = *in
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryConfig.
func (in *RegistryConfig) DeepCopy() *RegistryConfig {
	if in == nil {
		return nil
	}
	out := new(RegistryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotateCertificates) DeepCopyInto(out *RotateCertificates) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotateCertificates.
func (in *RotateCertificates) DeepCopy() *RotateCertificates {
	if in == nil {
		return nil
	}
	out := new(RotateCertificates)
	in.DeepCopyInto(out)
	return out
}
