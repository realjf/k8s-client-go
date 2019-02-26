package v1

import "k8s-client-go/resource"

type IResServiceAccount interface {
	resource.IResource

}

type ResServiceAccount struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Annotations map[string]string
		Labels map[string]string
	}
}


func NewResServiceAccount() *ResServiceAccount {
	return &ResServiceAccount{
		ApiVersion: "v1",
		Kind: resource.RESOURCE_SERVICE,
		Metadata: struct {
			Name        string
			Namespace   string
			Annotations map[string]string
			Labels      map[string]string
		}{Name: string(""), Namespace: string(""), Annotations: nil, Labels: nil},
	}
}

