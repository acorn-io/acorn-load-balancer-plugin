package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/acorn-io/baaah/pkg/restconfig"
	"github.com/sirupsen/logrus"
	"github.com/tylerslaton/acorn-load-balancer-plugin/pkg/controller"
	"github.com/tylerslaton/acorn-load-balancer-plugin/pkg/scheme"
	"github.com/tylerslaton/acorn-load-balancer-plugin/pkg/version"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var (
	versionFlag = flag.Bool("version", false, "print version")

	// TODO - Determine default
	path = flag.String("path", "annotations.yaml", "path to a yaml file of annotations")
)

func main() {
	flag.Parse()

	fmt.Printf("Version: %s\n", version.Get())
	if *versionFlag {
		return
	}

	logrus.Infof("Using file path %s for annotations", *path)

	annotations, err := parseAnnotations(*path)
	if err != nil {
		logrus.Fatal(err)
	}

	config, err := restconfig.Default()
	if err != nil {
		logrus.Fatal(err)
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	k8s := kubernetes.NewForConfigOrDie(config)

	ctx := signals.SetupSignalHandler()
	if err := controller.Start(ctx, controller.Options{
		K8s:         k8s,
		Annotations: annotations,
	}); err != nil {
		logrus.Fatal(err)
	}
	<-ctx.Done()
	logrus.Fatal(ctx.Err())
}

func parseAnnotations(path string) (map[string]string, error) {
	// Read the file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Create a struct to hold the YAML data
	var annotations map[string]string

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &annotations)
	if err != nil {
		return nil, err
	}

	return annotations, nil
}
