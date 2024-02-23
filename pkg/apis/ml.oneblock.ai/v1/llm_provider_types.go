package v1

import (
	v1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:scope=Namespaced
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=`.metadata.creationTimestamp`

// LLMProvider is the Schema for the LLM vendors
type LLMProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LLMProviderSpec   `json:"spec,omitempty"`
	Status LLMProviderStatus `json:"status,omitempty"`
}

type LLMProviderSpec struct {
	BaseURL            string                 `json:"baseURL"`
	DefaultModel       string                 `json:"defaultModel,omitempty"`
	DefaultTemperature string                 `json:"defaultTemperature,omitempty"`
	SecretRef          corev1.SecretReference `json:"secret"`
}

type LLMProviderStatus struct {
	// Conditions is an array of current conditions
	Conditions []v1.Condition        `json:"conditions"`
	State      corev1.ContainerState `json:"state"`
}
