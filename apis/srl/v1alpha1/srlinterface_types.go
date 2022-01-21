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

// Interface struct
type Interface struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	// +kubebuilder:default:="enable"
	Adminstate E_InterfaceAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
	// InterfaceBreakoutmode
	Breakoutmode []*InterfaceBreakoutmode `json:"breakout-mode,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Description *string `json:"description,omitempty"`
	// InterfaceEthernet
	Ethernet []*InterfaceEthernet `json:"ethernet,omitempty"`
	// InterfaceLag
	Lag          []*InterfaceLag `json:"lag,omitempty"`
	Loopbackmode *bool           `json:"loopback-mode,omitempty"`
	// kubebuilder:validation:Minimum=1500
	// kubebuilder:validation:Maximum=9500
	Mtu *uint16 `json:"mtu,omitempty"`
	// kubebuilder:validation:MinLength=3
	// kubebuilder:validation:MaxLength=20
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(mgmt0|mgmt0-standby|system0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])|lag(([1-9](\d){0,1})|(1[0-1]\d)|(12[0-8])))`
	Name *string `json:"name"`
	// InterfaceQos
	Qos []*InterfaceQos `json:"qos,omitempty"`
	// InterfaceSflow
	Sflow []*InterfaceSflow `json:"sflow,omitempty"`
	// InterfaceTransceiver
	Transceiver []*InterfaceTransceiver `json:"transceiver,omitempty"`
	Vlantagging *bool                   `json:"vlan-tagging,omitempty"`
}

// InterfaceBreakoutmode struct
type InterfaceBreakoutmode struct {
	// +kubebuilder:validation:Enum=`10G`;`25G`
	Channelspeed E_InterfaceBreakoutmodeChannelspeed `json:"channel-speed,omitempty"`
	// +kubebuilder:validation:Enum=`4`
	Numchannels E_InterfaceBreakoutmodeNumchannels `json:"num-channels,omitempty"`
}

// InterfaceEthernet struct
type InterfaceEthernet struct {
	Aggregateid   *string `json:"aggregate-id,omitempty"`
	Autonegotiate *bool   `json:"auto-negotiate,omitempty"`
	// +kubebuilder:validation:Enum=`full`;`half`
	Duplexmode E_InterfaceEthernetDuplexmode `json:"duplex-mode,omitempty"`
	//Duplexmode *string `json:"duplex-mode,omitempty"`
	// InterfaceEthernetFlowcontrol
	Flowcontrol []*InterfaceEthernetFlowcontrol `json:"flow-control,omitempty"`
	// InterfaceEthernetHoldtime
	Holdtime []*InterfaceEthernetHoldtime `json:"hold-time,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Lacpportpriority *uint16 `json:"lacp-port-priority,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`1T`;`200G`;`25G`;`400G`;`40G`;`50G`
	Portspeed E_InterfaceEthernetPortspeed `json:"port-speed,omitempty"`
	//Portspeed *string `json:"port-speed,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=86400
	Reloaddelay *uint32 `json:"reload-delay,omitempty"`
	// +kubebuilder:validation:Enum=`lacp`;`power-off`
	Standbysignaling E_InterfaceEthernetStandbysignaling `json:"standby-signaling,omitempty"`
	//Standbysignaling *string `json:"standby-signaling,omitempty"`
	// InterfaceEthernetStormcontrol
	Stormcontrol []*InterfaceEthernetStormcontrol `json:"storm-control,omitempty"`
}

// InterfaceEthernetFlowcontrol struct
type InterfaceEthernetFlowcontrol struct {
	Receive  *bool `json:"receive,omitempty"`
	Transmit *bool `json:"transmit,omitempty"`
}

// InterfaceEthernetHoldtime struct
type InterfaceEthernetHoldtime struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=86400
	Down *uint32 `json:"down,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=86400
	Up *uint32 `json:"up,omitempty"`
}

// InterfaceEthernetStormcontrol struct
type InterfaceEthernetStormcontrol struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	Broadcastrate *uint32 `json:"broadcast-rate,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	Multicastrate *uint32 `json:"multicast-rate,omitempty"`
	// +kubebuilder:validation:Enum=`kbps`;`percentage`
	// +kubebuilder:default:="percentage"
	Units E_InterfaceEthernetStormcontrolUnits `json:"units,omitempty"`
	//Units *string `json:"units,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=100000000
	Unknownunicastrate *uint32 `json:"unknown-unicast-rate,omitempty"`
}

// InterfaceLag struct
type InterfaceLag struct {
	// InterfaceLagLacp
	Lacp []*InterfaceLagLacp `json:"lacp,omitempty"`
	// +kubebuilder:validation:Enum=`static`
	Lacpfallbackmode E_InterfaceLagLacpfallbackmode `json:"lacp-fallback-mode,omitempty"`
	//Lacpfallbackmode *string `json:"lacp-fallback-mode,omitempty"`
	// kubebuilder:validation:Minimum=4
	// kubebuilder:validation:Maximum=3600
	Lacpfallbacktimeout *uint16 `json:"lacp-fallback-timeout,omitempty"`
	// +kubebuilder:validation:Enum=`lacp`;`static`
	// +kubebuilder:default:="static"
	Lagtype E_InterfaceLagLagtype `json:"lag-type,omitempty"`
	//Lagtype *string `json:"lag-type,omitempty"`
	// +kubebuilder:validation:Enum=`100G`;`100M`;`10G`;`10M`;`1G`;`25G`;`400G`;`40G`
	Memberspeed E_InterfaceLagMemberspeed `json:"member-speed,omitempty"`
	//Memberspeed *string `json:"member-speed,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=64
	// +kubebuilder:default:=1
	Minlinks *uint16 `json:"min-links,omitempty"`
}

