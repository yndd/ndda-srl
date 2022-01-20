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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="disable"
	Adminstate *string `json:"admin-state,omitempty"`
	// EthernetsegmentDfelection
	Dfelection []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelection `json:"df-election,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){9}`
	Esi       *string `json:"esi,omitempty"`
	Interface *string `json:"interface,omitempty"`
	// +kubebuilder:validation:Enum=`all-active`;`single-active`
	// +kubebuilder:default:="all-active"
	Multihomingmode *string `json:"multi-homing-mode,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name"`
	// EthernetsegmentRoutes
	Routes []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes `json:"routes,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelection struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelection struct {
	// EthernetsegmentDfelectionAlgorithm
	Algorithm []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithm `json:"algorithm,omitempty"`
	// EthernetsegmentDfelectionInterfacestandbysignalingonnondf
	Interfacestandbysignalingonnondf []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionInterfacestandbysignalingonnondf `json:"interface-standby-signaling-on-non-df,omitempty"`
	// EthernetsegmentDfelectionTimers
	Timers []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionTimers `json:"timers,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithm struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithm struct {
	// EthernetsegmentDfelectionAlgorithmPreferencealg
	Preferencealg []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealg `json:"preference-alg,omitempty"`
	// +kubebuilder:validation:Enum=`default`;`preference`
	// +kubebuilder:default:="default"
	Type *string `json:"type,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealg struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealg struct {
	// EthernetsegmentDfelectionAlgorithmPreferencealgCapabilities
	Capabilities []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilities `json:"capabilities,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	// +kubebuilder:default:=32767
	Preferencevalue *uint32 `json:"preference-value,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilities struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionAlgorithmPreferencealgCapabilities struct {
	// +kubebuilder:validation:Enum=`exclude`;`include`
	// +kubebuilder:default:="include"
	Acdf *string `json:"ac-df,omitempty"`
	// +kubebuilder:default:=false
	Nonrevertive *bool `json:"non-revertive,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionInterfacestandbysignalingonnondf struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionInterfacestandbysignalingonnondf struct {
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionTimers struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiDfelectionTimers struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100
	Activationtimer *uint32 `json:"activation-timer,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutes struct {
	// EthernetsegmentRoutesEthernetsegment
	Esi []*SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetsegment `json:"esi,omitempty"`
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:="use-system-ipv4-address"
	Nexthop *string `json:"next-hop,omitempty"`
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetsegment struct
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiRoutesEthernetsegment struct {
	// +kubebuilder:validation:Enum=`use-system-ipv4-address`
	// +kubebuilder:default:="use-system-ipv4-address"
	Originatingip *string `json:"originating-ip,omitempty"`
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec defines the desired state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec struct {
	BgpInstanceId                                        *string                                               `json:"bgp-instance-id"`
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi *SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi `json:"ethernet-segment,omitempty"`
}

// A SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus represents the observed state of a SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi.
type SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi is the Schema for the SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiSpec   `json:"spec,omitempty"`
	Status SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList contains a list of SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsis
type SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}, &SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiList{})
}

// SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi type metadata.
var (
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindKind         = reflect.TypeOf(SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}).Name()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiGroupKind        = schema.GroupKind{Group: Group, Kind: SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindKind}.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindAPIVersion   = SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindKind + "." + GroupVersion.String()
	SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiGroupVersionKind = GroupVersion.WithKind(SystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsiKindKind)
)
