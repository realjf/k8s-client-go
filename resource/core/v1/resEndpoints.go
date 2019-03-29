package v1

import "k8s-client-go/resource"

type IResEndpoints interface {
	resource.IResource
}

type ResEndpoints struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Subsets Subset
}

type Subset struct {
	Addresses []SubsetAddr
	Ports     []SubsetPort
}

type SubsetAddr struct {
	Ip string
}

type SubsetPort struct {
	Port int
}

func NewResEndpoints() *ResEndpoints {
	return &ResEndpoints{
		ApiVersion: "v1",
		Kind:       resource.RESOURCE_ENDPOINTS,
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
	}
}
