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

package srlschema

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/meta"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	srlv1alpha1 "github.com/yndd/ndda-srl/apis/srl/v1alpha1"
)

const (
	errCreateNetworkinstance = "cannot create Networkinstance"
	errDeleteNetworkinstance = "cannot delete Networkinstance"
	errGetNetworkinstance    = "cannot get Networkinstance"
)

type Networkinstance interface {
	// methods children
	NewNetworkinstanceStaticroutes(c resource.ClientApplicator, key string) NetworkinstanceStaticroutes
	NewNetworkinstanceAggregateroutes(c resource.ClientApplicator, key string) NetworkinstanceAggregateroutes
	NewNetworkinstanceProtocolsBgp(c resource.ClientApplicator, key string) NetworkinstanceProtocolsBgp
	NewNetworkinstanceProtocolsOspf(c resource.ClientApplicator, key string) NetworkinstanceProtocolsOspf
	NewNetworkinstanceProtocolsLinux(c resource.ClientApplicator, key string) NetworkinstanceProtocolsLinux
	NewNetworkinstanceNexthopgroups(c resource.ClientApplicator, key string) NetworkinstanceNexthopgroups
	NewNetworkinstanceProtocolsBgpevpn(c resource.ClientApplicator, key string) NetworkinstanceProtocolsBgpevpn
	NewNetworkinstanceProtocolsBgpvpn(c resource.ClientApplicator, key string) NetworkinstanceProtocolsBgpvpn
	NewNetworkinstanceProtocolsIsis(c resource.ClientApplicator, key string) NetworkinstanceProtocolsIsis
	GetNetworkinstanceStaticroutess() map[string]NetworkinstanceStaticroutes
	GetNetworkinstanceAggregateroutess() map[string]NetworkinstanceAggregateroutes
	GetNetworkinstanceProtocolsBgps() map[string]NetworkinstanceProtocolsBgp
	GetNetworkinstanceProtocolsOspfs() map[string]NetworkinstanceProtocolsOspf
	GetNetworkinstanceProtocolsLinuxs() map[string]NetworkinstanceProtocolsLinux
	GetNetworkinstanceNexthopgroupss() map[string]NetworkinstanceNexthopgroups
	GetNetworkinstanceProtocolsBgpevpns() map[string]NetworkinstanceProtocolsBgpevpn
	GetNetworkinstanceProtocolsBgpvpns() map[string]NetworkinstanceProtocolsBgpvpn
	GetNetworkinstanceProtocolsIsiss() map[string]NetworkinstanceProtocolsIsis
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.Networkinstance
	Update(x *srlv1alpha1.Networkinstance)
	AddNetworkinstanceBridgetable(ai *srlv1alpha1.NetworkinstanceBridgetable)
	AddNetworkinstanceInterface(ai *srlv1alpha1.NetworkinstanceInterface)
	AddNetworkinstanceIpforwarding(ai *srlv1alpha1.NetworkinstanceIpforwarding)
	AddNetworkinstanceIploadbalancing(ai *srlv1alpha1.NetworkinstanceIploadbalancing)
	AddNetworkinstanceMpls(ai *srlv1alpha1.NetworkinstanceMpls)
	AddNetworkinstanceMtu(ai *srlv1alpha1.NetworkinstanceMtu)
	AddNetworkinstanceProtocols(ai *srlv1alpha1.NetworkinstanceProtocols)
	AddNetworkinstanceSegmentrouting(ai *srlv1alpha1.NetworkinstanceSegmentrouting)
	AddNetworkinstanceTepolicies(ai *srlv1alpha1.NetworkinstanceTepolicies)
	AddNetworkinstanceTrafficengineering(ai *srlv1alpha1.NetworkinstanceTrafficengineering)
	AddNetworkinstanceVxlaninterface(ai *srlv1alpha1.NetworkinstanceVxlaninterface)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewNetworkinstance(c resource.ClientApplicator, p Device, key string) Networkinstance {
	newNetworkinstanceList := func() srlv1alpha1.IFSrlNetworkinstanceList { return &srlv1alpha1.SrlNetworkinstanceList{} }
	return &networkinstance{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		NetworkinstanceStaticroutes:     make(map[string]NetworkinstanceStaticroutes),
		NetworkinstanceAggregateroutes:  make(map[string]NetworkinstanceAggregateroutes),
		NetworkinstanceProtocolsBgp:     make(map[string]NetworkinstanceProtocolsBgp),
		NetworkinstanceProtocolsOspf:    make(map[string]NetworkinstanceProtocolsOspf),
		NetworkinstanceProtocolsLinux:   make(map[string]NetworkinstanceProtocolsLinux),
		NetworkinstanceNexthopgroups:    make(map[string]NetworkinstanceNexthopgroups),
		NetworkinstanceProtocolsBgpevpn: make(map[string]NetworkinstanceProtocolsBgpevpn),
		NetworkinstanceProtocolsBgpvpn:  make(map[string]NetworkinstanceProtocolsBgpvpn),
		NetworkinstanceProtocolsIsis:    make(map[string]NetworkinstanceProtocolsIsis),
		// data key
		//Networkinstance: &srlv1alpha1.Networkinstance{
		//	Name: &name,
		//},
		newNetworkinstanceList: newNetworkinstanceList,
	}
}

type networkinstance struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	NetworkinstanceStaticroutes     map[string]NetworkinstanceStaticroutes
	NetworkinstanceAggregateroutes  map[string]NetworkinstanceAggregateroutes
	NetworkinstanceProtocolsBgp     map[string]NetworkinstanceProtocolsBgp
	NetworkinstanceProtocolsOspf    map[string]NetworkinstanceProtocolsOspf
	NetworkinstanceProtocolsLinux   map[string]NetworkinstanceProtocolsLinux
	NetworkinstanceNexthopgroups    map[string]NetworkinstanceNexthopgroups
	NetworkinstanceProtocolsBgpevpn map[string]NetworkinstanceProtocolsBgpevpn
	NetworkinstanceProtocolsBgpvpn  map[string]NetworkinstanceProtocolsBgpvpn
	NetworkinstanceProtocolsIsis    map[string]NetworkinstanceProtocolsIsis
	// Data
	Networkinstance        *srlv1alpha1.Networkinstance
	newNetworkinstanceList func() srlv1alpha1.IFSrlNetworkinstanceList
}

