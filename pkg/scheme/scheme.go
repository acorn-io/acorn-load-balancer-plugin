package scheme

import (
	"github.com/rancher/wrangler/pkg/merr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	Scheme         = runtime.NewScheme()
	Codecs         = serializer.NewCodecFactory(Scheme)
	ParameterCodec = runtime.NewParameterCodec(Scheme)
)

func AddToScheme(scheme *runtime.Scheme) error {
	var errs []error
	errs = append(errs, corev1.AddToScheme(scheme))
	return merr.NewErrors(errs...)
}

func init() {
	utilruntime.Must(AddToScheme(Scheme))
}
