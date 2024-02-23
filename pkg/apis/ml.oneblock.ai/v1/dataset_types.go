package v1

import (
	v1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ds,scope=Namespaced
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=`.metadata.creationTimestamp`

// Dataset is the Schema for the dataset API
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasetSpec   `json:"spec,omitempty"`
	Status DatasetStatus `json:"status,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	Data     DataSpec            `json:"data,omitempty"`
	Template DatasetTemplateSpec `json:"template,omitempty"`
}

type DataSpec struct {
	Source      DatasetSource      `json:"source,omitempty"`
	Destination DatasetDestination `json:"destination,omitempty"`
	Config      DatasetConfig      `json:"config,omitempty"`
}

type DatasetConfig struct {
}

type DatasetDestination struct {
}

type DatasetSource struct {
}

type DatasetTemplateSpec struct {
	Spec corev1.PodSpec `json:"spec,omitempty"`
}

// DatasetStatus defines the observed state of Dataset
type DatasetStatus struct {
	// Conditions is an array of current conditions
	Conditions []v1.Condition `json:"conditions"`
	// ContainerState is the state of underlying container.
	State corev1.ContainerState `json:"state"`
}