// key type/method

type NetworkinstanceKey struct {
	Name string
}

func WithNetworkinstanceKey(key *NetworkinstanceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
func (x *networkinstance) NewNetworkinstanceStaticroutes(c resource.ClientApplicator, key string) NetworkinstanceStaticroutes {
	if _, ok := x.NetworkinstanceStaticroutes[key]; !ok {
		x.NetworkinstanceStaticroutes[key] = NewNetworkinstanceStaticroutes(c, x, key)
	}
	return x.NetworkinstanceStaticroutes[key]
}
func (x *networkinstance) NewNetworkinstanceAggregateroutes(c resource.ClientApplicator, key string) NetworkinstanceAggregateroutes {
	if _, ok := x.NetworkinstanceAggregateroutes[key]; !ok {
		x.NetworkinstanceAggregateroutes[key] = NewNetworkinstanceAggregateroutes(c, x, key)
	}
	return x.NetworkinstanceAggregateroutes[key]
}
func (x *networkinstance) NewNetworkinstanceProtocolsBgp(c resource.ClientApplicator, key string) NetworkinstanceProtocolsBgp {
	if _, ok := x.NetworkinstanceProtocolsBgp[key]; !ok {
		x.NetworkinstanceProtocolsBgp[key] = NewNetworkinstanceProtocolsBgp(c, x, key)
	}
	return x.NetworkinstanceProtocolsBgp[key]
}
func (x *networkinstance) NewNetworkinstanceProtocolsOspf(c resource.ClientApplicator, key string) NetworkinstanceProtocolsOspf {
	if _, ok := x.NetworkinstanceProtocolsOspf[key]; !ok {
		x.NetworkinstanceProtocolsOspf[key] = NewNetworkinstanceProtocolsOspf(c, x, key)
	}
	return x.NetworkinstanceProtocolsOspf[key]
}
func (x *networkinstance) NewNetworkinstanceProtocolsLinux(c resource.ClientApplicator, key string) NetworkinstanceProtocolsLinux {
	if _, ok := x.NetworkinstanceProtocolsLinux[key]; !ok {
		x.NetworkinstanceProtocolsLinux[key] = NewNetworkinstanceProtocolsLinux(c, x, key)
	}
	return x.NetworkinstanceProtocolsLinux[key]
}
func (x *networkinstance) NewNetworkinstanceNexthopgroups(c resource.ClientApplicator, key string) NetworkinstanceNexthopgroups {
	if _, ok := x.NetworkinstanceNexthopgroups[key]; !ok {
		x.NetworkinstanceNexthopgroups[key] = NewNetworkinstanceNexthopgroups(c, x, key)
	}
	return x.NetworkinstanceNexthopgroups[key]
}
func (x *networkinstance) NewNetworkinstanceProtocolsBgpevpn(c resource.ClientApplicator, key string) NetworkinstanceProtocolsBgpevpn {
	if _, ok := x.NetworkinstanceProtocolsBgpevpn[key]; !ok {
		x.NetworkinstanceProtocolsBgpevpn[key] = NewNetworkinstanceProtocolsBgpevpn(c, x, key)
	}
	return x.NetworkinstanceProtocolsBgpevpn[key]
}
func (x *networkinstance) NewNetworkinstanceProtocolsBgpvpn(c resource.ClientApplicator, key string) NetworkinstanceProtocolsBgpvpn {
	if _, ok := x.NetworkinstanceProtocolsBgpvpn[key]; !ok {
		x.NetworkinstanceProtocolsBgpvpn[key] = NewNetworkinstanceProtocolsBgpvpn(c, x, key)
	}
	return x.NetworkinstanceProtocolsBgpvpn[key]
}
func (x *networkinstance) NewNetworkinstanceProtocolsIsis(c resource.ClientApplicator, key string) NetworkinstanceProtocolsIsis {
	if _, ok := x.NetworkinstanceProtocolsIsis[key]; !ok {
		x.NetworkinstanceProtocolsIsis[key] = NewNetworkinstanceProtocolsIsis(c, x, key)
	}
	return x.NetworkinstanceProtocolsIsis[key]
}
func (x *networkinstance) GetNetworkinstanceStaticroutess() map[string]NetworkinstanceStaticroutes {
	return x.NetworkinstanceStaticroutes
}
func (x *networkinstance) GetNetworkinstanceAggregateroutess() map[string]NetworkinstanceAggregateroutes {
	return x.NetworkinstanceAggregateroutes
}
func (x *networkinstance) GetNetworkinstanceProtocolsBgps() map[string]NetworkinstanceProtocolsBgp {
	return x.NetworkinstanceProtocolsBgp
}
func (x *networkinstance) GetNetworkinstanceProtocolsOspfs() map[string]NetworkinstanceProtocolsOspf {
	return x.NetworkinstanceProtocolsOspf
}
func (x *networkinstance) GetNetworkinstanceProtocolsLinuxs() map[string]NetworkinstanceProtocolsLinux {
	return x.NetworkinstanceProtocolsLinux
}
func (x *networkinstance) GetNetworkinstanceNexthopgroupss() map[string]NetworkinstanceNexthopgroups {
	return x.NetworkinstanceNexthopgroups
}
func (x *networkinstance) GetNetworkinstanceProtocolsBgpevpns() map[string]NetworkinstanceProtocolsBgpevpn {
	return x.NetworkinstanceProtocolsBgpevpn
}
func (x *networkinstance) GetNetworkinstanceProtocolsBgpvpns() map[string]NetworkinstanceProtocolsBgpvpn {
	return x.NetworkinstanceProtocolsBgpvpn
}
func (x *networkinstance) GetNetworkinstanceProtocolsIsiss() map[string]NetworkinstanceProtocolsIsis {
	return x.NetworkinstanceProtocolsIsis
}

// Data methods
func (x *networkinstance) Update(d *srlv1alpha1.Networkinstance) {
	x.Networkinstance = d
}

// methods data
func (x *networkinstance) Get() *srlv1alpha1.Networkinstance {
	return x.Networkinstance
}

func (x *networkinstance) GetKey() []string {
	return strings.Split(x.key, ".")
}

// Networkinstance bridge-table networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceBridgetable(ai *srlv1alpha1.NetworkinstanceBridgetable) {
	x.Networkinstance.Bridgetable = append(x.Networkinstance.Bridgetable, ai)
}

