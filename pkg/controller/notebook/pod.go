package notebook

import (
	"strings"
	"time"

	cond "github.com/rancher/wrangler/v2/pkg/condition"
	"github.com/rancher/wrangler/v2/pkg/relatedresource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	mgmtv1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
)

const notebookNameLabel = "notebook-name"

// ReconcileNotebookPodOwners reconciles the owner notebook by pod
func (h *handler) ReconcileNotebookPodOwners(_, _ string, obj runtime.Object) ([]relatedresource.Key, error) {
	if hc, ok := obj.(*corev1.Pod); ok {
		for k, v := range hc.GetLabels() {
			if strings.ToLower(k) == notebookNameLabel {
				return []relatedresource.Key{
					{
						Name:      v,
						Namespace: hc.Namespace,
					},
				}, nil
			}
		}
	}

	return nil, nil
}

// PodCondToNotebookCond Note: this is referred to kubeflow notebook controller
// https://github.com/kubeflow/kubeflow/tree/master/components/notebook-controller
func PodCondToNotebookCond(pCond corev1.PodCondition) mgmtv1.Condition {
	condition := mgmtv1.Condition{}

	if len(pCond.Type) > 0 {
		condition.Type = cond.Cond(pCond.Type)
	}

	if len(pCond.Status) > 0 {
		condition.Status = metav1.ConditionStatus(pCond.Status)
	}

	if len(pCond.Message) > 0 {
		condition.Message = pCond.Message
	}

	if len(pCond.Reason) > 0 {
		condition.Reason = pCond.Reason
	}

	// check if pCond.LastProbeTime is null. If so initialize
	// the field with metav1.Now()
	check := pCond.LastProbeTime.Time.Equal(time.Time{})
	if !check {
		condition.LastUpdateTime = pCond.LastProbeTime
	} else {
		condition.LastUpdateTime = metav1.Now()
	}

	// check if pCond.LastTransitionTime is null. If so initialize
	// the field with metav1.Now()
	check = pCond.LastTransitionTime.Time.Equal(time.Time{})
	if !check {
		condition.LastTransitionTime = pCond.LastTransitionTime
	} else {
		condition.LastTransitionTime = metav1.Now()
	}

	return condition
}
