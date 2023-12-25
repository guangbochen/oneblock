package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Display Name",type="string",JSONPath=`.displayName`

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpec struct {
	// optional
	// +kubebuilder:default=2
	// +kubebuilder:validation:Minimum=0
	MaxWorkers int `json:"maxWorkers,omitempty"`
	// optional
	UpScalingSpeed float64 `json:"upScalingSpeed,omitempty"`
	// optional
	IdleTimeoutMinutes int `json:"idleTimeoutMinutes,omitempty"`
	// optional
	Docker DockerConfig `json:"docker,omitempty"`
	// +kubebuilder:validation:Required
	Provider ProviderConfig `json:"provider"`
	// +kubebuilder:validation:Required
	Auth AuthConfig `json:"auth"`
	// optional
	AvailableNodeTypes []NodeType `json:"availableNodeTypes,omitempty"`
	// optional
	HeadNodeType string `json:"headNodeType,omitempty"`
	// optional
	//FileMounts []FileMount `json:"fileMounts,omitempty"`
	// optional
	ClusterSyncedFiles []string `json:"clusterSyncedFiles,omitempty"`
	// optional
	RsyncExclude []string `json:"rsyncExclude,omitempty"`
	// optional
	RsyncFilter []string `json:"rsyncFilter,omitempty"`
	// optional
	InitializationCommands []string `json:"initializationCommands,omitempty"`
	// optional
	SetupCommands []string `json:"setupCommands,omitempty"`
	// optional
	HeadSetupCommands []string `json:"headSetupCommands,omitempty"`
	// optional
	WorkerSetupCommands []string `json:"workerSetupCommands,omitempty"`
	// optional
	HeadStartRayCommands []string `json:"headStartRayCommands,omitempty"`
	// optional
	WorkerStartRayCommands []string `json:"workerStartRayCommands,omitempty"`
}

type DockerConfig struct {
	// optional
	Image string `json:"image"`
	// optional
	HeadImage string `json:"headImage,omitempty"`
	// optional
	WorkerImage string `json:"workerImage,omitempty"`
	// optional
	ContainerName string `json:"containerName,omitempty"`
	// optional
	PullBeforeRun bool `json:"pullBeforeRun,omitempty"`
	// optional
	RunOptions []string `json:"runOptions,omitempty"`
	// optional
	HeadRunOptions []string `json:"headRunOptions,omitempty"`
	// optional
	WorkerRunOptions []string `json:"workerRunOptions,omitempty"`
	// optional
	DisableAutomaticRuntimeDetection bool `json:"disableAutomaticRuntimeDetection,omitempty"`
	// optional
	DisableShmSizeDetection bool `json:"disableShmSizeDetection,omitempty"`
}

type ProviderType string

const (
	ProviderTypeAWS ProviderType = "aws"
	ProviderTypeGCP ProviderType = "gcp"
)

type ProviderConfig struct {
	// +kubebuilder:validation:Required
	Type ProviderType `json:"type"`
	// optional
	Region string `json:"region,omitempty"`
	// optional
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	// optional
	Location string `json:"location,omitempty"`
	// optional
	ResourceGroup string `json:"resourceGroup,omitempty"`
	// optional
	SubscriptionID string `json:"subscriptionID,omitempty"`
	// optional
	ProjectID string `json:"projectID,omitempty"`
	// optional
	CacheStoppedNodes bool `json:"cacheStoppedNodes,omitempty"`
	// optional
	SecurityGroup SecurityGroup `json:"securityGroup,omitempty"`
	// optional
	UseInternalIPs bool `json:"useInternalIPs,omitempty"`
}

type SecurityGroup struct {
	// optional
	GroupName string `json:"groupName,omitempty"`
	// optional
	//IpPermissions []IpPermission `json:"ipPermissions,omitempty"`
}

type AuthConfig struct {
	// +kubebuilder:validation:Required
	SSHUser string `json:"sshUser"`
	// optional
	SSHPrivateKey string `json:"sshPrivateKey,omitempty"`
	// optional
	SSHPublicKey string `json:"sshPublicKey,omitempty"`
}

type NodeType struct {
	// optional
	NodeConfig map[string]interface{} `json:"nodeConfig,omitempty"`
	// optional
	Resources Resources `json:"resources,omitempty"`
	// optional
	MinWorkers int `json:"minWorkers,omitempty"`
	// optional
	MaxWorkers int `json:"maxWorkers,omitempty"`
	// optional
	WorkerSetupCommands []string `json:"workerSetupCommands,omitempty"`
	// optional
	Docker NodeDocker `json:"docker,omitempty"`
}

type Resources struct {
	// optional
	CPU int `json:"cpu,omitempty"`
	// optional
	GPU int `json:"gpu,omitempty"`
	// optional
	ObjectStoreMemory int `json:"objectStoreMemory,omitempty"`
	// optional
	Memory int `json:"memory,omitempty"`
}

type NodeDocker struct {
	// optional
	WorkerImage string `json:"workerImage,omitempty"`
	// optional
	PullBeforeRun bool `json:"pullBeforeRun,omitempty"`
	// optional
	WorkerRunOptions []string `json:"workerRunOptions,omitempty"`
	// optional
	DisableAutomaticRuntimeDetection bool `json:"disableAutomaticRuntimeDetection,omitempty"`
	// optional
	DisableShmSizeDetection bool `json:"disableShmSizeDetection,omitempty"`
}

type ClusterStatus struct {
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}
