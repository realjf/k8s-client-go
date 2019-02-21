package v1

import "k8s-client-go/resource"

type IResService interface {
	resource.IResource

}

type ResService struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Annotations map[string]string
	}
	Spec ServiceSpec
}

type ServiceSpec struct {
	Selector resource.Selector
	Ports []ServicePort
	ClusterIP string `yaml:"clusterIP"`
	LoadBalancerIP string `yaml:"loadBalancerIP"`
	Type string
}

type ServicePort struct {
	Protocol string // TCP|UDP
	Port     int
	TargetPort int  `yaml:"targetPort"`
}

func NewResService() *ResService {
	return &ResService{
		ApiVersion: "v1",
		Kind: resource.RESOURCE_SERVICE,
		Metadata: struct {
			Name      string
			Namespace string
			Annotations map[string]string
		}{Name: "", Namespace: "", Annotations: map[string]string{}},
	}
}
