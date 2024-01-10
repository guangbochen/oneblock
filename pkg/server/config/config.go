package config

import (
	"context"

	dashboardapi "github.com/kubernetes/dashboard/src/app/backend/auth/api"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v2/pkg/apply"
	corev1 "github.com/rancher/wrangler/v2/pkg/generated/controllers/core"
	rbacv1 "github.com/rancher/wrangler/v2/pkg/generated/controllers/rbac"
	"github.com/rancher/wrangler/v2/pkg/generic"
	"github.com/rancher/wrangler/v2/pkg/start"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/oneblock-ai/oneblock/pkg/auth"
	obcorev1 "github.com/oneblock-ai/oneblock/pkg/generated/controllers/core.oneblock.ai"
	obmgmtv1 "github.com/oneblock-ai/oneblock/pkg/generated/controllers/management.oneblock.ai"
	nvidiav1 "github.com/oneblock-ai/oneblock/pkg/generated/controllers/nvidia.com"
	kuberayv1 "github.com/oneblock-ai/oneblock/pkg/generated/controllers/ray.io"
)

type Management struct {
	ctx        context.Context
	Namespace  string
	ClientSet  *kubernetes.Clientset
	RestConfig *rest.Config
	Apply      apply.Apply

	OneBlockCoreFactory *obcorev1.Factory
	OneBlockMgmtFactory *obmgmtv1.Factory
	CoreFactory         *corev1.Factory
	RbacFactory         *rbacv1.Factory
	KubeRayFactory      *kuberayv1.Factory
	NvidiaFactory       *nvidiav1.Factory
	TokenManager        dashboardapi.TokenManager

	starters []start.Starter
}

func SetupManagement(ctx context.Context, restConfig *rest.Config, namespace string) (*Management, error) {
	mgmt := &Management{
		ctx:       ctx,
		Namespace: namespace,
	}

	apply, err := apply.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	mgmt.Apply = apply

	mgmt.RestConfig = restConfig
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	mgmt.ClientSet = clientSet

	factory, err := controller.NewSharedControllerFactoryFromConfig(mgmt.RestConfig, Scheme)
	if err != nil {
		return nil, err
	}

	factoryOpts := &generic.FactoryOptions{
		SharedControllerFactory: factory,
	}

	oneblockCore, err := obcorev1.NewFactoryFromConfigWithOptions(restConfig, factoryOpts)
	if err != nil {
		return nil, err
	}
	mgmt.OneBlockCoreFactory = oneblockCore
	mgmt.starters = append(mgmt.starters, oneblockCore)

	oneblockMgmt, err := obmgmtv1.NewFactoryFromConfigWithOptions(restConfig, factoryOpts)
	if err != nil {
		return nil, err
	}
	mgmt.OneBlockMgmtFactory = oneblockMgmt
	mgmt.starters = append(mgmt.starters, oneblockMgmt)

	core, err := corev1.NewFactoryFromConfigWithOptions(restConfig, factoryOpts)
	if err != nil {
		return nil, err
	}
	mgmt.CoreFactory = core
	mgmt.starters = append(mgmt.starters, core)

	rbac, err := rbacv1.NewFactoryFromConfigWithOptions(restConfig, factoryOpts)
	if err != nil {
		return nil, err
	}
	mgmt.RbacFactory = rbac
	mgmt.starters = append(mgmt.starters, rbac)

	kuberay, err := kuberayv1.NewFactoryFromConfigWithOptions(restConfig, factoryOpts)
	if err != nil {
		return nil, err
	}
	mgmt.KubeRayFactory = kuberay
	mgmt.starters = append(mgmt.starters, kuberay)

	nvidia, err := nvidiav1.NewFactoryFromConfigWithOptions(restConfig, factoryOpts)
	if err != nil {
		return nil, err
	}
	mgmt.NvidiaFactory = nvidia
	mgmt.starters = append(mgmt.starters, nvidia)

	mgmt.TokenManager, err = auth.NewJWETokenManager(core.Core().V1().Secret(), namespace)
	if err != nil {
		return nil, err
	}

	return mgmt, nil
}

func (m *Management) Start(threadiness int) error {
	return start.All(m.ctx, threadiness, m.starters...)
}
