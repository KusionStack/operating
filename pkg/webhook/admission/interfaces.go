/*
Copyright 2023 The KusionStack Authors.

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

package admission

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// Interface is an abstract, pluggable interface for Admission Control decisions.
type Interface interface {
	Name() string
}

type PluginInterface interface {
	Interface
	// Validate makes an admission decision based on the request attributes.  It is NOT allowed to mutate
	Validate(ctx context.Context, req admission.Request, obj runtime.Object) error
	// Admit makes an admission decision based on the request attributes
	Admit(ctx context.Context, req admission.Request, obj runtime.Object) error
}

type ValidationFunc func(ctx context.Context, req admission.Request, obj runtime.Object) error
type MutationFunc func(ctx context.Context, req admission.Request, obj runtime.Object) error

type DispatchHandler interface {
	Handle(context.Context, admission.Request) admission.Response
}
