// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateServiceBrokerConfig) DeepCopyInto(out *TemplateServiceBrokerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.TemplateNamespaces != nil {
		in, out := &in.TemplateNamespaces, &out.TemplateNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateServiceBrokerConfig.
func (in *TemplateServiceBrokerConfig) DeepCopy() *TemplateServiceBrokerConfig {
	if in == nil {
		return nil
	}
	out := new(TemplateServiceBrokerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateServiceBrokerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}