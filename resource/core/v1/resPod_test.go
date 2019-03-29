package v1

import (
	"k8s-client-go/resource"
	"testing"
)

func TestNewPod(t *testing.T) {
	var pod IResPod

	pod = NewResPod("pod1")
	pod.SetNamespace("namespace")
	labels := map[string]string{
		"app": "app",
		"val": "val",
	}
	pod.SetLabels(labels)
	container := resource.NewContainer("container1", "image1")
	pod.AddContainer(container)
	volume := resource.Volume{}
	volume.Name = "vol1"
	volume.Secret = &resource.Secret{
		SecretName: "secret1",
		Items:      []map[string]string{},
	}
	pod.AddVolume(&volume)
	pod.SetRestartPolicy("policy1")

	t.Fatalf("%v", pod)
}
