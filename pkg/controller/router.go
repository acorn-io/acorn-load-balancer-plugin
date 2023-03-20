package controller

import (
	"github.com/acorn-io/baaah/pkg/router"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
)

var (
	acornManagedSelector = labels.SelectorFromSet(map[string]string{
		"acorn.io/managed": "true",
	})

	appNameLabel        = "acorn.io/app-name"
	appNamespaceLabel   = "acorn.io/app-namespace"
	servicePublishLabel = "acorn.io/service-publish"
)

func RegisterRoutes(router *router.Router, client kubernetes.Interface, annotations map[string]string) error {
	h := Handler{
		client:      client,
		annotations: annotations,
	}

	managedSelector, err := getAcornPublishedServiceSelector()
	if err != nil {
		return err
	}

	router.Type(&corev1.Service{}).Selector(managedSelector).HandlerFunc(h.AddAnnotations)

	return nil
}

func getAcornPublishedServiceSelector() (labels.Selector, error) {
	r1, err := labels.NewRequirement(appNameLabel, selection.Exists, nil)
	if err != nil {
		return nil, err
	}
	r2, err := labels.NewRequirement(appNamespaceLabel, selection.Exists, nil)
	if err != nil {
		return nil, err
	}
	r3, err := labels.NewRequirement(servicePublishLabel, selection.Equals, []string{"true"})
	if err != nil {
		return nil, err
	}
	acornManagedSelector.Add(*r1, *r2, *r3)
	return acornManagedSelector, nil
}
