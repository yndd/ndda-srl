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
	errCreateRoutingpolicyAspathset = "cannot create RoutingpolicyAspathset"
	errDeleteRoutingpolicyAspathset = "cannot delete RoutingpolicyAspathset"
	errGetRoutingpolicyAspathset    = "cannot get RoutingpolicyAspathset"
)

type RoutingpolicyAspathset interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *srlv1alpha1.RoutingpolicyAspathset
	Update(x *srlv1alpha1.RoutingpolicyAspathset)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewRoutingpolicyAspathset(c resource.ClientApplicator, p Device, key string) RoutingpolicyAspathset {
	newRoutingpolicyAspathsetList := func() srlv1alpha1.IFSrlRoutingpolicyAspathsetList {
		return &srlv1alpha1.SrlRoutingpolicyAspathsetList{}
	}
	return &routingpolicyaspathset{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//RoutingpolicyAspathset: &srlv1alpha1.RoutingpolicyAspathset{
		//	Name: &name,
		//},
		newRoutingpolicyAspathsetList: newRoutingpolicyAspathsetList,
	}
}

type routingpolicyaspathset struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	RoutingpolicyAspathset        *srlv1alpha1.RoutingpolicyAspathset
	newRoutingpolicyAspathsetList func() srlv1alpha1.IFSrlRoutingpolicyAspathsetList
}

// key type/method

type RoutingpolicyAspathsetKey struct {
	Name string
}

func WithRoutingpolicyAspathsetKey(key *RoutingpolicyAspathsetKey) string {
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
func (x *routingpolicyaspathset) Update(d *srlv1alpha1.RoutingpolicyAspathset) {
	x.RoutingpolicyAspathset = d
}

// methods data
func (x *routingpolicyaspathset) Get() *srlv1alpha1.RoutingpolicyAspathset {
	return x.RoutingpolicyAspathset
}

func (x *routingpolicyaspathset) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *routingpolicyaspathset) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.RoutingpolicyAspathset)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s RoutingpolicyAspathset: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s RoutingpolicyAspathset: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *routingpolicyaspathset) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateRoutingpolicyAspathset)
		}
	}

	return nil
}
func (x *routingpolicyaspathset) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *srlv1alpha1.SrlRoutingpolicyAspathset {
	key0 := strings.ReplaceAll(*x.RoutingpolicyAspathset.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[srlv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[srlv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[srlv1alpha1.LabelNddaDevice] = deviceName
	//labels[srlv1alpha1.LabelNddaItfce] = itfceName
	return &srlv1alpha1.SrlRoutingpolicyAspathset{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: srlv1alpha1.RoutingpolicyAspathsetSpec{
			RoutingpolicyAspathset: x.RoutingpolicyAspathset,
		},
	}
}

func (x *routingpolicyaspathset) InitializeDummySchema() {
}

func (x *routingpolicyaspathset) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{srlv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newRoutingpolicyAspathsetList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetRoutingpolicyAspathsets() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *routingpolicyaspathset) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.RoutingpolicyAspathset.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[srlv1alpha1.RoutingpolicyAspathsetKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *routingpolicyaspathset) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[srlv1alpha1.RoutingpolicyAspathsetKindKind]; ok {
		for resName := range res {
			o := &srlv1alpha1.SrlRoutingpolicyAspathset{
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
