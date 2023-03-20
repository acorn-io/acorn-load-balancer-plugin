package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

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
	path        = flag.String("path", "", "path to a yaml file of annotations")
	annotations = flag.String("annotations", "", "annotations in the form of key=value,foo=bar")
)

func main() {
	flag.Parse()

	fmt.Printf("Version: %s\n", version.Get())
	if *versionFlag {
		return
	}

	loadBalancerAnnotations := make(map[string]string)
	if *path != "" {
		logrus.Infof("Reading file path %s for annotations", *path)
		err := addFileAnnotations(*path, loadBalancerAnnotations)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	if *annotations != "" {
		err := addFlagAnnotations(*annotations, loadBalancerAnnotations)
		if err != nil {
			logrus.Fatal(err)
		}

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
		Annotations: loadBalancerAnnotations,
	}); err != nil {
		logrus.Fatal(err)
	}
	<-ctx.Done()
	logrus.Fatal(ctx.Err())
}

func addFileAnnotations(path string, annotations map[string]string) error {
	// Read the file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Create a struct to hold the YAML data
	var newAnnotations map[string]string

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &annotations)
	if err != nil {
		return err
	}

	for k, v := range newAnnotations {
		annotations[k] = v
	}

	return nil
}

func addFlagAnnotations(flag string, annotations map[string]string) error {
	for _, pair := range strings.Split(flag, ",") {
		parsed := strings.Split(pair, "=")
		if len(parsed) != 2 {
			return fmt.Errorf("specified annotations are invalid")
		}
		annotations[parsed[0]] = parsed[1]
	}
	return nil
}
