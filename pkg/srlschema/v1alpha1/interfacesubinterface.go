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
	"strconv"
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
	errCreateInterfaceSubinterface = "cannot create InterfaceSubinterface"
	errDeleteInterfaceSubinterface = "cannot delete InterfaceSubinterface"
	errGetInterfaceSubinterface    = "cannot get InterfaceSubinterface"
)

type InterfaceSubinterface interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.InterfaceSubinterface
	Update(x *srlv1alpha1.InterfaceSubinterface)
	AddInterfaceSubinterfaceAcl(ai *srlv1alpha1.InterfaceSubinterfaceAcl)
	AddInterfaceSubinterfaceAnycastgw(ai *srlv1alpha1.InterfaceSubinterfaceAnycastgw)
	AddInterfaceSubinterfaceBridgetable(ai *srlv1alpha1.InterfaceSubinterfaceBridgetable)
	AddInterfaceSubinterfaceIpv4(ai *srlv1alpha1.InterfaceSubinterfaceIpv4)
	AddInterfaceSubinterfaceIpv6(ai *srlv1alpha1.InterfaceSubinterfaceIpv6)
	AddInterfaceSubinterfaceLocalmirrordestination(ai *srlv1alpha1.InterfaceSubinterfaceLocalmirrordestination)
	AddInterfaceSubinterfaceQos(ai *srlv1alpha1.InterfaceSubinterfaceQos)
	AddInterfaceSubinterfaceRaguard(ai *srlv1alpha1.InterfaceSubinterfaceRaguard)
	AddInterfaceSubinterfaceVlan(ai *srlv1alpha1.InterfaceSubinterfaceVlan)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewInterfaceSubinterface(c resource.ClientApplicator, p Interface, key string) InterfaceSubinterface {
	newInterfaceSubinterfaceList := func() srlv1alpha1.IFSrlInterfaceSubinterfaceList { return &srlv1alpha1.SrlInterfaceSubinterfaceList{} }
	return &interfacesubinterface{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//InterfaceSubinterface: &srlv1alpha1.InterfaceSubinterface{
		//	Name: &name,
		//},
		newInterfaceSubinterfaceList: newInterfaceSubinterfaceList,
	}
}

type interfacesubinterface struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Interface
	// children
	// Data
	InterfaceSubinterface        *srlv1alpha1.InterfaceSubinterface
	newInterfaceSubinterfaceList func() srlv1alpha1.IFSrlInterfaceSubinterfaceList
}

// key type/method

type InterfaceSubinterfaceKey struct {
	Index string
}

func WithInterfaceSubinterfaceKey(key *InterfaceSubinterfaceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
// Data methods
func (x *interfacesubinterface) Update(d *srlv1alpha1.InterfaceSubinterface) {
	x.InterfaceSubinterface = d
}

// methods data
func (x *interfacesubinterface) Get() *srlv1alpha1.InterfaceSubinterface {
	return x.InterfaceSubinterface
}

func (x *interfacesubinterface) GetKey() []string {
	return strings.Split(x.key, ".")
}

// InterfaceSubinterface acl subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceAcl(ai *srlv1alpha1.InterfaceSubinterfaceAcl) {
	x.InterfaceSubinterface.Acl = append(x.InterfaceSubinterface.Acl, ai)
}

// InterfaceSubinterface anycast-gw subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceAnycastgw(ai *srlv1alpha1.InterfaceSubinterfaceAnycastgw) {
	x.InterfaceSubinterface.Anycastgw = append(x.InterfaceSubinterface.Anycastgw, ai)
}

// InterfaceSubinterface bridge-table subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceBridgetable(ai *srlv1alpha1.InterfaceSubinterfaceBridgetable) {
	x.InterfaceSubinterface.Bridgetable = append(x.InterfaceSubinterface.Bridgetable, ai)
}

// InterfaceSubinterface ipv4 subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceIpv4(ai *srlv1alpha1.InterfaceSubinterfaceIpv4) {
	x.InterfaceSubinterface.Ipv4 = append(x.InterfaceSubinterface.Ipv4, ai)
}

// InterfaceSubinterface ipv6 subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceIpv6(ai *srlv1alpha1.InterfaceSubinterfaceIpv6) {
	x.InterfaceSubinterface.Ipv6 = append(x.InterfaceSubinterface.Ipv6, ai)
}

// InterfaceSubinterface local-mirror-destination subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceLocalmirrordestination(ai *srlv1alpha1.InterfaceSubinterfaceLocalmirrordestination) {
	x.InterfaceSubinterface.Localmirrordestination = append(x.InterfaceSubinterface.Localmirrordestination, ai)
}

// InterfaceSubinterface qos subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceQos(ai *srlv1alpha1.InterfaceSubinterfaceQos) {
	x.InterfaceSubinterface.Qos = append(x.InterfaceSubinterface.Qos, ai)
}

// InterfaceSubinterface ra-guard subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceRaguard(ai *srlv1alpha1.InterfaceSubinterfaceRaguard) {
	x.InterfaceSubinterface.Raguard = append(x.InterfaceSubinterface.Raguard, ai)
}

// InterfaceSubinterface vlan subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceVlan(ai *srlv1alpha1.InterfaceSubinterfaceVlan) {
	x.InterfaceSubinterface.Vlan = append(x.InterfaceSubinterface.Vlan, ai)
}

// methods schema

func (x *interfacesubinterface) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.InterfaceSubinterface)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s InterfaceSubinterface: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s InterfaceSubinterface: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *interfacesubinterface) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateInterfaceSubinterface)
		}
	}

	return nil
}
func (x *interfacesubinterface) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlInterfaceSubinterface {
	parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
	//1
	key0 := strconv.Itoa(int(*x.InterfaceSubinterface.Index))

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			// 0
			strings.ToLower(parent0Key0),
			//1
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlInterfaceSubinterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.InterfaceSubinterfaceSpec{
			InterfaceName: &parent0Key0,
			//1
			InterfaceSubinterface: x.InterfaceSubinterface,
		},
	}
}

func (x *interfacesubinterface) InitializeDummySchema() {
}

func (x *interfacesubinterface) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newInterfaceSubinterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetInterfaceSubinterfaces() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *interfacesubinterface) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		parent0Key0 := strings.ReplaceAll(x.parent.GetKey()[0], "/", "-")
		//1
		key0 := strconv.Itoa(int(*x.InterfaceSubinterface.Index))

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(parent0Key0),
				//1
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.InterfaceSubinterfaceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *interfacesubinterface) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.InterfaceSubinterfaceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlInterfaceSubinterface{
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

	return nil
}
