package controller

import (
	"context"

	"github.com/acorn-io/baaah"
	"github.com/tylerslaton/acorn-load-balancer-plugin/pkg/scheme"
	"k8s.io/client-go/kubernetes"
)

type Options struct {
	K8s         kubernetes.Interface
	Annotations map[string]string
}

func Start(ctx context.Context, opt Options) error {
	router, err := baaah.DefaultRouter("load-balancer-controller", scheme.Scheme)
	if err != nil {
		return err
	}

	if err := RegisterRoutes(router, opt.K8s, opt.Annotations); err != nil {
		return err
	}

	return router.Start(ctx)
}
