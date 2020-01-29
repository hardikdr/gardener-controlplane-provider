/*

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

// GardenerControlPlaneSpec defines the desired state of GardenerControlPlane
type GardenerControlPlaneSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GardenerControlPlane. Edit GardenerControlPlane_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// GardenerControlPlaneStatus defines the observed state of GardenerControlPlane
type GardenerControlPlaneStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// GardenerControlPlane is the Schema for the gardenercontrolplanes API
type GardenerControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GardenerControlPlaneSpec   `json:"spec,omitempty"`
	Status GardenerControlPlaneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GardenerControlPlaneList contains a list of GardenerControlPlane
type GardenerControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GardenerControlPlane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GardenerControlPlane{}, &GardenerControlPlaneList{})
}
