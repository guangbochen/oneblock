package auth

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	mgmtv1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
	"github.com/oneblock-ai/oneblock/pkg/server/config"
)

const (
	annoAdminCreatedKey    = "management.oneblock.ai/admin-created"
	defaultAdminLabelKey   = "management.oneblock.ai/default-admin"
	defaultAdminLabelValue = "true"
	defaultAdminPassword   = "password"
)

var defaultAdminLabel = map[string]string{
	defaultAdminLabelKey: defaultAdminLabelValue,
}

func BootstrapAdminUser(mgmt *config.Management) error {
	ns, err := mgmt.CoreFactory.Core().V1().Namespace().Get(mgmt.Namespace, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// return if default admin already created
	if ns.Annotations != nil && ns.Annotations[annoAdminCreatedKey] == "true" {
		return nil
	}

	set := labels.Set(defaultAdminLabel)
	admins, err := mgmt.OneBlockMgmtFactory.Management().V1().User().List(metav1.ListOptions{LabelSelector: set.String()})
	if err != nil {
		return err
	}

	if len(admins.Items) > 0 {
		logrus.Info("Default admin already exist, skip creating")
		return nil
	}

	// admin user not exist, attempt to create the default admin user
	hash, err := HashPasswordString(defaultAdminPassword)
	if err != nil {
		return err
	}

	user, err := mgmt.OneBlockMgmtFactory.Management().V1().User().Create(&mgmtv1.User{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "user-",
			Labels:       defaultAdminLabel,
		},
		DisplayName: "Default Admin",
		Username:    "admin",
		Password:    hash,
		IsAdmin:     true,
	})
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	_, err = mgmt.RbacFactory.Rbac().V1().ClusterRoleBinding().Create(
		&rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "default-admin-",
				Labels: map[string]string{
					defaultAdminLabelKey: defaultAdminLabelValue,
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
			Subjects: []rbacv1.Subject{
				{
					Kind:     "User",
					APIGroup: rbacv1.GroupName,
					Name:     user.Name,
				},
			},
			RoleRef: rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     "cluster-admin",
			},
		})
	if err != nil {
		return fmt.Errorf("failed to create default admin cluster role binding: %v", err)
	}
	logrus.Info("successfully created default admin user and cluster role binding")

	if ns.Annotations == nil {
		ns.Annotations = make(map[string]string, 1)
	}
	nsCopy := ns.DeepCopy()
	nsCopy.Annotations[annoAdminCreatedKey] = "true"
	if _, err = mgmt.CoreFactory.Core().V1().Namespace().Update(nsCopy); err != nil {
		return err
	}

	return nil
}

func HashPasswordString(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
