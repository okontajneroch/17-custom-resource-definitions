//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Starfighter) DeepCopyInto(out *Starfighter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Starfighter.
func (in *Starfighter) DeepCopy() *Starfighter {
	if in == nil {
		return nil
	}
	out := new(Starfighter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Starfighter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarfighterSpec) DeepCopyInto(out *StarfighterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarfighterSpec.
func (in *StarfighterSpec) DeepCopy() *StarfighterSpec {
	if in == nil {
		return nil
	}
	out := new(StarfighterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarfighterStatus) DeepCopyInto(out *StarfighterStatus) {
	*out = *in
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarfighterStatus.
func (in *StarfighterStatus) DeepCopy() *StarfighterStatus {
	if in == nil {
		return nil
	}
	out := new(StarfighterStatus)
	in.DeepCopyInto(out)
	return out
}
