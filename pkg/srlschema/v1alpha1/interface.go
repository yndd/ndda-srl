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
	errCreateInterface = "cannot create Interface"
	errDeleteInterface = "cannot delete Interface"
	errGetInterface    = "cannot get Interface"
)

type Interface interface {
	// methods children
	NewInterfaceSubinterface(c resource.ClientApplicator, key string) InterfaceSubinterface
	GetInterfaceSubinterfaces() map[string]InterfaceSubinterface
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.Interface
	Update(x *srlv1alpha1.Interface)
	AddInterfaceBreakoutmode(ai *srlv1alpha1.InterfaceBreakoutmode)
	AddInterfaceEthernet(ai *srlv1alpha1.InterfaceEthernet)
	AddInterfaceLag(ai *srlv1alpha1.InterfaceLag)
	AddInterfaceQos(ai *srlv1alpha1.InterfaceQos)
	AddInterfaceSflow(ai *srlv1alpha1.InterfaceSflow)
	AddInterfaceTransceiver(ai *srlv1alpha1.InterfaceTransceiver)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewInterface(c resource.ClientApplicator, p Device, key string) Interface {
	newInterfaceList := func() srlv1alpha1.IFSrlInterfaceList { return &srlv1alpha1.SrlInterfaceList{} }
	return &itfce{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		InterfaceSubinterface: make(map[string]InterfaceSubinterface),
		// data key
		//Interface: &srlv1alpha1.Interface{
		//	Name: &name,
		//},
		newInterfaceList: newInterfaceList,
	}
}

type itfce struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	InterfaceSubinterface map[string]InterfaceSubinterface
	// Data
	Interface        *srlv1alpha1.Interface
	newInterfaceList func() srlv1alpha1.IFSrlInterfaceList
}

// key type/method

type InterfaceKey struct {
	Name string
}

func WithInterfaceKey(key *InterfaceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
func (x *itfce) NewInterfaceSubinterface(c resource.ClientApplicator, key string) InterfaceSubinterface {
	if _, ok := x.InterfaceSubinterface[key]; !ok {
		x.InterfaceSubinterface[key] = NewInterfaceSubinterface(c, x, key)
	}
	return x.InterfaceSubinterface[key]
}
func (x *itfce) GetInterfaceSubinterfaces() map[string]InterfaceSubinterface {
	return x.InterfaceSubinterface
}

// Data methods
func (x *itfce) Update(d *srlv1alpha1.Interface) {
	x.Interface = d
}

// methods data
func (x *itfce) Get() *srlv1alpha1.Interface {
	return x.Interface
}

func (x *itfce) GetKey() []string {
	return strings.Split(x.key, ".")
}

// Interface breakout-mode interface Interface [interface]
func (x *itfce) AddInterfaceBreakoutmode(ai *srlv1alpha1.InterfaceBreakoutmode) {
	x.Interface.Breakoutmode = append(x.Interface.Breakoutmode, ai)
}

// Interface ethernet interface Interface [interface]
func (x *itfce) AddInterfaceEthernet(ai *srlv1alpha1.InterfaceEthernet) {
	x.Interface.Ethernet = append(x.Interface.Ethernet, ai)
}

// Interface lag interface Interface [interface]
func (x *itfce) AddInterfaceLag(ai *srlv1alpha1.InterfaceLag) {
	x.Interface.Lag = append(x.Interface.Lag, ai)
}

// Interface qos interface Interface [interface]
func (x *itfce) AddInterfaceQos(ai *srlv1alpha1.InterfaceQos) {
	x.Interface.Qos = append(x.Interface.Qos, ai)
}

// Interface sflow interface Interface [interface]
func (x *itfce) AddInterfaceSflow(ai *srlv1alpha1.InterfaceSflow) {
	x.Interface.Sflow = append(x.Interface.Sflow, ai)
}

// Interface transceiver interface Interface [interface]
func (x *itfce) AddInterfaceTransceiver(ai *srlv1alpha1.InterfaceTransceiver) {
	x.Interface.Transceiver = append(x.Interface.Transceiver, ai)
}

// methods schema

func (x *itfce) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.Interface)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s Interface: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s Interface: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetInterfaceSubinterfaces() {
		i.Print(key, n)
	}
}

func (x *itfce) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateInterface)
		}
	}
	for _, r := range x.GetInterfaceSubinterfaces() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}
func (x *itfce) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlInterface {
	key0 := strings.ReplaceAll(*x.Interface.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.InterfaceSpec{
			Interface: x.Interface,
		},
	}
}

func (x *itfce) InitializeDummySchema() {
	c0 := x.NewInterfaceSubinterface(x.client, "dummy")
	c0.InitializeDummySchema()
}

func (x *itfce) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newInterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetInterfaces() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	for _, i := range x.GetInterfaceSubinterfaces() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *itfce) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.Interface.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.InterfaceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	for _, i := range x.GetInterfaceSubinterfaces() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *itfce) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.InterfaceKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlInterface{
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
	for _, i := range x.GetInterfaceSubinterfaces() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}

	return nil
}
