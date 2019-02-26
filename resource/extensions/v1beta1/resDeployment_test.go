package v1beta1

import (
	"testing"
	"k8s-client-go/resource"
)

func TestNewResDeployment(t *testing.T) {
	var deploy IResDeployment
	deploy = NewResDeployment()

	deploy.SetMetadataName("name")
	deploy.SetNamespace("namespace")
	labels := map[string]string{
		"app": "nginx",
	}
	deploy.SetMatchLabels(labels)
	deploy.SetTemplateLabels(labels)
	container := resource.NewContainer("container", "image")
	deploy.AddContainer(container)

	deploy.ToYamlFile()

	t.Fatalf("%v", deploy)
}
