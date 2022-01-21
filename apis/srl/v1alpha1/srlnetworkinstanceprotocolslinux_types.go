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

// NetworkinstanceProtocolsLinux struct
type NetworkinstanceProtocolsLinux struct {
	// +kubebuilder:default:=true
	Exportneighbors *bool `json:"export-neighbors,omitempty"`
	// +kubebuilder:default:=false
	Exportroutes *bool `json:"export-routes,omitempty"`
	// +kubebuilder:default:=false
	Importroutes *bool `json:"import-routes,omitempty"`
}

// A NetworkinstanceProtocolsLinuxSpec defines the desired state of a NetworkinstanceProtocolsLinux.
type NetworkinstanceProtocolsLinuxSpec struct {
	NetworkInstanceName           *string                        `json:"network-instance-name"`
	NetworkinstanceProtocolsLinux *NetworkinstanceProtocolsLinux `json:"linux,omitempty"`
}

// A NetworkinstanceProtocolsLinuxStatus represents the observed state of a NetworkinstanceProtocolsLinux.
type NetworkinstanceProtocolsLinuxStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsLinux is the Schema for the NetworkinstanceProtocolsLinux API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={ndda,srl}
type SrlNetworkinstanceProtocolsLinux struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkinstanceProtocolsLinuxSpec   `json:"spec,omitempty"`
	Status NetworkinstanceProtocolsLinuxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SrlNetworkinstanceProtocolsLinuxList contains a list of NetworkinstanceProtocolsLinuxs
type SrlNetworkinstanceProtocolsLinuxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SrlNetworkinstanceProtocolsLinux `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SrlNetworkinstanceProtocolsLinux{}, &SrlNetworkinstanceProtocolsLinuxList{})
}

// NetworkinstanceProtocolsLinux type metadata.
var (
	NetworkinstanceProtocolsLinuxKindKind         = reflect.TypeOf(SrlNetworkinstanceProtocolsLinux{}).Name()
	NetworkinstanceProtocolsLinuxGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkinstanceProtocolsLinuxKindKind}.String()
	NetworkinstanceProtocolsLinuxKindAPIVersion   = NetworkinstanceProtocolsLinuxKindKind + "." + GroupVersion.String()
	NetworkinstanceProtocolsLinuxGroupVersionKind = GroupVersion.WithKind(NetworkinstanceProtocolsLinuxKindKind)
)
