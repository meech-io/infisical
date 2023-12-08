//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
	out.ServiceToken = in.ServiceToken
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecret) DeepCopyInto(out *InfisicalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecret.
func (in *InfisicalSecret) DeepCopy() *InfisicalSecret {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretList) DeepCopyInto(out *InfisicalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfisicalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretList.
func (in *InfisicalSecretList) DeepCopy() *InfisicalSecretList {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretSpec) DeepCopyInto(out *InfisicalSecretSpec) {
	*out = *in
	out.TokenSecretReference = in.TokenSecretReference
	out.Authentication = in.Authentication
	out.ManagedSecretReference = in.ManagedSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretSpec.
func (in *InfisicalSecretSpec) DeepCopy() *InfisicalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretStatus) DeepCopyInto(out *InfisicalSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretStatus.
func (in *InfisicalSecretStatus) DeepCopy() *InfisicalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSecretReference) DeepCopyInto(out *KubeSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSecretReference.
func (in *KubeSecretReference) DeepCopy() *KubeSecretReference {
	if in == nil {
		return nil
	}
	out := new(KubeSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MangedKubeSecretConfig) DeepCopyInto(out *MangedKubeSecretConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MangedKubeSecretConfig.
func (in *MangedKubeSecretConfig) DeepCopy() *MangedKubeSecretConfig {
	if in == nil {
		return nil
	}
	out := new(MangedKubeSecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretScopeInWorkspace) DeepCopyInto(out *SecretScopeInWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretScopeInWorkspace.
func (in *SecretScopeInWorkspace) DeepCopy() *SecretScopeInWorkspace {
	if in == nil {
		return nil
	}
	out := new(SecretScopeInWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountDetails) DeepCopyInto(out *ServiceAccountDetails) {
	*out = *in
	out.ServiceAccountSecretReference = in.ServiceAccountSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountDetails.
func (in *ServiceAccountDetails) DeepCopy() *ServiceAccountDetails {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTokenDetails) DeepCopyInto(out *ServiceTokenDetails) {
	*out = *in
	out.ServiceTokenSecretReference = in.ServiceTokenSecretReference
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTokenDetails.
func (in *ServiceTokenDetails) DeepCopy() *ServiceTokenDetails {
	if in == nil {
		return nil
	}
	out := new(ServiceTokenDetails)
	in.DeepCopyInto(out)
	return out
}
