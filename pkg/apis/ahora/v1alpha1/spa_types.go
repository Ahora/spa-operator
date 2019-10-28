package v1alpha1

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SPASpec defines the desired state of SPA
// +k8s:openapi-gen=true
type SPASpec struct {
	Replicas      *int32               `json:"replicas"`
	SPAArchiveURL string               `json:"SPAArchiveURL"`
	TLS           []v1beta1.IngressTLS `json:"tls"`
	Hosts         []string             `json:"hosts,omitempty" protobuf:"bytes,1,rep,name=hosts"`
}

// SPAStatus defines the observed state of SPA
// +k8s:openapi-gen=true
type SPAStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SPA is the Schema for the spas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=spas,scope=Namespaced
type SPA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SPASpec   `json:"spec,omitempty"`
	Status SPAStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SPAList contains a list of SPA
type SPAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SPA `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SPA{}, &SPAList{})
}
