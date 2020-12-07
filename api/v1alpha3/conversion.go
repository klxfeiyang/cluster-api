/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *Cluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.Cluster)

	return Convert_v1alpha3_Cluster_To_v1alpha4_Cluster(src, dst, nil)
}

func (dst *Cluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.Cluster)

	return Convert_v1alpha4_Cluster_To_v1alpha3_Cluster(src, dst, nil)
}

func (src *ClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.ClusterList)

	return Convert_v1alpha3_ClusterList_To_v1alpha4_ClusterList(src, dst, nil)
}

func (dst *ClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.ClusterList)

	return Convert_v1alpha4_ClusterList_To_v1alpha3_ClusterList(src, dst, nil)
}

func (src *Machine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.Machine)

	return Convert_v1alpha3_Machine_To_v1alpha4_Machine(src, dst, nil)
}

func (dst *Machine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.Machine)

	return Convert_v1alpha4_Machine_To_v1alpha3_Machine(src, dst, nil)
}

func (src *MachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineList)

	return Convert_v1alpha3_MachineList_To_v1alpha4_MachineList(src, dst, nil)
}

func (dst *MachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineList)

	return Convert_v1alpha4_MachineList_To_v1alpha3_MachineList(src, dst, nil)
}

func (src *MachineSet) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineSet)

	return Convert_v1alpha3_MachineSet_To_v1alpha4_MachineSet(src, dst, nil)
}

func (dst *MachineSet) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineSet)

	return Convert_v1alpha4_MachineSet_To_v1alpha3_MachineSet(src, dst, nil)
}

func (src *MachineSetList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineSetList)

	return Convert_v1alpha3_MachineSetList_To_v1alpha4_MachineSetList(src, dst, nil)
}

func (dst *MachineSetList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineSetList)

	return Convert_v1alpha4_MachineSetList_To_v1alpha3_MachineSetList(src, dst, nil)
}

func (src *MachineDeployment) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineDeployment)

	return Convert_v1alpha3_MachineDeployment_To_v1alpha4_MachineDeployment(src, dst, nil)
}

func (dst *MachineDeployment) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineDeployment)

	return Convert_v1alpha4_MachineDeployment_To_v1alpha3_MachineDeployment(src, dst, nil)
}

func (src *MachineDeploymentList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineDeploymentList)

	return Convert_v1alpha3_MachineDeploymentList_To_v1alpha4_MachineDeploymentList(src, dst, nil)
}

func (dst *MachineDeploymentList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineDeploymentList)

	return Convert_v1alpha4_MachineDeploymentList_To_v1alpha3_MachineDeploymentList(src, dst, nil)
}

func (src *MachineHealthCheck) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineHealthCheck)

	return Convert_v1alpha3_MachineHealthCheck_To_v1alpha4_MachineHealthCheck(src, dst, nil)
}

func (dst *MachineHealthCheck) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineHealthCheck)

	return Convert_v1alpha4_MachineHealthCheck_To_v1alpha3_MachineHealthCheck(src, dst, nil)
}

func (src *MachineHealthCheckList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha4.MachineHealthCheckList)

	return Convert_v1alpha3_MachineHealthCheckList_To_v1alpha4_MachineHealthCheckList(src, dst, nil)
}

func (dst *MachineHealthCheckList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha4.MachineHealthCheckList)

	return Convert_v1alpha4_MachineHealthCheckList_To_v1alpha3_MachineHealthCheckList(src, dst, nil)
}

// Convert_v1alpha3_Bootstrap_To_v1alpha4_Bootstrap is an autogenerated conversion function.
func Convert_v1alpha3_Bootstrap_To_v1alpha4_Bootstrap(in *Bootstrap, out *v1alpha4.Bootstrap, s apiconversion.Scope) error { //nolint
	return autoConvert_v1alpha3_Bootstrap_To_v1alpha4_Bootstrap(in, out, s)
}
