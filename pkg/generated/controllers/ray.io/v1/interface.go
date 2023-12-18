/*
Copyright 2023 1block.ai.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v2/pkg/generic"
	"github.com/rancher/wrangler/v2/pkg/schemes"
	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	RayCluster() RayClusterController
	RayJob() RayJobController
	RayService() RayServiceController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) RayCluster() RayClusterController {
	return generic.NewController[*v1.RayCluster, *v1.RayClusterList](schema.GroupVersionKind{Group: "ray.io", Version: "v1", Kind: "RayCluster"}, "rayclusters", true, v.controllerFactory)
}

func (v *version) RayJob() RayJobController {
	return generic.NewController[*v1.RayJob, *v1.RayJobList](schema.GroupVersionKind{Group: "ray.io", Version: "v1", Kind: "RayJob"}, "rayjobs", true, v.controllerFactory)
}

func (v *version) RayService() RayServiceController {
	return generic.NewController[*v1.RayService, *v1.RayServiceList](schema.GroupVersionKind{Group: "ray.io", Version: "v1", Kind: "RayService"}, "rayservices", true, v.controllerFactory)
}
