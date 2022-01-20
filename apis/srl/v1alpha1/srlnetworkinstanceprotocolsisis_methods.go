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

/*
import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)
*/
// Isis
type E_NetworkinstanceProtocolsIsisInstanceAdminstate string

const (
	E_NetworkinstanceProtocolsIsisInstanceAdminstate_Disable E_NetworkinstanceProtocolsIsisInstanceAdminstate = "disable"
	E_NetworkinstanceProtocolsIsisInstanceAdminstate_Enable  E_NetworkinstanceProtocolsIsisInstanceAdminstate = "enable"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceLevelcapability string

const (
	E_NetworkinstanceProtocolsIsisInstanceLevelcapability_L1   E_NetworkinstanceProtocolsIsisInstanceLevelcapability = "L1"
	E_NetworkinstanceProtocolsIsisInstanceLevelcapability_L1l2 E_NetworkinstanceProtocolsIsisInstanceLevelcapability = "L1L2"
	E_NetworkinstanceProtocolsIsisInstanceLevelcapability_L2   E_NetworkinstanceProtocolsIsisInstanceLevelcapability = "L2"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceInterfaceAdminstate string

const (
	E_NetworkinstanceProtocolsIsisInstanceInterfaceAdminstate_Disable E_NetworkinstanceProtocolsIsisInstanceInterfaceAdminstate = "disable"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceAdminstate_Enable  E_NetworkinstanceProtocolsIsisInstanceInterfaceAdminstate = "enable"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceInterfaceCircuittype string

const (
	E_NetworkinstanceProtocolsIsisInstanceInterfaceCircuittype_Broadcast    E_NetworkinstanceProtocolsIsisInstanceInterfaceCircuittype = "broadcast"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceCircuittype_PointToPoint E_NetworkinstanceProtocolsIsisInstanceInterfaceCircuittype = "point-to-point"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding string

const (
	E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding_Adaptive E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding = "adaptive"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding_Disable  E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding = "disable"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding_Loose    E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding = "loose"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding_Strict   E_NetworkinstanceProtocolsIsisInstanceInterfaceHellopadding = "strict"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate string

const (
	E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate_Disable E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate_Enable  E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv4unicastAdminstate = "enable"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate string

const (
	E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate_Disable E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate_Enable  E_NetworkinstanceProtocolsIsisInstanceInterfaceIpv6unicastAdminstate = "enable"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace string

const (
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_Adjacencies     E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "adjacencies"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsAll      E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-all"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL1Csnp   E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l1-csnp"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL1Hello  E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l1-hello"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL1Lsp    E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l1-lsp"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL1Psnp   E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l1-psnp"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL2Csnp   E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l2-csnp"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL2Hello  E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l2-hello"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL2Lsp    E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l2-lsp"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsL2Psnp   E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-l2-psnp"
	E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace_PacketsP2pHello E_NetworkinstanceProtocolsIsisInstanceInterfaceTraceoptionsTraceTrace = "packets-p2p-hello"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate string

const (
	E_NetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate_Disable E_NetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate_Enable  E_NetworkinstanceProtocolsIsisInstanceIpv4unicastAdminstate = "enable"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate string

const (
	E_NetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate_Disable E_NetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate = "disable"
	E_NetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate_Enable  E_NetworkinstanceProtocolsIsisInstanceIpv6unicastAdminstate = "enable"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceLevelMetricstyle string

const (
	E_NetworkinstanceProtocolsIsisInstanceLevelMetricstyle_Narrow E_NetworkinstanceProtocolsIsisInstanceLevelMetricstyle = "narrow"
	E_NetworkinstanceProtocolsIsisInstanceLevelMetricstyle_Wide   E_NetworkinstanceProtocolsIsisInstanceLevelMetricstyle = "wide"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace string

const (
	E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace_Adjacencies E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace = "adjacencies"
	E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace_Lsdb        E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace = "lsdb"
	E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace_Routes      E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace = "routes"
	E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace_Spf         E_NetworkinstanceProtocolsIsisInstanceLevelTraceoptionsTraceTrace = "spf"
)

// Isis
type E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace string

const (
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_Adjacencies      E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "adjacencies"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_GracefulRestart  E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "graceful-restart"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_Interfaces       E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "interfaces"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsAll       E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-all"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL1Csnp    E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l1-csnp"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL1Hello   E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l1-hello"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL1Lsp     E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l1-lsp"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL1Psnp    E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l1-psnp"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL2Csnp    E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l2-csnp"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL2Hello   E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l2-hello"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL2Lsp     E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l2-lsp"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsL2Psnp    E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-l2-psnp"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_PacketsP2pHello  E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "packets-p2p-hello"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_Routes           E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "routes"
	E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace_SummaryAddresses E_NetworkinstanceProtocolsIsisInstanceTraceoptionsTraceTrace = "summary-addresses"
)

func NewNetworkinstanceProtocolsIsis() *NetworkinstanceProtocolsIsis {
	return &NetworkinstanceProtocolsIsis{}
}