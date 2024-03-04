package modeltemplate

import (
	"fmt"
	"strings"

	mlv1 "github.com/oneblock-ai/oneblock/pkg/apis/ml.oneblock.ai/v1"
	"gopkg.in/yaml.v2"
)

type RayLLMModelConfig struct {
	DeploymentConfig DeploymentConfig `yaml:"deployment_config"`
	EngineConfig     EngineConfig     `yaml:"engine_config"`
	ScalingConfig    ScalingConfig    `yaml:"scaling_config"`
}

type ScalingConfig struct {
	NumWorkers         int32             `yaml:"num_workers"`
	NumGPUsPerWorker   int32             `yaml:"num_gpus_per_worker"`
	NumCPUsPerWorker   int32             `yaml:"num_cpus_per_worker"`
	PlacementStrategy  string            `yaml:"placement_strategy"`
	ResourcesPerWorker map[string]string `yaml:"resources_per_worker"`
}

type EngineConfig struct {
	ModelID         string                 `yaml:"model_id"`
	HFModelID       string                 `yaml:"hf_model_id,omitempty"`
	S3MirrorConfig  MirrorConfig           `yaml:"s3_mirror_config,omitempty"`
	GCSMirrorConfig MirrorConfig           `yaml:"gcs_mirror_config,omitempty"`
	Type            string                 `yaml:"type"`
	EngineKwargs    map[string]interface{} `yaml:"engine_kwargs"`
	MaxTotalTokens  int32                  `yaml:"max_total_tokens"`
	Generation      GenerationConfig       `yaml:"generation"`
}

type MirrorConfig struct {
	BucketURI string `yaml:"bucket_uri"`
}

type GenerationConfig struct {
	PromptFormat      PromptFormat `yaml:"prompt_format"`
	StoppingSequences []string     `yaml:"stopping_sequences"`
}

type PromptFormat struct {
	System                            string `yaml:"system"`
	Assistant                         string `yaml:"assistant"`
	TrailingAssistant                 string `yaml:"trailing_assistant"`
	User                              string `yaml:"user"`
	DefaultSystemMessage              string `yaml:"default_system_message"`
	SystemInUser                      bool   `yaml:"system_in_user"`
	AddSystemTagsEvenIfMessageIsEmpty bool   `yaml:"add_system_tags_even_if_message_is_empty"`
	StripWhitespace                   bool   `yaml:"strip_whitespace"`
}

type DeploymentConfig struct {
	AutoScalingConfig    AutoScalingConfig `yaml:"auto_scaling_config"`
	MaxConcurrentQueries int32             `yaml:"max_concurrent_queries"`
	RayActorOptions      RayActorOptions   `yaml:"ray_actor_options"`
}

type RayActorOptions struct {
	Resources map[string]string `yaml:"resources"`
}

type AutoScalingConfig struct {
	MinReplicas                        int32   `yaml:"min_replicas"`
	MaxReplicas                        int32   `yaml:"max_replicas"`
	InitialReplicas                    int32   `yaml:"initial_replicas"`
	TargetNumOngoingRequestsPerReplica int32   `yaml:"target_num_ongoing_requests_per_replica"`
	MetricsIntervalS                   float32 `yaml:"metrics_interval_s"`
	LookBackPeriodS                    float32 `yaml:"look_back_period_s"`
	SmoothingFactor                    float32 `yaml:"smoothing_factor"`
	DownscaleDelayS                    float32 `yaml:"downscale_delay_s"`
	UpscaleDelayS                      float32 `yaml:"upscale_delay_s"`
}

const (
	MaxConcurrentRatio = 40
)

