package user

import (
	"context"
	"fmt"

	ctlrbacv1 "github.com/rancher/wrangler/v2/pkg/generated/controllers/rbac/v1"
	"github.com/sirupsen/logrus"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mgmtv1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
	ctlmgmtv1 "github.com/oneblock-ai/oneblock/pkg/generated/controllers/management.oneblock.ai/v1"
	"github.com/oneblock-ai/oneblock/pkg/indexeres"
	"github.com/oneblock-ai/oneblock/pkg/server/config"
)

const (
	usernameLabelKey     = "management.oneblock.ai/username"
	adminRole            = "cluster-admin"
	publicInfoViewerRole = "system:public-info-viewer"
	userControllerName   = "ob-user-controller"
)

// userHandler reconcile the user's clusterRole and clusterRoleBinding
type userHandler struct {
	users                   ctlmgmtv1.UserClient
	clusterRoleBindings     ctlrbacv1.ClusterRoleBindingClient
	clusterRoleBindingCache ctlrbacv1.ClusterRoleBindingCache
}

func Register(ctx context.Context, management *config.Management) error {
	users := management.OneBlockMgmtFactory.Management().V1().User()

	userRBACController := &userHandler{
		users:                   users,
		clusterRoleBindings:     management.RbacFactory.Rbac().V1().ClusterRoleBinding(),
		clusterRoleBindingCache: management.RbacFactory.Rbac().V1().ClusterRoleBinding().Cache(),
	}

	users.OnChange(ctx, userControllerName, userRBACController.OnChanged)
	return nil
}

func (h *userHandler) OnChanged(_ string, user *mgmtv1.User) (*mgmtv1.User, error) {
	if user == nil || user.DeletionTimestamp != nil {
		return user, nil
	}

	roleName := publicInfoViewerRole
	if user.IsAdmin {
		roleName = adminRole
	}

	if err := h.ensureClusterBinding(roleName, user); err != nil {
		return user, err
	}

	return user, nil
}

func (h *userHandler) ensureClusterBinding(roleName string, user *mgmtv1.User) error {
	subject := rbacv1.Subject{
		Kind: "User",
		Name: user.Name,
	}

	// find if there is a clusterRoleBinding with the same role and subject
	key := indexeres.GetCrbKey(roleName, subject)
	crbs, err := h.clusterRoleBindingCache.GetByIndex(indexeres.ClusterRoleBindingNameIndex, key)
	if err != nil {
		return err
	}
	if len(crbs) > 0 {
		logrus.Infof("ClusterRoleBinding with role %v for subject %v already exists", roleName, subject.Name)
		return nil
	}

	logrus.Infof("Creating clusterRoleBinding with role %v for subject %v", roleName, subject.Name)
	_, err = h.clusterRoleBindings.Create(&rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", user.Name),
			Labels: map[string]string{
				usernameLabelKey: user.Name,
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: mgmtv1.SchemeGroupVersion.String(),
					Kind:       "User",
					Name:       user.Name,
					UID:        user.UID,
				},
			},
		},
		Subjects: []rbacv1.Subject{subject},
		RoleRef: rbacv1.RoleRef{
			Kind: "ClusterRole",
			Name: roleName,
		},
	})

	return err
}
