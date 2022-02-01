//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package shared

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeMacPool) DeepCopyInto(out *KubeMacPool) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeMacPool.
func (in *KubeMacPool) DeepCopy() *KubeMacPool {
	if in == nil {
		return nil
	}
	out := new(KubeMacPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxBridge) DeepCopyInto(out *LinuxBridge) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxBridge.
func (in *LinuxBridge) DeepCopy() *LinuxBridge {
	if in == nil {
		return nil
	}
	out := new(LinuxBridge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacvtapCni) DeepCopyInto(out *MacvtapCni) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacvtapCni.
func (in *MacvtapCni) DeepCopy() *MacvtapCni {
	if in == nil {
		return nil
	}
	out := new(MacvtapCni)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Multus) DeepCopyInto(out *Multus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Multus.
func (in *Multus) DeepCopy() *Multus {
	if in == nil {
		return nil
	}
	out := new(Multus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NMState) DeepCopyInto(out *NMState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NMState.
func (in *NMState) DeepCopy() *NMState {
	if in == nil {
		return nil
	}
	out := new(NMState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddonsConfig) DeepCopyInto(out *NetworkAddonsConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddonsConfig.
func (in *NetworkAddonsConfig) DeepCopy() *NetworkAddonsConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkAddonsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddonsConfigSpec) DeepCopyInto(out *NetworkAddonsConfigSpec) {
	*out = *in
	if in.Multus != nil {
		in, out := &in.Multus, &out.Multus
		*out = new(Multus)
		**out = **in
	}
	if in.LinuxBridge != nil {
		in, out := &in.LinuxBridge, &out.LinuxBridge
		*out = new(LinuxBridge)
		**out = **in
	}
	if in.Ovs != nil {
		in, out := &in.Ovs, &out.Ovs
		*out = new(Ovs)
		**out = **in
	}
	if in.KubeMacPool != nil {
		in, out := &in.KubeMacPool, &out.KubeMacPool
		*out = new(KubeMacPool)
		**out = **in
	}
	if in.NMState != nil {
		in, out := &in.NMState, &out.NMState
		*out = new(NMState)
		**out = **in
	}
	if in.MacvtapCni != nil {
		in, out := &in.MacvtapCni, &out.MacvtapCni
		*out = new(MacvtapCni)
		**out = **in
	}
	if in.SelfSignConfiguration != nil {
		in, out := &in.SelfSignConfiguration, &out.SelfSignConfiguration
		*out = new(SelfSignConfiguration)
		**out = **in
	}
	if in.PlacementConfiguration != nil {
		in, out := &in.PlacementConfiguration, &out.PlacementConfiguration
		*out = new(PlacementConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddonsConfigSpec.
func (in *NetworkAddonsConfigSpec) DeepCopy() *NetworkAddonsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkAddonsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddonsConfigStatus) DeepCopyInto(out *NetworkAddonsConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]Container, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddonsConfigStatus.
func (in *NetworkAddonsConfigStatus) DeepCopy() *NetworkAddonsConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkAddonsConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ovs) DeepCopyInto(out *Ovs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ovs.
func (in *Ovs) DeepCopy() *Ovs {
	if in == nil {
		return nil
	}
	out := new(Ovs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Affinity.DeepCopyInto(&out.Affinity)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementConfiguration) DeepCopyInto(out *PlacementConfiguration) {
	*out = *in
	if in.Infra != nil {
		in, out := &in.Infra, &out.Infra
		*out = new(Placement)
		(*in).DeepCopyInto(*out)
	}
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = new(Placement)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementConfiguration.
func (in *PlacementConfiguration) DeepCopy() *PlacementConfiguration {
	if in == nil {
		return nil
	}
	out := new(PlacementConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignConfiguration) DeepCopyInto(out *SelfSignConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignConfiguration.
func (in *SelfSignConfiguration) DeepCopy() *SelfSignConfiguration {
	if in == nil {
		return nil
	}
	out := new(SelfSignConfiguration)
	in.DeepCopyInto(out)
	return out
}
