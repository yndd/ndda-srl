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

var _ IFSrlNetworkinstanceProtocolsBgpList = &SrlNetworkinstanceProtocolsBgpList{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceProtocolsBgpList interface {
	client.ObjectList

	GetNetworkinstanceProtocolsBgps() []IFSrlNetworkinstanceProtocolsBgp
}

func (x *SrlNetworkinstanceProtocolsBgpList) GetNetworkinstanceProtocolsBgps() []IFSrlNetworkinstanceProtocolsBgp {
	xs := make([]IFSrlNetworkinstanceProtocolsBgp, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFSrlNetworkinstanceProtocolsBgp = &SrlNetworkinstanceProtocolsBgp{}

// +k8s:deepcopy-gen=false
type IFSrlNetworkinstanceProtocolsBgp interface {
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
	GetBgpAdminState() E_NetworkinstanceProtocolsBgpAdminstate
	GetBgpAsPathOptions() []*NetworkinstanceProtocolsBgpAspathoptions
}

// GetCondition
func (x *SrlNetworkinstanceProtocolsBgp) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *SrlNetworkinstanceProtocolsBgp) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *SrlNetworkinstanceProtocolsBgp) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsBgp) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsBgp) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsBgp) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsBgp) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsBgp) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *SrlNetworkinstanceProtocolsBgp) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpAdminState() E_NetworkinstanceProtocolsBgpAdminstate {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Adminstate).IsZero() {
		return ""
	}
	return E_NetworkinstanceProtocolsBgpAdminstate(*x.Spec.NetworkinstanceProtocolsBgp.Adminstate)
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpAsPathOptions() []*NetworkinstanceProtocolsBgpAspathoptions {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Aspathoptions).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Aspathoptions
}