// Networkinstance interface networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceInterface(ai *srlv1alpha1.NetworkinstanceInterface) {
	x.Networkinstance.Interface = append(x.Networkinstance.Interface, ai)
}

// Networkinstance ip-forwarding networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceIpforwarding(ai *srlv1alpha1.NetworkinstanceIpforwarding) {
	x.Networkinstance.Ipforwarding = append(x.Networkinstance.Ipforwarding, ai)
}

// Networkinstance ip-load-balancing networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceIploadbalancing(ai *srlv1alpha1.NetworkinstanceIploadbalancing) {
	x.Networkinstance.Iploadbalancing = append(x.Networkinstance.Iploadbalancing, ai)
}

// Networkinstance mpls networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceMpls(ai *srlv1alpha1.NetworkinstanceMpls) {
	x.Networkinstance.Mpls = append(x.Networkinstance.Mpls, ai)
}

// Networkinstance mtu networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceMtu(ai *srlv1alpha1.NetworkinstanceMtu) {
	x.Networkinstance.Mtu = append(x.Networkinstance.Mtu, ai)
}

// Networkinstance protocols networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceProtocols(ai *srlv1alpha1.NetworkinstanceProtocols) {
	x.Networkinstance.Protocols = append(x.Networkinstance.Protocols, ai)
}

