package mlservice

import (
	"fmt"

	mlv1 "github.com/oneblock-ai/oneblock/pkg/apis/ml.oneblock.ai/v1"
	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const MLServiceKind = "MLService"

func (h *Handler) syncRayServiceStatus(_ string, rayService *rayv1.RayService) (*rayv1.RayService, error) {
	if rayService == nil || rayService.DeletionTimestamp != nil {
		return nil, nil
	}

	if rayService.OwnerReferences == nil || len(rayService.OwnerReferences) == 0 {
		return nil, nil
	}

	ownerRef := getMLServiceOwner(rayService.OwnerReferences)
	if ownerRef == nil {
		return nil, nil
	}

	mlService, err := h.mlServiceCache.Get(rayService.Namespace, ownerRef.Name)
	if err != nil {
		return rayService, err
	}

	if mlService.Status.RayServiceStatuses.ServiceStatus != rayService.Status.ServiceStatus {
		mlServiceCpy := mlService.DeepCopy()
		mlServiceCpy.Status.RayServiceStatuses = rayService.Status
		if rayService.Status.ServiceStatus == rayv1.Running {
			mlv1.MLServiceReady.True(mlServiceCpy)
		} else {
			mlv1.MLServiceReady.False(mlServiceCpy)
			mlv1.MLServicePending.True(mlServiceCpy)
			mlv1.MLServicePending.Reason(mlServiceCpy, string(rayService.Status.ServiceStatus))
		}
		fmt.Println("updating ray service")
		if _, err = h.mlService.UpdateStatus(mlServiceCpy); err != nil {
			return rayService, err
		}
	}

	return nil, nil
}

func getMLServiceOwner(ownerRefers []v1.OwnerReference) *v1.OwnerReference {
	for _, owner := range ownerRefers {
		if owner.Kind == MLServiceKind {
			return &owner
		}
	}

	return nil
}
