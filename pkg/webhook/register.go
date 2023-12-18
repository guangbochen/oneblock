package webhook

import (
	"context"
	"fmt"

	"github.com/oneblock-ai/webhook/pkg/server"
	"github.com/oneblock-ai/webhook/pkg/server/admission"
	"k8s.io/client-go/rest"

	"github.com/oneblock-ai/oneblock/pkg/webhook/config"
	"github.com/oneblock-ai/oneblock/pkg/webhook/user"
)

func register(mgmt *config.Management) (validators []admission.Validator, mutators []admission.Mutator) {
	validators = []admission.Validator{
		user.NewValidator(mgmt),
	}

	mutators = []admission.Mutator{
		user.NewMutator(),
	}

	return
}

func Register(ctx context.Context, restConfig *rest.Config, threadiness int, ws *server.WebhookServer) error {
	mgmt, err := config.SetupManagement(ctx, restConfig)
	if err != nil {
		return fmt.Errorf("setup management failed: %w", err)
	}

	validators, mutators := register(mgmt)

	if err := ws.RegisterValidators(validators...); err != nil {
		return fmt.Errorf("register validators failed: %w", err)
	}

	if err := ws.RegisterMutators(mutators...); err != nil {
		return fmt.Errorf("register mutators failed: %w", err)
	}

	if err := mgmt.Start(threadiness); err != nil {
		return fmt.Errorf("start management failed: %w", err)
	}

	return nil
}
