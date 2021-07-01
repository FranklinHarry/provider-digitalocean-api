// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Object) DeepCopyInto(out *Object) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Object.
func (in *Object) DeepCopy() *Object {
	if in == nil {
		return nil
	}
	out := new(Object)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Object) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectList) DeepCopyInto(out *ObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Object, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectList.
func (in *ObjectList) DeepCopy() *ObjectList {
	if in == nil {
		return nil
	}
	out := new(ObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSpec) DeepCopyInto(out *ObjectSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ObjectSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSpec.
func (in *ObjectSpec) DeepCopy() *ObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSpecResource) DeepCopyInto(out *ObjectSpecResource) {
	*out = *in
	if in.Acl != nil {
		in, out := &in.Acl, &out.Acl
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CacheControl != nil {
		in, out := &in.CacheControl, &out.CacheControl
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.ContentBase64 != nil {
		in, out := &in.ContentBase64, &out.ContentBase64
		*out = new(string)
		**out = **in
	}
	if in.ContentDisposition != nil {
		in, out := &in.ContentDisposition, &out.ContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ContentEncoding != nil {
		in, out := &in.ContentEncoding, &out.ContentEncoding
		*out = new(string)
		**out = **in
	}
	if in.ContentLanguage != nil {
		in, out := &in.ContentLanguage, &out.ContentLanguage
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
	if in.WebsiteRedirect != nil {
		in, out := &in.WebsiteRedirect, &out.WebsiteRedirect
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSpecResource.
func (in *ObjectSpecResource) DeepCopy() *ObjectSpecResource {
	if in == nil {
		return nil
	}
	out := new(ObjectSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucket) DeepCopyInto(out *SpacesBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucket.
func (in *SpacesBucket) DeepCopy() *SpacesBucket {
	if in == nil {
		return nil
	}
	out := new(SpacesBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpacesBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketList) DeepCopyInto(out *SpacesBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpacesBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketList.
func (in *SpacesBucketList) DeepCopy() *SpacesBucketList {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpacesBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpec) DeepCopyInto(out *SpacesBucketSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(SpacesBucketSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpec.
func (in *SpacesBucketSpec) DeepCopy() *SpacesBucketSpec {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpecCorsRule) DeepCopyInto(out *SpacesBucketSpecCorsRule) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxAgeSeconds != nil {
		in, out := &in.MaxAgeSeconds, &out.MaxAgeSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpecCorsRule.
func (in *SpacesBucketSpecCorsRule) DeepCopy() *SpacesBucketSpecCorsRule {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpecCorsRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpecLifecycleRule) DeepCopyInto(out *SpacesBucketSpecLifecycleRule) {
	*out = *in
	if in.AbortIncompleteMultipartUploadDays != nil {
		in, out := &in.AbortIncompleteMultipartUploadDays, &out.AbortIncompleteMultipartUploadDays
		*out = new(int64)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = new(SpacesBucketSpecLifecycleRuleExpiration)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NoncurrentVersionExpiration != nil {
		in, out := &in.NoncurrentVersionExpiration, &out.NoncurrentVersionExpiration
		*out = new(SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration)
		(*in).DeepCopyInto(*out)
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpecLifecycleRule.
func (in *SpacesBucketSpecLifecycleRule) DeepCopy() *SpacesBucketSpecLifecycleRule {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpecLifecycleRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpecLifecycleRuleExpiration) DeepCopyInto(out *SpacesBucketSpecLifecycleRuleExpiration) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpecLifecycleRuleExpiration.
func (in *SpacesBucketSpecLifecycleRuleExpiration) DeepCopy() *SpacesBucketSpecLifecycleRuleExpiration {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpecLifecycleRuleExpiration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration) DeepCopyInto(out *SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration.
func (in *SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration) DeepCopy() *SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpecLifecycleRuleNoncurrentVersionExpiration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpecResource) DeepCopyInto(out *SpacesBucketSpecResource) {
	*out = *in
	if in.Acl != nil {
		in, out := &in.Acl, &out.Acl
		*out = new(string)
		**out = **in
	}
	if in.BucketDomainName != nil {
		in, out := &in.BucketDomainName, &out.BucketDomainName
		*out = new(string)
		**out = **in
	}
	if in.CorsRule != nil {
		in, out := &in.CorsRule, &out.CorsRule
		*out = make([]SpacesBucketSpecCorsRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]SpacesBucketSpecLifecycleRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Urn != nil {
		in, out := &in.Urn, &out.Urn
		*out = new(string)
		**out = **in
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = new(SpacesBucketSpecVersioning)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpecResource.
func (in *SpacesBucketSpecResource) DeepCopy() *SpacesBucketSpecResource {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketSpecVersioning) DeepCopyInto(out *SpacesBucketSpecVersioning) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketSpecVersioning.
func (in *SpacesBucketSpecVersioning) DeepCopy() *SpacesBucketSpecVersioning {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketSpecVersioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpacesBucketStatus) DeepCopyInto(out *SpacesBucketStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpacesBucketStatus.
func (in *SpacesBucketStatus) DeepCopy() *SpacesBucketStatus {
	if in == nil {
		return nil
	}
	out := new(SpacesBucketStatus)
	in.DeepCopyInto(out)
	return out
}
