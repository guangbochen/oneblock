package mlservice

import (
	"fmt"

	mlv1 "github.com/oneblock-ai/oneblock/pkg/apis/ml.oneblock.ai/v1"
	"github.com/oneblock-ai/oneblock/pkg/utils/constant"
	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServeConfig struct {
	Applications []ServeApplication `yaml:"applications,omitempty"`
}

type ServeApplication struct {
	Name        string    `yaml:"name,omitempty"`
	RoutePrefix string    `yaml:"route_prefix,omitempty"`
	ImportPath  string    `yaml:"import_path,omitempty"`
	Args        ServeArgs `yaml:"args,omitempty"`
}

type ServeArgs struct {
	Models []string `json:"models,omitempty"`
}

func getRayServiceConfig(mlService *mlv1.MLService, modelTmpVersion *mlv1.ModelTemplateVersion,
	owners []metav1.OwnerReference, releaseName string) (*rayv1.RayService, error) {

	serveConfig, err := getServeConfigV2(mlService.Name, getModelConfigPath(modelTmpVersion.Name))
	if err != nil {
		return nil, err
	}

	rayClusterSpec, err := GetRayClusterSpecConfig(mlService, modelTmpVersion, releaseName)
	if err != nil {
		return nil, err
	}

	raySvc := &rayv1.RayService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mlService.Name,
			Namespace: mlService.Namespace,
			Annotations: map[string]string{
				constant.AnnotationRayFTEnabledKey: "true",
			},
			OwnerReferences: owners,
		},
		Spec: rayv1.RayServiceSpec{
			ServeConfigV2:  serveConfig,
			RayClusterSpec: *rayClusterSpec,
		},
	}

	return raySvc, nil
}

func getModelConfigPath(modelName string) string {
	return fmt.Sprintf("./models/%s.yaml", modelName)
}
