/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HomePageEntrySpec defines the desired state of HomePageEntry
type HomePageEntrySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

    Type string `json:"type"`

    Name string `json:"name"`

    Namespace string `json:"namespace"`
	
	DisplayName string `json:"displayName"`
}

// HomePageEntryStatus defines the observed state of HomePageEntry
type HomePageEntryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Active bool `json:"active"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HomePageEntry is the Schema for the homepageentries API
type HomePageEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HomePageEntrySpec   `json:"spec,omitempty"`
	Status HomePageEntryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HomePageEntryList contains a list of HomePageEntry
type HomePageEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HomePageEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HomePageEntry{}, &HomePageEntryList{})
}
