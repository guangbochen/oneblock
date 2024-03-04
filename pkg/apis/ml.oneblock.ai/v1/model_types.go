package v1

import (
	"github.com/rancher/wrangler/v2/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
)

var (
	ModelTemplateVersionConfigured condition.Cond = "configured"
)

type EngineType string

const (
	EngineTypeVLLM      EngineType = "VLLMEngine"
	EngineTypeEmbedding EngineType = "EmbeddingEngine"
)

type PlacementStrategy string

const (
	PlacementStrategyStrictPack PlacementStrategy = "STRICT_PACK"
	PlacementStrategyPack       PlacementStrategy = "PACK"
	PlacementStrategySpread     PlacementStrategy = "SPREAD"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:scope=Namespaced
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=`.metadata.creationTimestamp`

// ModelTemplateVersion is the Schema for the LLM model
type ModelTemplateVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModelTemplateVersionSpec   `json:"spec,omitempty"`
	Status ModelTemplateVersionStatus `json:"status,omitempty"`
}

type ModelTemplateVersionSpec struct {
	// +kubebuilder:validation:Required
	// ModelID is the ID that refers to the model in the OpenAI API.
	ModelID string `json:"modelID"`
	// HFModelID is the Hugging Face model ID. If not specified, defaults to modelID
	HFModelID string `json:"hfModelID,omitempty"`
	// MirrorConfig helps to add a private model, you can either choose to use an S3 or GCS mirror.
	MirrorConfig     string           `json:"mirrorConfig,omitempty"`
	EngineConfig     EngineConfig     `json:"engineConfig,omitempty"`
	DeploymentConfig DeploymentConfig `json:"deploymentConfig,omitempty"`
	ScalingConfig    ScalingConfig    `json:"scalingConfig,omitempty"`
	WorkerGroupSpec  WorkerGroupSpec  `json:"workerGroupSpec,omitempty"`
}

type ScalingConfig struct {
	// +kubebuilder:default:=1
	NumWorkers int32 `json:"numWorkers"`
	// +kubebuilder:default:=1
	NumGPUsPerWorker int32 `json:"numGPUsPerWorker,omitempty"`
	// +kubebuilder:default:=5
	NumCPUsPerWorker int32 `json:"numCPUsPerWorker"`
	// +kubebuilder:default:=STRICT_PACK
	PlacementStrategy PlacementStrategy `json:"placementStrategy"`
	// You can use custom resources to specify the instance type/accelerator type to use for the model.
	// e.g., accelerator_type_a10: 0.01
	ResourcesPerWorker map[string]string `json:"resourcesPerWorker,omitempty"`
}

type EngineConfig struct {
	// +kubebuilder:default:=VLLMEngine
	Type EngineType `json:"type,omitempty"`
	// +kubebuilder:validation:Required
	MaxTotalTokens int32 `json:"maxTotalTokens"`
	// More details about engine config can be referred to:
	// vLLM: https://github.com/vllm-project/vllm/blob/main/vllm/config.py
	VLLMArgs   string           `json:"vLLMArgs,omitempty"`
	Generation GenerationConfig `json:"generation,omitempty"`
}

type GenerationConfig struct {
	PromptFormat      `json:"promptFormat,omitempty"`
	StoppingSequences []string `json:"stoppingSequences,omitempty"`
}

type PromptFormat struct {
	// The format of the prompt. The following fields are available:
	// System message. Will default to empty
	System string `json:"system,omitempty"`
	// Past assistant message. Used in chat completions API.
	Assistant string `json:"assistant,omitempty"`
	// New assistant message. After this point, model will generate tokens.
	TrailingAssistant string `json:"trailingAssistant,omitempty"`
	// User message
	User string `json:"user,omitempty"`
	// Default system message.
	DefaultSystemMessage string `json:"defaultSystemMessage,omitempty"`
	// Whether the system prompt is inside the user prompt. If true, the user field should include '{system}'
	SystemInUser bool `json:"systemInUser,omitempty"`
	// Whether to include the system tags even if the user message is empty.
	AddSystemTagsEvenIfMessageIsEmpty bool `json:"addSystemTagsEvenIfMessageIsEmpty,omitempty"`
	// Whether to automatically strip whitespace from left and right of user supplied messages for chat completions
	// +kubebuilder:default:=true
	StripWhitespace bool `json:"stripWhitespace,omitempty"`
}

type DeploymentConfig struct {
	// +kubebuilder:default:=1
	Replicas int32 `json:"replicas"`
	// +kubebuilder:default:=1
	MinReplicas int32 `json:"minReplicas"`
	// +kubebuilder:default:=5
	MaxReplicas int32 `json:"maxReplicas"`
	// +kubebuilder:default:=64
	MaxConcurrentQueries int32 `json:"maxConcurrentQueries"`
	// Auto scale up/down the number of replicas if the average number of ongoing requests is above/below this value.
	// Automatically set this to 40% of the maxConcurrentQueries if not specified.
	TargetNumOngoingRequests int32 `json:"targetNumOngoingRequests,omitempty"`
}

type ModelTemplateVersionStatus struct {
	// Conditions is an array of current conditions
	Conditions           []v1.Condition `json:"conditions,omitempty"`
	GeneratedModelConfig string         `json:"generatedModelConfig,omitempty"`
}
