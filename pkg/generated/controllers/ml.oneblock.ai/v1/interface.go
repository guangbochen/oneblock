/*
Copyright 2024 1block.ai.

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
	v1 "github.com/oneblock-ai/oneblock/pkg/apis/ml.oneblock.ai/v1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v2/pkg/generic"
	"github.com/rancher/wrangler/v2/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	Dataset() DatasetController
	Notebook() NotebookController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) Dataset() DatasetController {
	return generic.NewController[*v1.Dataset, *v1.DatasetList](schema.GroupVersionKind{Group: "ml.oneblock.ai", Version: "v1", Kind: "Dataset"}, "datasets", true, v.controllerFactory)
}

func (v *version) Notebook() NotebookController {
	return generic.NewController[*v1.Notebook, *v1.NotebookList](schema.GroupVersionKind{Group: "ml.oneblock.ai", Version: "v1", Kind: "Notebook"}, "notebooks", true, v.controllerFactory)
}
