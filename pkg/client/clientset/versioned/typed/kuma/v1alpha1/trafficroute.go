/*
Copyright 2020 The Flux authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/fluxcd/flagger/pkg/apis/kuma/v1alpha1"
	scheme "github.com/fluxcd/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// TrafficRoutesGetter has a method to return a TrafficRouteInterface.
// A group's client should implement this interface.
type TrafficRoutesGetter interface {
	TrafficRoutes() TrafficRouteInterface
}

// TrafficRouteInterface has methods to work with TrafficRoute resources.
type TrafficRouteInterface interface {
	Create(ctx context.Context, trafficRoute *v1alpha1.TrafficRoute, opts v1.CreateOptions) (*v1alpha1.TrafficRoute, error)
	Update(ctx context.Context, trafficRoute *v1alpha1.TrafficRoute, opts v1.UpdateOptions) (*v1alpha1.TrafficRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TrafficRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TrafficRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TrafficRoute, err error)
	TrafficRouteExpansion
}

// trafficRoutes implements TrafficRouteInterface
type trafficRoutes struct {
	*gentype.ClientWithList[*v1alpha1.TrafficRoute, *v1alpha1.TrafficRouteList]
}

// newTrafficRoutes returns a TrafficRoutes
func newTrafficRoutes(c *KumaV1alpha1Client) *trafficRoutes {
	return &trafficRoutes{
		gentype.NewClientWithList[*v1alpha1.TrafficRoute, *v1alpha1.TrafficRouteList](
			"trafficroutes",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.TrafficRoute { return &v1alpha1.TrafficRoute{} },
			func() *v1alpha1.TrafficRouteList { return &v1alpha1.TrafficRouteList{} }),
	}
}