func generateRayLLMModelConfig(modelTmpVersion *mlv1.ModelTemplateVersion) (string, error) {
	if modelTmpVersion == nil {
		return "", fmt.Errorf("model template version is empty")
	}

	rayLLMModelConfig := RayLLMModelConfig{}
	// define deployment config
	maxConcurrentQueries := modelTmpVersion.Spec.DeploymentConfig.MaxConcurrentQueries
	deploymentConfig := DeploymentConfig{
		AutoScalingConfig: AutoScalingConfig{
			MinReplicas:                        modelTmpVersion.Spec.DeploymentConfig.MinReplicas,
			MaxReplicas:                        modelTmpVersion.Spec.DeploymentConfig.MaxReplicas,
			InitialReplicas:                    modelTmpVersion.Spec.DeploymentConfig.Replicas,
			TargetNumOngoingRequestsPerReplica: modelTmpVersion.Spec.DeploymentConfig.TargetNumOngoingRequests,
			MetricsIntervalS:                   10.0,
			LookBackPeriodS:                    30.0,
			SmoothingFactor:                    0.6,
			DownscaleDelayS:                    300.0,
			UpscaleDelayS:                      60.0,
		},
		MaxConcurrentQueries: maxConcurrentQueries,
		RayActorOptions: RayActorOptions{
			Resources: modelTmpVersion.Spec.ScalingConfig.ResourcesPerWorker,
		},
	}
	if deploymentConfig.AutoScalingConfig.TargetNumOngoingRequestsPerReplica == 0 {
		deploymentConfig.AutoScalingConfig.TargetNumOngoingRequestsPerReplica = (maxConcurrentQueries * MaxConcurrentRatio) / 100
	}
	rayLLMModelConfig.DeploymentConfig = deploymentConfig

	// define engine config
	prompt := modelTmpVersion.Spec.EngineConfig.Generation.PromptFormat
	engineConfig := EngineConfig{
		ModelID:        modelTmpVersion.Spec.ModelID,
		HFModelID:      modelTmpVersion.Spec.HFModelID,
		Type:           string(modelTmpVersion.Spec.EngineConfig.Type),
		MaxTotalTokens: modelTmpVersion.Spec.EngineConfig.MaxTotalTokens,
		Generation: GenerationConfig{
			PromptFormat: PromptFormat{
				System:                            prompt.System,
				Assistant:                         prompt.Assistant,
				TrailingAssistant:                 prompt.TrailingAssistant,
				User:                              prompt.User,
				DefaultSystemMessage:              prompt.DefaultSystemMessage,
				SystemInUser:                      prompt.SystemInUser,
				AddSystemTagsEvenIfMessageIsEmpty: prompt.AddSystemTagsEvenIfMessageIsEmpty,
				StripWhitespace:                   prompt.StripWhitespace,
			},
			StoppingSequences: modelTmpVersion.Spec.EngineConfig.Generation.StoppingSequences,
		},
	}
	if modelTmpVersion.Spec.EngineConfig.VLLMArgs != "" {
		if err := yaml.Unmarshal([]byte(modelTmpVersion.Spec.EngineConfig.VLLMArgs), &engineConfig.EngineKwargs); err != nil {
			return "", fmt.Errorf("failed to convert vllmArgs, error: %s", err.Error())
		}
	} else {
		engineConfig.EngineKwargs = map[string]interface{}{
			"trust_remote_code":      true,
			"max_num_seq":            32,
			"max_num_batched_tokens": modelTmpVersion.Spec.EngineConfig.MaxTotalTokens,
			"gpu_memory_utilization": 0.9,
		}
	}

	if engineConfig.HFModelID == "" && modelTmpVersion.Spec.MirrorConfig != "" {
		if strings.Contains(modelTmpVersion.Spec.MirrorConfig, "s3://") {
			engineConfig.S3MirrorConfig.BucketURI = modelTmpVersion.Spec.MirrorConfig
		} else if strings.Contains(modelTmpVersion.Spec.MirrorConfig, "gs://") {
			engineConfig.GCSMirrorConfig.BucketURI = modelTmpVersion.Spec.MirrorConfig
		} else {
			return "", fmt.Errorf("invalid mirror config: %s", modelTmpVersion.Spec.MirrorConfig)
		}
	}

	// set default generation config if its value is empty string
	engineConfig.Generation = setDefaultGeneration(engineConfig.Generation)

	rayLLMModelConfig.EngineConfig = engineConfig

	// define scaling config
	scalingConfig := ScalingConfig{
		NumWorkers:         modelTmpVersion.Spec.ScalingConfig.NumWorkers,
		NumGPUsPerWorker:   modelTmpVersion.Spec.ScalingConfig.NumGPUsPerWorker,
		NumCPUsPerWorker:   modelTmpVersion.Spec.ScalingConfig.NumCPUsPerWorker,
		PlacementStrategy:  string(modelTmpVersion.Spec.ScalingConfig.PlacementStrategy),
		ResourcesPerWorker: modelTmpVersion.Spec.ScalingConfig.ResourcesPerWorker,
	}
	rayLLMModelConfig.ScalingConfig = scalingConfig
	yamlModelConfig, err := yaml.Marshal(&rayLLMModelConfig)
	if err != nil {
		return "", fmt.Errorf("failed to convert to YAML config, error: %s", err.Error())

	}
	return string(yamlModelConfig), nil
}

func setDefaultGeneration(generation GenerationConfig) GenerationConfig {
	promptFormat := generation.PromptFormat
	if promptFormat.System == "" {
		promptFormat.System = "{instruction}\\n\\n"
	}
	if promptFormat.Assistant == "" {
		promptFormat.Assistant = " {instruction} </s><s>"
	}
	if promptFormat.User == "" {
		promptFormat.User = "[INST] {system}{instruction} [/INST]"
	}
	if promptFormat.SystemInUser == false {
		promptFormat.SystemInUser = true
	}
	generation.PromptFormat = promptFormat
	if generation.StoppingSequences == nil {
		generation.StoppingSequences = []string{"\"<unk>\""}
	}
	return generation
}
