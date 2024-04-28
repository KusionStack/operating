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

package collaset

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	appsv1alpha1 "kusionstack.io/operating/apis/apps/v1alpha1"
	commonutils "kusionstack.io/operating/pkg/utils"
	"kusionstack.io/operating/pkg/utils/mixin"
)

type MutatingHandler struct {
	*mixin.WebhookHandlerMixin
}

func NewMutatingHandler() *MutatingHandler {
	return &MutatingHandler{
		WebhookHandlerMixin: mixin.NewWebhookHandlerMixin(),
	}
}

func (h *MutatingHandler) Handle(ctx context.Context, req admission.Request) (resp admission.Response) {
	if req.Operation != admissionv1.Update && req.Operation != admissionv1.Create {
		return admission.Allowed("")
	}

	logger := h.Logger.WithValues(
		"op", req.Operation,
		"collaset", commonutils.AdmissionRequestObjectKeyString(req),
	)

	cls := &appsv1alpha1.CollaSet{}
	if err := h.Decoder.Decode(req, cls); err != nil {
		logger.Error(err, "failed to decode collaset")
		return admission.Errored(http.StatusBadRequest, err)
	}
	appsv1alpha1.SetDetaultCollaSet(cls)
	marshalled, err := json.Marshal(cls)
	if err != nil {
		logger.Error(err, "failed to marshal collaset to json")
		return admission.Errored(http.StatusInternalServerError, err)
	}

	for _, podName := range cls.Spec.ScaleStrategy.PodToDelete {
		pod := &corev1.Pod{}
		if err := h.Client.Get(ctx, types.NamespacedName{Namespace: cls.Namespace, Name: podName}, pod); err != nil {
			if errors.IsNotFound(err) {
				return admission.Errored(http.StatusBadRequest, fmt.Errorf("no found Pod named %s to delete", podName))
			}
			return admission.Errored(http.StatusInternalServerError, fmt.Errorf("failed to find Pod %s: %v", podName, err))
		}
		owned := false
		for _, ref := range pod.OwnerReferences {
			if ref.Name == cls.Name {
				owned = true
			}
		}
		if !owned {
			return admission.Errored(http.StatusBadRequest, fmt.Errorf("not allwed to delete pod %s which is not owned by collaset %s", podName, cls.Name))
		}
	}

	return admission.PatchResponseFromRaw(req.AdmissionRequest.Object.Raw, marshalled)
}

var _ inject.Client = &MutatingHandler{}

func (h *MutatingHandler) InjectClient(c client.Client) error {
	h.Client = c
	return nil
}

var _ admission.DecoderInjector = &MutatingHandler{}

func (h *MutatingHandler) InjectDecoder(d *admission.Decoder) error {
	h.Decoder = d
	return nil
}
