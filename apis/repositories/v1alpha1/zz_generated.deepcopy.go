// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionState) DeepCopyInto(out *ConnectionState) {
	*out = *in
	if in.ModifiedAt != nil {
		in, out := &in.ModifiedAt, &out.ModifiedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionState.
func (in *ConnectionState) DeepCopy() *ConnectionState {
	if in == nil {
		return nil
	}
	out := new(ConnectionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryObservation) DeepCopyInto(out *RepositoryObservation) {
	*out = *in
	in.ConnectionState.DeepCopyInto(&out.ConnectionState)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryObservation.
func (in *RepositoryObservation) DeepCopy() *RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryParameters) DeepCopyInto(out *RepositoryParameters) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.PasswordRef != nil {
		in, out := &in.PasswordRef, &out.PasswordRef
		*out = new(SecretReference)
		**out = **in
	}
	if in.SSHPrivateKeyRef != nil {
		in, out := &in.SSHPrivateKeyRef, &out.SSHPrivateKeyRef
		*out = new(SecretReference)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.EnableLFS != nil {
		in, out := &in.EnableLFS, &out.EnableLFS
		*out = new(bool)
		**out = **in
	}
	if in.TLSClientCertDataRef != nil {
		in, out := &in.TLSClientCertDataRef, &out.TLSClientCertDataRef
		*out = new(SecretReference)
		**out = **in
	}
	if in.TLSClientCertKeyRef != nil {
		in, out := &in.TLSClientCertKeyRef, &out.TLSClientCertKeyRef
		*out = new(SecretReference)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.InheritedCreds != nil {
		in, out := &in.InheritedCreds, &out.InheritedCreds
		*out = new(bool)
		**out = **in
	}
	if in.EnableOCI != nil {
		in, out := &in.EnableOCI, &out.EnableOCI
		*out = new(bool)
		**out = **in
	}
	if in.GithubAppPrivateKeyRef != nil {
		in, out := &in.GithubAppPrivateKeyRef, &out.GithubAppPrivateKeyRef
		*out = new(SecretReference)
		**out = **in
	}
	if in.GithubAppID != nil {
		in, out := &in.GithubAppID, &out.GithubAppID
		*out = new(int64)
		**out = **in
	}
	if in.GithubAppInstallationID != nil {
		in, out := &in.GithubAppInstallationID, &out.GithubAppInstallationID
		*out = new(int64)
		**out = **in
	}
	if in.GitHubAppEnterpriseBaseURL != nil {
		in, out := &in.GitHubAppEnterpriseBaseURL, &out.GitHubAppEnterpriseBaseURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryParameters.
func (in *RepositoryParameters) DeepCopy() *RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}