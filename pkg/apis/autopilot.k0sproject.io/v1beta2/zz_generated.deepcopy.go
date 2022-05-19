//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlNode) DeepCopyInto(out *ControlNode) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.TypeMeta = in.TypeMeta
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlNode.
func (in *ControlNode) DeepCopy() *ControlNode {
	if in == nil {
		return nil
	}
	out := new(ControlNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlNodeList) DeepCopyInto(out *ControlNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControlNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlNodeList.
func (in *ControlNodeList) DeepCopy() *ControlNodeList {
	if in == nil {
		return nil
	}
	out := new(ControlNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlNodeStatus) DeepCopyInto(out *ControlNodeStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlNodeStatus.
func (in *ControlNodeStatus) DeepCopy() *ControlNodeStatus {
	if in == nil {
		return nil
	}
	out := new(ControlNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plan) DeepCopyInto(out *Plan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plan.
func (in *Plan) DeepCopy() *Plan {
	if in == nil {
		return nil
	}
	out := new(Plan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommand) DeepCopyInto(out *PlanCommand) {
	*out = *in
	if in.K0sUpdate != nil {
		in, out := &in.K0sUpdate, &out.K0sUpdate
		*out = new(PlanCommandK0sUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.AirgapUpdate != nil {
		in, out := &in.AirgapUpdate, &out.AirgapUpdate
		*out = new(PlanCommandAirgapUpdate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommand.
func (in *PlanCommand) DeepCopy() *PlanCommand {
	if in == nil {
		return nil
	}
	out := new(PlanCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandAirgapUpdate) DeepCopyInto(out *PlanCommandAirgapUpdate) {
	*out = *in
	if in.Platforms != nil {
		in, out := &in.Platforms, &out.Platforms
		*out = make(PlanPlatformResourceURLMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Workers.DeepCopyInto(&out.Workers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandAirgapUpdate.
func (in *PlanCommandAirgapUpdate) DeepCopy() *PlanCommandAirgapUpdate {
	if in == nil {
		return nil
	}
	out := new(PlanCommandAirgapUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandAirgapUpdateStatus) DeepCopyInto(out *PlanCommandAirgapUpdateStatus) {
	*out = *in
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]PlanCommandTargetStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandAirgapUpdateStatus.
func (in *PlanCommandAirgapUpdateStatus) DeepCopy() *PlanCommandAirgapUpdateStatus {
	if in == nil {
		return nil
	}
	out := new(PlanCommandAirgapUpdateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandK0sUpdate) DeepCopyInto(out *PlanCommandK0sUpdate) {
	*out = *in
	if in.Platforms != nil {
		in, out := &in.Platforms, &out.Platforms
		*out = make(PlanPlatformResourceURLMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Targets.DeepCopyInto(&out.Targets)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandK0sUpdate.
func (in *PlanCommandK0sUpdate) DeepCopy() *PlanCommandK0sUpdate {
	if in == nil {
		return nil
	}
	out := new(PlanCommandK0sUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandK0sUpdateStatus) DeepCopyInto(out *PlanCommandK0sUpdateStatus) {
	*out = *in
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]PlanCommandTargetStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]PlanCommandTargetStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandK0sUpdateStatus.
func (in *PlanCommandK0sUpdateStatus) DeepCopy() *PlanCommandK0sUpdateStatus {
	if in == nil {
		return nil
	}
	out := new(PlanCommandK0sUpdateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandStatus) DeepCopyInto(out *PlanCommandStatus) {
	*out = *in
	if in.K0sUpdate != nil {
		in, out := &in.K0sUpdate, &out.K0sUpdate
		*out = new(PlanCommandK0sUpdateStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.AirgapUpdate != nil {
		in, out := &in.AirgapUpdate, &out.AirgapUpdate
		*out = new(PlanCommandAirgapUpdateStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandStatus.
func (in *PlanCommandStatus) DeepCopy() *PlanCommandStatus {
	if in == nil {
		return nil
	}
	out := new(PlanCommandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTarget) DeepCopyInto(out *PlanCommandTarget) {
	*out = *in
	in.Discovery.DeepCopyInto(&out.Discovery)
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTarget.
func (in *PlanCommandTarget) DeepCopy() *PlanCommandTarget {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTargetDiscovery) DeepCopyInto(out *PlanCommandTargetDiscovery) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(PlanCommandTargetDiscoveryStatic)
		(*in).DeepCopyInto(*out)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(PlanCommandTargetDiscoverySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTargetDiscovery.
func (in *PlanCommandTargetDiscovery) DeepCopy() *PlanCommandTargetDiscovery {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTargetDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTargetDiscoverySelector) DeepCopyInto(out *PlanCommandTargetDiscoverySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTargetDiscoverySelector.
func (in *PlanCommandTargetDiscoverySelector) DeepCopy() *PlanCommandTargetDiscoverySelector {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTargetDiscoverySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTargetDiscoveryStatic) DeepCopyInto(out *PlanCommandTargetDiscoveryStatic) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTargetDiscoveryStatic.
func (in *PlanCommandTargetDiscoveryStatic) DeepCopy() *PlanCommandTargetDiscoveryStatic {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTargetDiscoveryStatic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTargetLimits) DeepCopyInto(out *PlanCommandTargetLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTargetLimits.
func (in *PlanCommandTargetLimits) DeepCopy() *PlanCommandTargetLimits {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTargetLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTargetStatus) DeepCopyInto(out *PlanCommandTargetStatus) {
	*out = *in
	in.LastUpdatedTimestamp.DeepCopyInto(&out.LastUpdatedTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTargetStatus.
func (in *PlanCommandTargetStatus) DeepCopy() *PlanCommandTargetStatus {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanCommandTargets) DeepCopyInto(out *PlanCommandTargets) {
	*out = *in
	in.Controllers.DeepCopyInto(&out.Controllers)
	in.Workers.DeepCopyInto(&out.Workers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanCommandTargets.
func (in *PlanCommandTargets) DeepCopy() *PlanCommandTargets {
	if in == nil {
		return nil
	}
	out := new(PlanCommandTargets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanList) DeepCopyInto(out *PlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanList.
func (in *PlanList) DeepCopy() *PlanList {
	if in == nil {
		return nil
	}
	out := new(PlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PlanPlatformResourceURLMap) DeepCopyInto(out *PlanPlatformResourceURLMap) {
	{
		in := &in
		*out = make(PlanPlatformResourceURLMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanPlatformResourceURLMap.
func (in PlanPlatformResourceURLMap) DeepCopy() PlanPlatformResourceURLMap {
	if in == nil {
		return nil
	}
	out := new(PlanPlatformResourceURLMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanResourceURL) DeepCopyInto(out *PlanResourceURL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanResourceURL.
func (in *PlanResourceURL) DeepCopy() *PlanResourceURL {
	if in == nil {
		return nil
	}
	out := new(PlanResourceURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanSpec) DeepCopyInto(out *PlanSpec) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]PlanCommand, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanSpec.
func (in *PlanSpec) DeepCopy() *PlanSpec {
	if in == nil {
		return nil
	}
	out := new(PlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanStatus) DeepCopyInto(out *PlanStatus) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]PlanCommandStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanStatus.
func (in *PlanStatus) DeepCopy() *PlanStatus {
	if in == nil {
		return nil
	}
	out := new(PlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateConfig) DeepCopyInto(out *UpdateConfig) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateConfig.
func (in *UpdateConfig) DeepCopy() *UpdateConfig {
	if in == nil {
		return nil
	}
	out := new(UpdateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpdateConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateConfigList) DeepCopyInto(out *UpdateConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpdateConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateConfigList.
func (in *UpdateConfigList) DeepCopy() *UpdateConfigList {
	if in == nil {
		return nil
	}
	out := new(UpdateConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpdateConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateSpec) DeepCopyInto(out *UpdateSpec) {
	*out = *in
	out.UpgradeStrategy = in.UpgradeStrategy
	in.PlanSpec.DeepCopyInto(&out.PlanSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateSpec.
func (in *UpdateSpec) DeepCopy() *UpdateSpec {
	if in == nil {
		return nil
	}
	out := new(UpdateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeStrategy) DeepCopyInto(out *UpgradeStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeStrategy.
func (in *UpgradeStrategy) DeepCopy() *UpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(UpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}