// InterfaceLagLacp struct
type InterfaceLagLacp struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=65535
	Adminkey *uint16 `json:"admin-key,omitempty"`
	// +kubebuilder:validation:Enum=`FAST`;`SLOW`
	// +kubebuilder:default:="SLOW"
	Interval E_InterfaceLagLacpInterval `json:"interval,omitempty"`
	//Interval *string `json:"interval,omitempty"`
	// +kubebuilder:validation:Enum=`ACTIVE`;`PASSIVE`
	// +kubebuilder:default:="ACTIVE"
	Lacpmode E_InterfaceLagLacpLacpmode `json:"lacp-mode,omitempty"`
	//Lacpmode *string `json:"lacp-mode,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	Systemidmac *string `json:"system-id-mac,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=65535
	Systempriority *uint16 `json:"system-priority,omitempty"`
}

// InterfaceQos struct
type InterfaceQos struct {
	// InterfaceQosOutput
	Output []*InterfaceQosOutput `json:"output,omitempty"`
}

// InterfaceQosOutput struct
type InterfaceQosOutput struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	// InterfaceQosOutputMulticastqueue
	Multicastqueue []*InterfaceQosOutputMulticastqueue `json:"multicast-queue,omitempty"`
	// InterfaceQosOutputScheduler
	Scheduler []*InterfaceQosOutputScheduler `json:"scheduler,omitempty"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1024
	// InterfaceQosOutputUnicastqueue
	Unicastqueue []*InterfaceQosOutputUnicastqueue `json:"unicast-queue,omitempty"`
}

// InterfaceQosOutputMulticastqueue struct
type InterfaceQosOutputMulticastqueue struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Queueid *uint8 `json:"queue-id"`
	// InterfaceQosOutputMulticastqueueScheduling
	Scheduling []*InterfaceQosOutputMulticastqueueScheduling `json:"scheduling,omitempty"`
	Template   *string                                       `json:"template,omitempty"`
}

// InterfaceQosOutputMulticastqueueScheduling struct
type InterfaceQosOutputMulticastqueueScheduling struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Peakratepercent *uint8 `json:"peak-rate-percent,omitempty"`
}

// InterfaceQosOutputScheduler struct
type InterfaceQosOutputScheduler struct {
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=1
	// InterfaceQosOutputSchedulerTier
	Tier []*InterfaceQosOutputSchedulerTier `json:"tier,omitempty"`
}

// InterfaceQosOutputSchedulerTier struct
type InterfaceQosOutputSchedulerTier struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4
	Level *uint8 `json:"level"`
	//+kubebuilder:validation:MinItems=0
	//+kubebuilder:validation:MaxItems=12
	// InterfaceQosOutputSchedulerTierNode
	Node []*InterfaceQosOutputSchedulerTierNode `json:"node,omitempty"`
}

// InterfaceQosOutputSchedulerTierNode struct
type InterfaceQosOutputSchedulerTierNode struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=11
	Nodenumber     *uint8 `json:"node-number"`
	Strictpriority *bool  `json:"strict-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=127
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
}

// InterfaceQosOutputUnicastqueue struct
type InterfaceQosOutputUnicastqueue struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=7
	Queueid *uint8 `json:"queue-id"`
	// InterfaceQosOutputUnicastqueueScheduling
	Scheduling  []*InterfaceQosOutputUnicastqueueScheduling `json:"scheduling,omitempty"`
	Template    *string                                     `json:"template,omitempty"`
	Voqtemplate *string                                     `json:"voq-template,omitempty"`
}

// InterfaceQosOutputUnicastqueueScheduling struct
type InterfaceQosOutputUnicastqueueScheduling struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=100
	// +kubebuilder:default:=100
	Peakratepercent *uint8 `json:"peak-rate-percent,omitempty"`
	// +kubebuilder:default:=true
	Strictpriority *bool `json:"strict-priority,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=255
	// +kubebuilder:default:=1
	Weight *uint8 `json:"weight,omitempty"`
}

// InterfaceSflow struct
type InterfaceSflow struct {
	// +kubebuilder:validation:Enum=`disable`;`enable`
	Adminstate E_InterfaceSflowAdminstate `json:"admin-state,omitempty"`
	//Adminstate *string `json:"admin-state,omitempty"`
}

// InterfaceTransceiver struct
type InterfaceTransceiver struct {
	Ddmevents *bool `json:"ddm-events,omitempty"`
	// +kubebuilder:validation:Enum=`base-r`;`disabled`;`rs-108`;`rs-528`;`rs-544`
	Forwarderrorcorrection E_InterfaceTransceiverForwarderrorcorrection `json:"forward-error-correction,omitempty"`
	//Forwarderrorcorrection *string `json:"forward-error-correction,omitempty"`
	Txlaser *bool `json:"tx-laser,omitempty"`
}

// A InterfaceSpec defines the desired state of a Interface.
type InterfaceSpec struct {
	Interface *Interface `json:"interface,omitempty"`
}

// A InterfaceStatus represents the observed state of a Interface.
type InterfaceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlInterface is the Schema for the Interface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndda,srl}
type SrlInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlInterfaceList contains a list of Interfaces
type SrlInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlInterface{}, &SrlInterfaceList{})
}

// Interface type metadata.
var (
	InterfaceKindKind         = reflect.TypeOf(SrlInterface{}).Name()
	InterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceKindKind}.String()
	InterfaceKindAPIVersion   = InterfaceKindKind + "." + GroupVersion.String()
	InterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceKindKind)
)
