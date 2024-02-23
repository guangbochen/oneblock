package serve

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	mlv1 "github.com/oneblock-ai/oneblock/pkg/apis/ml.oneblock.ai/v1"
	ctlmlv1 "github.com/oneblock-ai/oneblock/pkg/generated/controllers/ml.oneblock.ai/v1"
	"github.com/oneblock-ai/oneblock/pkg/server/config"
	rayutils "github.com/ray-project/kuberay/ray-operator/controllers/ray/utils"
	"k8s.io/apimachinery/pkg/runtime"
)

type Handler struct {
	scheme *runtime.Scheme
	serve  ctlmlv1.ServeClient
	client http.Client
}

const (
	serveControllerOnChange = "serve.onChange"
)

func Register(ctx context.Context, mgmt *config.Management) error {
	serves := mgmt.OneBlockMLFactory.Ml().V1().Serve()
	h := Handler{
		scheme: mgmt.Scheme,
		serve:  serves,
	}

	serves.OnChange(ctx, serveControllerOnChange, h.OnChanged)
	return nil
}

func (h *Handler) OnChanged(_ string, serve *mlv1.MLServe) (*mlv1.MLServe, error) {
	if serve == nil || serve.DeletionTimestamp != nil {
		return nil, nil
	}

	// convert spec yaml to json
	applications, err := json.Marshal(serve.Spec.Applications)
	if err != nil {
		return serve, err
	}
	fmt.Println(applications)
	clientURL, err := rayutils.FetchHeadServiceURL(ctx, &r.Log, r.Client, rayClusterInstance, common.DashboardAgentListenPortName)
	if err != nil || clientURL == "" {
		return serve, err
	}

	rayDashboardClient := rayutils.GetRayDashboardClient()
	rayDashboardClient.InitClient(clientURL)

	return serve, nil
}
