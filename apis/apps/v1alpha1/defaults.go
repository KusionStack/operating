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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	k8scorev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	appsv1alpha1 "kusionstack.io/kube-api/apps/v1alpha1"
)

func SetDefaultCollaSet(cls *appsv1alpha1.CollaSet) {
	SetDefaultPodSpec(cls)
	SetDefaultCollaSetUpdateStrategy(cls)
}

func SetDefaultPodSpec(in *appsv1alpha1.CollaSet) {
	k8scorev1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	SetDefaultPodSpecVolumes(in.Spec.Template.Spec.Volumes)
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		k8scorev1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					k8scorev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		k8scorev1.SetDefaults_ResourceList(&a.Resources.Limits)
		k8scorev1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			k8scorev1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			k8scorev1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.StartupProbe != nil {
			k8scorev1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.StartupProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					k8scorev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					k8scorev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		k8scorev1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					k8scorev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		k8scorev1.SetDefaults_ResourceList(&a.Resources.Limits)
		k8scorev1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			k8scorev1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			k8scorev1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.StartupProbe != nil {
			k8scorev1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.StartupProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					k8scorev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					k8scorev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.EphemeralContainers {
		a := &in.Spec.Template.Spec.EphemeralContainers[i]
		k8scorev1.SetDefaults_EphemeralContainer(a)
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.EphemeralContainerCommon.Env {
			b := &a.EphemeralContainerCommon.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					k8scorev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		k8scorev1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Limits)
		k8scorev1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Requests)
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			k8scorev1.SetDefaults_Probe(a.EphemeralContainerCommon.LivenessProbe)
			if a.EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			k8scorev1.SetDefaults_Probe(a.EphemeralContainerCommon.ReadinessProbe)
			if a.EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			k8scorev1.SetDefaults_Probe(a.EphemeralContainerCommon.StartupProbe)
			if a.EphemeralContainerCommon.StartupProbe.Handler.HTTPGet != nil {
				k8scorev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.StartupProbe.Handler.HTTPGet)
			}
		}
		if a.EphemeralContainerCommon.Lifecycle != nil {
			if a.EphemeralContainerCommon.Lifecycle.PostStart != nil {
				if a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet != nil {
					k8scorev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.EphemeralContainerCommon.Lifecycle.PreStop != nil {
				if a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet != nil {
					k8scorev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	k8scorev1.SetDefaults_ResourceList(&in.Spec.Template.Spec.Overhead)
}

func SetDefaultPodSpecVolumes(volumes []corev1.Volume) {
	for i := range volumes {
		a := &volumes[i]
		k8scorev1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			k8scorev1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			k8scorev1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			k8scorev1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			k8scorev1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			k8scorev1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					k8scorev1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			k8scorev1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			k8scorev1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			k8scorev1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							k8scorev1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
				if b.ServiceAccountToken != nil {
					k8scorev1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			k8scorev1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
		if a.VolumeSource.Ephemeral != nil {
			if a.VolumeSource.Ephemeral.VolumeClaimTemplate != nil {
				k8scorev1.SetDefaults_PersistentVolumeClaimSpec(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec)
				k8scorev1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Limits)
				k8scorev1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Requests)
			}
		}
	}

}

func SetDefaultCollaSetUpdateStrategy(cls *appsv1alpha1.CollaSet) {
	if cls.Spec.UpdateStrategy.PodUpdatePolicy == "" {
		cls.Spec.UpdateStrategy.PodUpdatePolicy = appsv1alpha1.CollaSetInPlaceIfPossiblePodUpdateStrategyType
	}

	if cls.Spec.UpdateStrategy.RollingUpdate == nil {
		cls.Spec.UpdateStrategy.RollingUpdate = &appsv1alpha1.RollingUpdateCollaSetStrategy{}
	}

	if cls.Spec.UpdateStrategy.RollingUpdate.ByPartition == nil && cls.Spec.UpdateStrategy.RollingUpdate.ByLabel == nil {
		cls.Spec.UpdateStrategy.RollingUpdate.ByPartition = &appsv1alpha1.ByPartition{}
	}
}
