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

type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AgentSpec   `json:"spec,omitempty"`
	Status AgentStatus `json:"status,omitempty"`
}

type AgentSpec struct {
	Prompt        string            `json:"prompt,omitempty"`
	OpeningDialog string            `json:"openingDialog,omitempty"`
	OpeningQuery  string            `json:"openingQuery,omitempty"`
	Datasets      []DatasetRef      `json:"datasets,omitempty"`
	LLMProvider   LLMProviderRef    `json:"llmProviderRef,omitempty"`
	Tools         []AgentTool       `json:"tools,omitempty"`
	Inputs        map[string]string `json:"inputs,omitempty"`
	Outputs       map[string]string `json:"outputs,omitempty"`
}

type AgentTool struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Headers     map[string]string      `json:"headers,omitempty"`
	URL         string                 `json:"url,omitempty"`
	Path        string                 `json:"path,omitempty"`
	SecretRef   corev1.SecretReference `json:"secretRef,omitempty"`
}

type DatasetRef struct {
	Name      string                 `json:"name,omitempty"`
	Namespace string                 `json:"namespace,omitempty"`
	URL       string                 `json:"url,omitempty"`
	SecretRef corev1.SecretReference `json:"secretRef,omitempty"`
}

type LLMProviderRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

type AgentStatus struct {
	// Conditions is an array of current conditions
	Conditions []v1.Condition        `json:"conditions"`
	State      corev1.ContainerState `json:"state"`
}
