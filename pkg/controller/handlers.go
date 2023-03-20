package controller

import (
	"github.com/acorn-io/baaah/pkg/router"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Handler struct {
	client      kubernetes.Interface
	annotations map[string]string
}

// AddAnnotations adds the arg specified annotations to a Service of type LoadBalancer
func (h Handler) AddAnnotations(req router.Request, resp router.Response) error {
	service := req.Object.(*corev1.Service)

	if service.Spec.Type != corev1.ServiceTypeLoadBalancer {
		return nil
	}

	if service.Annotations == nil {
		service.Annotations = map[string]string{}
	}

	logrus.Infof("Updating service %v with injected annotations", service.Name)
	for key, value := range h.annotations {
		service.Annotations[key] = value
	}
	if err := req.Client.Update(req.Ctx, service); err != nil {
		return err
	}
	return nil
}
