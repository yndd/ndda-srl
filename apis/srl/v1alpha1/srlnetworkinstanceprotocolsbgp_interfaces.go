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
	GetBgpAuthentication() []*NetworkinstanceProtocolsBgpAuthentication
	GetBgpAutonomousSystem() uint32
	GetBgpConvergence() []*NetworkinstanceProtocolsBgpConvergence
	GetBgpDynamicNeighbors() []*NetworkinstanceProtocolsBgpDynamicneighbors
	GetBgpEbgpDefaultPolicy() []*NetworkinstanceProtocolsBgpEbgpdefaultpolicy
	GetBgpEvpn() []*NetworkinstanceProtocolsBgpEvpn
	GetBgpExportPolicy() string
	GetBgpFailureDetection() []*NetworkinstanceProtocolsBgpFailuredetection
	GetBgpGracefulRestart() []*NetworkinstanceProtocolsBgpGracefulrestart
	GetBgpGroup() []*NetworkinstanceProtocolsBgpGroup
	GetBgpImportPolicy() string
	GetBgpIpv4Unicast() []*NetworkinstanceProtocolsBgpIpv4unicast
	GetBgpIpv6Unicast() []*NetworkinstanceProtocolsBgpIpv6unicast
	GetBgpLocalPreference() uint32
	GetBgpNeighbor() []*NetworkinstanceProtocolsBgpNeighbor
	GetBgpPreference() []*NetworkinstanceProtocolsBgpPreference
	GetBgpRouteAdvertisement() []*NetworkinstanceProtocolsBgpRouteadvertisement
	GetBgpRouteReflector() []*NetworkinstanceProtocolsBgpRoutereflector
	GetBgpRouterId() string
	GetBgpSendCommunity() []*NetworkinstanceProtocolsBgpSendcommunity
	GetBgpTraceOptions() []*NetworkinstanceProtocolsBgpTraceoptions
	GetBgpTransport() []*NetworkinstanceProtocolsBgpTransport
	// add based on type
	AddBgpAsPathOptions(a *NetworkinstanceProtocolsBgpAspathoptions)
	AddBgpAuthentication(a *NetworkinstanceProtocolsBgpAuthentication)
	AddBgpConvergence(a *NetworkinstanceProtocolsBgpConvergence)
	AddBgpDynamicNeighbors(a *NetworkinstanceProtocolsBgpDynamicneighbors)
	AddBgpEbgpDefaultPolicy(a *NetworkinstanceProtocolsBgpEbgpdefaultpolicy)
	AddBgpEvpn(a *NetworkinstanceProtocolsBgpEvpn)
	AddBgpFailureDetection(a *NetworkinstanceProtocolsBgpFailuredetection)
	AddBgpGracefulRestart(a *NetworkinstanceProtocolsBgpGracefulrestart)
	AddBgpGroup(a *NetworkinstanceProtocolsBgpGroup)
	AddBgpIpv4Unicast(a *NetworkinstanceProtocolsBgpIpv4unicast)
	AddBgpIpv6Unicast(a *NetworkinstanceProtocolsBgpIpv6unicast)
	AddBgpNeighbor(a *NetworkinstanceProtocolsBgpNeighbor)
	AddBgpPreference(a *NetworkinstanceProtocolsBgpPreference)
	AddBgpRouteAdvertisement(a *NetworkinstanceProtocolsBgpRouteadvertisement)
	AddBgpRouteReflector(a *NetworkinstanceProtocolsBgpRoutereflector)
	AddBgpSendCommunity(a *NetworkinstanceProtocolsBgpSendcommunity)
	AddBgpTraceOptions(a *NetworkinstanceProtocolsBgpTraceoptions)
	AddBgpTransport(a *NetworkinstanceProtocolsBgpTransport)
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
	return E_NetworkinstanceProtocolsBgpAdminstate(x.Spec.NetworkinstanceProtocolsBgp.Adminstate)
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpAsPathOptions() []*NetworkinstanceProtocolsBgpAspathoptions {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Aspathoptions).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Aspathoptions
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpAuthentication() []*NetworkinstanceProtocolsBgpAuthentication {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Authentication).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Authentication
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpAutonomousSystem() uint32 {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Autonomoussystem).IsZero() {
		return 0
	}
	return *x.Spec.NetworkinstanceProtocolsBgp.Autonomoussystem
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpConvergence() []*NetworkinstanceProtocolsBgpConvergence {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Convergence).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Convergence
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpDynamicNeighbors() []*NetworkinstanceProtocolsBgpDynamicneighbors {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Dynamicneighbors).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Dynamicneighbors
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpEbgpDefaultPolicy() []*NetworkinstanceProtocolsBgpEbgpdefaultpolicy {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Ebgpdefaultpolicy).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Ebgpdefaultpolicy
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpEvpn() []*NetworkinstanceProtocolsBgpEvpn {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Evpn).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Evpn
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpExportPolicy() string {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Exportpolicy).IsZero() {
		return ""
	}
	return *x.Spec.NetworkinstanceProtocolsBgp.Exportpolicy
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpFailureDetection() []*NetworkinstanceProtocolsBgpFailuredetection {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Failuredetection).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Failuredetection
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpGracefulRestart() []*NetworkinstanceProtocolsBgpGracefulrestart {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Gracefulrestart).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Gracefulrestart
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpGroup() []*NetworkinstanceProtocolsBgpGroup {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Group).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Group
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpImportPolicy() string {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Importpolicy).IsZero() {
		return ""
	}
	return *x.Spec.NetworkinstanceProtocolsBgp.Importpolicy
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpIpv4Unicast() []*NetworkinstanceProtocolsBgpIpv4unicast {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Ipv4unicast).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Ipv4unicast
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpIpv6Unicast() []*NetworkinstanceProtocolsBgpIpv6unicast {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Ipv6unicast).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Ipv6unicast
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpLocalPreference() uint32 {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Localpreference).IsZero() {
		return 0
	}
	return *x.Spec.NetworkinstanceProtocolsBgp.Localpreference
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpNeighbor() []*NetworkinstanceProtocolsBgpNeighbor {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Neighbor).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Neighbor
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpPreference() []*NetworkinstanceProtocolsBgpPreference {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Preference).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Preference
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpRouteAdvertisement() []*NetworkinstanceProtocolsBgpRouteadvertisement {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Routeadvertisement).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Routeadvertisement
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpRouteReflector() []*NetworkinstanceProtocolsBgpRoutereflector {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Routereflector).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Routereflector
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpRouterId() string {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Routerid).IsZero() {
		return ""
	}
	return *x.Spec.NetworkinstanceProtocolsBgp.Routerid
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpSendCommunity() []*NetworkinstanceProtocolsBgpSendcommunity {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Sendcommunity).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Sendcommunity
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpTraceOptions() []*NetworkinstanceProtocolsBgpTraceoptions {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Traceoptions).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Traceoptions
}
func (x *SrlNetworkinstanceProtocolsBgp) GetBgpTransport() []*NetworkinstanceProtocolsBgpTransport {
	if reflect.ValueOf(x.Spec.NetworkinstanceProtocolsBgp.Transport).IsZero() {
		return nil
	}
	return x.Spec.NetworkinstanceProtocolsBgp.Transport
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpAsPathOptions(a *NetworkinstanceProtocolsBgpAspathoptions) {
	x.Spec.NetworkinstanceProtocolsBgp.Aspathoptions = append(x.Spec.NetworkinstanceProtocolsBgp.Aspathoptions, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpAuthentication(a *NetworkinstanceProtocolsBgpAuthentication) {
	x.Spec.NetworkinstanceProtocolsBgp.Authentication = append(x.Spec.NetworkinstanceProtocolsBgp.Authentication, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpConvergence(a *NetworkinstanceProtocolsBgpConvergence) {
	x.Spec.NetworkinstanceProtocolsBgp.Convergence = append(x.Spec.NetworkinstanceProtocolsBgp.Convergence, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpDynamicNeighbors(a *NetworkinstanceProtocolsBgpDynamicneighbors) {
	x.Spec.NetworkinstanceProtocolsBgp.Dynamicneighbors = append(x.Spec.NetworkinstanceProtocolsBgp.Dynamicneighbors, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpEbgpDefaultPolicy(a *NetworkinstanceProtocolsBgpEbgpdefaultpolicy) {
	x.Spec.NetworkinstanceProtocolsBgp.Ebgpdefaultpolicy = append(x.Spec.NetworkinstanceProtocolsBgp.Ebgpdefaultpolicy, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpEvpn(a *NetworkinstanceProtocolsBgpEvpn) {
	x.Spec.NetworkinstanceProtocolsBgp.Evpn = append(x.Spec.NetworkinstanceProtocolsBgp.Evpn, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpFailureDetection(a *NetworkinstanceProtocolsBgpFailuredetection) {
	x.Spec.NetworkinstanceProtocolsBgp.Failuredetection = append(x.Spec.NetworkinstanceProtocolsBgp.Failuredetection, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpGracefulRestart(a *NetworkinstanceProtocolsBgpGracefulrestart) {
	x.Spec.NetworkinstanceProtocolsBgp.Gracefulrestart = append(x.Spec.NetworkinstanceProtocolsBgp.Gracefulrestart, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpGroup(a *NetworkinstanceProtocolsBgpGroup) {
	x.Spec.NetworkinstanceProtocolsBgp.Group = append(x.Spec.NetworkinstanceProtocolsBgp.Group, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpIpv4Unicast(a *NetworkinstanceProtocolsBgpIpv4unicast) {
	x.Spec.NetworkinstanceProtocolsBgp.Ipv4unicast = append(x.Spec.NetworkinstanceProtocolsBgp.Ipv4unicast, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpIpv6Unicast(a *NetworkinstanceProtocolsBgpIpv6unicast) {
	x.Spec.NetworkinstanceProtocolsBgp.Ipv6unicast = append(x.Spec.NetworkinstanceProtocolsBgp.Ipv6unicast, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpNeighbor(a *NetworkinstanceProtocolsBgpNeighbor) {
	x.Spec.NetworkinstanceProtocolsBgp.Neighbor = append(x.Spec.NetworkinstanceProtocolsBgp.Neighbor, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpPreference(a *NetworkinstanceProtocolsBgpPreference) {
	x.Spec.NetworkinstanceProtocolsBgp.Preference = append(x.Spec.NetworkinstanceProtocolsBgp.Preference, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpRouteAdvertisement(a *NetworkinstanceProtocolsBgpRouteadvertisement) {
	x.Spec.NetworkinstanceProtocolsBgp.Routeadvertisement = append(x.Spec.NetworkinstanceProtocolsBgp.Routeadvertisement, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpRouteReflector(a *NetworkinstanceProtocolsBgpRoutereflector) {
	x.Spec.NetworkinstanceProtocolsBgp.Routereflector = append(x.Spec.NetworkinstanceProtocolsBgp.Routereflector, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpSendCommunity(a *NetworkinstanceProtocolsBgpSendcommunity) {
	x.Spec.NetworkinstanceProtocolsBgp.Sendcommunity = append(x.Spec.NetworkinstanceProtocolsBgp.Sendcommunity, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpTraceOptions(a *NetworkinstanceProtocolsBgpTraceoptions) {
	x.Spec.NetworkinstanceProtocolsBgp.Traceoptions = append(x.Spec.NetworkinstanceProtocolsBgp.Traceoptions, a)
}
func (x *SrlNetworkinstanceProtocolsBgp) AddBgpTransport(a *NetworkinstanceProtocolsBgpTransport) {
	x.Spec.NetworkinstanceProtocolsBgp.Transport = append(x.Spec.NetworkinstanceProtocolsBgp.Transport, a)
}