// Networkinstance segment-routing networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceSegmentrouting(ai *srlv1alpha1.NetworkinstanceSegmentrouting) {
	x.Networkinstance.Segmentrouting = append(x.Networkinstance.Segmentrouting, ai)
}

// Networkinstance te-policies networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceTepolicies(ai *srlv1alpha1.NetworkinstanceTepolicies) {
	x.Networkinstance.Tepolicies = append(x.Networkinstance.Tepolicies, ai)
}

// Networkinstance traffic-engineering networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceTrafficengineering(ai *srlv1alpha1.NetworkinstanceTrafficengineering) {
	x.Networkinstance.Trafficengineering = append(x.Networkinstance.Trafficengineering, ai)
}

// Networkinstance vxlan-interface networkinstance Networkinstance [network-instance]
func (x *networkinstance) AddNetworkinstanceVxlaninterface(ai *srlv1alpha1.NetworkinstanceVxlaninterface) {
	x.Networkinstance.Vxlaninterface = append(x.Networkinstance.Vxlaninterface, ai)
}

// methods schema

func (x *networkinstance) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.Networkinstance)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s Networkinstance: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s Networkinstance: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetNetworkinstanceStaticroutess() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceAggregateroutess() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceProtocolsBgps() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceProtocolsOspfs() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceProtocolsLinuxs() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceNexthopgroupss() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceProtocolsBgpevpns() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceProtocolsBgpvpns() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkinstanceProtocolsIsiss() {
		i.Print(key, n)
	}
}

func (x *networkinstance) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkinstance)
		}
	}
	for _, r := range x.GetNetworkinstanceStaticroutess() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceAggregateroutess() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceProtocolsBgps() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceProtocolsOspfs() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceProtocolsLinuxs() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceNexthopgroupss() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceProtocolsBgpevpns() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceProtocolsBgpvpns() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkinstanceProtocolsIsiss() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}
func (x *networkinstance) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlNetworkinstance {
	key0 := strings.ReplaceAll(*x.Networkinstance.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlNetworkinstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.NetworkinstanceSpec{
			Networkinstance: x.Networkinstance,
		},
	}
}

func (x *networkinstance) InitializeDummySchema() {
	c0 := x.NewNetworkinstanceStaticroutes(x.client, "dummy")
	c0.InitializeDummySchema()
	c1 := x.NewNetworkinstanceAggregateroutes(x.client, "dummy")
	c1.InitializeDummySchema()
	c2 := x.NewNetworkinstanceProtocolsBgp(x.client, "dummy")
	c2.InitializeDummySchema()
	c3 := x.NewNetworkinstanceProtocolsOspf(x.client, "dummy")
	c3.InitializeDummySchema()
	c4 := x.NewNetworkinstanceProtocolsLinux(x.client, "dummy")
	c4.InitializeDummySchema()
	c5 := x.NewNetworkinstanceNexthopgroups(x.client, "dummy")
	c5.InitializeDummySchema()
	c6 := x.NewNetworkinstanceProtocolsBgpevpn(x.client, "dummy")
	c6.InitializeDummySchema()
	c7 := x.NewNetworkinstanceProtocolsBgpvpn(x.client, "dummy")
	c7.InitializeDummySchema()
	c8 := x.NewNetworkinstanceProtocolsIsis(x.client, "dummy")
	c8.InitializeDummySchema()
}

func (x *networkinstance) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newNetworkinstanceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkinstances() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	for _, i := range x.GetNetworkinstanceStaticroutess() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceAggregateroutess() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgps() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsOspfs() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsLinuxs() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceNexthopgroupss() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgpevpns() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgpvpns() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsIsiss() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *networkinstance) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.Networkinstance.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.NetworkinstanceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	for _, i := range x.GetNetworkinstanceStaticroutess() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceAggregateroutess() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgps() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsOspfs() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsLinuxs() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceNexthopgroupss() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgpevpns() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgpvpns() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsIsiss() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *networkinstance) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.NetworkinstanceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlNetworkinstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resName,
					Namespace: mg.GetNamespace(),
				},
			}
			if err := x.client.Delete(ctx, o); err != nil {
				return err
			}
		}
	}

	// children
	for _, i := range x.GetNetworkinstanceStaticroutess() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceAggregateroutess() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgps() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsOspfs() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsLinuxs() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceNexthopgroupss() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgpevpns() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsBgpvpns() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkinstanceProtocolsIsiss() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}

	return nil
}
