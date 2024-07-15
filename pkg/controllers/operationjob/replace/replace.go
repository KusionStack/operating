/*
Copyright 2024 The KusionStack Authors.

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

package replace

import (
	"encoding/json"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "kusionstack.io/operating/apis/apps/v1alpha1"
	"kusionstack.io/operating/pkg/controllers/collaset/podcontrol"
	. "kusionstack.io/operating/pkg/controllers/operationjob/opscontrol"
)

type PodReplaceControl struct {
	*OperateInfo
	PodControl podcontrol.Interface
}

func (p *PodReplaceControl) OperateTarget(candidate *OpsCandidate) error {
	// parse replace information from origin pod
	var replaceIndicated, replaceByReplaceUpdate, replaceNewPodExists bool
	_, replaceIndicated = candidate.Pod.Labels[appsv1alpha1.PodReplaceIndicationLabelKey]
	_, replaceByReplaceUpdate = candidate.Pod.Labels[appsv1alpha1.PodReplaceByReplaceUpdateLabelKey]
	_, replaceNewPodExists = candidate.Pod.Labels[appsv1alpha1.PodReplacePairNewId]

	// label pod to trigger replace
	replaceTriggered := replaceIndicated || replaceByReplaceUpdate || replaceNewPodExists
	if !replaceTriggered {
		patch := client.RawPatch(types.StrategicMergePatchType, []byte(fmt.Sprintf(`{"metadata":{"labels":{"%s":"%v"}}}`, appsv1alpha1.PodReplaceIndicationLabelKey, true)))
		if err := p.Client.Patch(p.Context, candidate.Pod, patch); err != nil {
			return fmt.Errorf("fail to label origin pod %s/%s with replace indicate label by replaceUpdate: %s", candidate.Pod.Namespace, candidate.Pod.Name, err)
		}
	}

	return nil
}

func (p *PodReplaceControl) FulfilTargetOpsStatus(candidate *OpsCandidate) error {
	// try to find replaceNewPod
	if candidate.Pod != nil && candidate.CollaSet != nil {
		newPodId, exist := candidate.Pod.Labels[appsv1alpha1.PodReplacePairNewId]
		if exist {
			filteredPods, err := p.PodControl.GetFilteredPods(candidate.CollaSet.Spec.Selector, candidate.CollaSet)
			if err != nil {
				return err
			}
			for _, newPod := range filteredPods {
				if newPodId == newPod.Labels[appsv1alpha1.PodInstanceIDLabelKey] {
					p.Recorder.Eventf(candidate.Pod, corev1.EventTypeNormal, "ReplaceNewPod", "replace by pod %s with operationjob %s", candidate.PodName, p.OperationJob.Name)
					p.Recorder.Eventf(newPod, corev1.EventTypeNormal, "ReplaceOriginPod", "replace pod %s with operationjob %s", newPod.Name, p.OperationJob.Name)
					candidate.OpsStatus.Reason = appsv1alpha1.ReasonReplacedByNewPod
					candidate.OpsStatus.Message = newPod.Name
				}
			}
		}
	}

	// origin pod is deleted not exist, mark as succeeded
	if candidate.Pod == nil {
		candidate.OpsStatus.Progress = appsv1alpha1.OperationProgressSucceeded
		if candidate.OpsStatus.Reason != appsv1alpha1.ReasonReplacedByNewPod {
			candidate.OpsStatus.Reason = appsv1alpha1.ReasonPodNotFound
		}
	} else {
		candidate.OpsStatus.Progress = appsv1alpha1.OperationProgressProcessing
	}

	return nil
}

func (p *PodReplaceControl) ReleaseTarget(candidate *OpsCandidate) error {
	if candidate.Pod == nil || candidate.Pod.DeletionTimestamp != nil {
		return nil
	}

	// try to remove replace label from origin pod
	patchOperation := map[string]string{
		"op":   "remove",
		"path": fmt.Sprintf("/metadata/labels/%s", strings.ReplaceAll(appsv1alpha1.PodReplaceIndicationLabelKey, "/", "~1")),
	}

	patchBytes, err := json.Marshal([]map[string]string{patchOperation})
	if err != nil {
		return err
	}

	if err := p.PodControl.PatchPod(candidate.Pod, client.RawPatch(types.JSONPatchType, patchBytes)); err != nil {
		return fmt.Errorf("failed to remove to-replace label %s/%s: %s", candidate.Pod.Namespace, candidate.Pod.Name, err)
	}
	return nil
}
