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

// TunnelEndpointSpec defines the desired state of TunnelEndpoint
type TunnelEndpointSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Role (leader/member) shows the role of Gateway node if HA is enabled
	// +kubebuilder:validation:Enum=Leader;Member;Unknown
	// +kubebuilder:validation:Required
	Role string `json:"role,omitempty"`
	// ClusterID of the member cluster.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Required
	ClusterID string `json:"clusterID,omitempty"`
	// Hostname of Gateway node
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Required
	Hostname string `json:"hostname,omitempty"`
	// Subnets used by member cluster in each node.
	// +kubebuilder:validation:Optional
	Subnets []string `json:"subnets,omitempty"`
	// PrivateIP of Gateway node
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`
	// +kubebuilder:validation:Optional
	PrivateIP string `json:"privateIP,omitempty"`
	// PublicIP of Gateway node
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`
	// +kubebuilder:validation:Optional
	PublicIP string `json:"publicIP,omitempty"`
}

// TunnelEndpointStatus defines the observed state of TunnelEndpoint
type TunnelEndpointStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Status of TunnelEndpoint, one word of True, False, Unknown to indicate the TunnelPoint
	// +kubebuilder:validation:Type=string
	Status string `json:"status,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
// +kubebuilder:resource:path=tunnelendpoints,scope=Cluster

// TunnelEndpoint is the Schema for the tunnelendpoints API
type TunnelEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TunnelEndpointSpec   `json:"spec,omitempty"`
	Status TunnelEndpointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TunnelEndpointList contains a list of TunnelEndpoint
type TunnelEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TunnelEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TunnelEndpoint{}, &TunnelEndpointList{})
}
