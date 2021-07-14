/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemberClusterInfoSpec defines the desired state of MemberClusterInfo
type MemberClusterInfoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ClusterID of the member cluster.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Required
	ClusterID string `json:"clusterID,omitempty"`
	// Service's ClusterIP CIDR
	// +kubebuilder:validation:Required
	ServiceCIDR []string `json:"serviceCIDR,omitempty"`
	// Cluster's PoD Global CIDR
	// +kubebuilder:validation:Required
	ClusterCIDR []string `json:"clusterCIDR,omitempty"`
}

//+kubebuilder:object:root=true
// +kubebuilder:resource:path=memberclusterinfoes,scope=Cluster
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemberClusterInfo is the Schema for the memberclusterinfoes API
type MemberClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MemberClusterInfoSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// MemberClusterInfoList contains a list of MemberClusterInfo
type MemberClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemberClusterInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemberClusterInfo{}, &MemberClusterInfoList{})
}
