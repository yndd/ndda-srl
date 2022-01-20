/*
Copyright 2022 NDDA.

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

package v1alpha1

import (
	"reflect"

	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ IFSrlTunnelinterfaceVxlaninterfaceList = &SrlTunnelinterfaceVxlaninterfaceList{}

// +k8s:deepcopy-gen=false
type IFSrlTunnelinterfaceVxlaninterfaceList interface {
	client.ObjectList

	GetTunnelinterfaceVxlaninterfaces() []IFSrlTunnelinterfaceVxlaninterface
}

func (x *SrlTunnelinterfaceVxlaninterfaceList) GetTunnelinterfaceVxlaninterfaces() []IFSrlTunnelinterfaceVxlaninterface {
	xs := make([]IFSrlTunnelinterfaceVxlaninterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlTunnelinterfaceVxlaninterface = &SrlTunnelinterfaceVxlaninterface{}

// +k8s:deepcopy-gen=false
type IFSrlTunnelinterfaceVxlaninterface interface {
	resource.Object
	resource.Conditioned

	GetCondition(ct nddv1.ConditionKind) nddv1.Condition
	SetConditions(c ...nddv1.Condition)
	// getters based on labels
	GetOwner() string
	GetDeploymentPolicy() string
	GetDeviceName() string
	GetEndpointGroup() string
	GetOrganization() string
	GetDeployment() string
	GetAvailabilityZone() string
	// getters based on type
	GetVxlaninterfaceBridgeTable() []*TunnelinterfaceVxlaninterfaceBridgetable
}

// GetCondition
func (x *SrlTunnelinterfaceVxlaninterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlTunnelinterfaceVxlaninterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlTunnelinterfaceVxlaninterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlTunnelinterfaceVxlaninterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlTunnelinterfaceVxlaninterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlTunnelinterfaceVxlaninterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlTunnelinterfaceVxlaninterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlTunnelinterfaceVxlaninterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlTunnelinterfaceVxlaninterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlTunnelinterfaceVxlaninterface) GetVxlaninterfaceBridgeTable() []*TunnelinterfaceVxlaninterfaceBridgetable {
	if reflect.ValueOf(x.Spec.TunnelinterfaceVxlaninterface.Bridgetable).IsZero() {
		return nil
	}
	return x.Spec.TunnelinterfaceVxlaninterface.Bridgetable
}
