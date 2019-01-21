package resource

import "testing"

func TestNewResHorizontalPodAutoscaler(t *testing.T) {
	var hpa *ResHorizontalPodAutoscaler
	hpa = NewResHorizontalPodAutoscaler()

	hpa.SetMetadataName("hpa")
	hpa.SetNamespace("hpa_namespace")
	labels := map[string]string{
		"app": "nginx",
	}
	hpa.SetMatchLabels(labels)
	//yaml, _ := hpa.ToYamlFile()

	t.Fatalf("%v", hpa)
}
