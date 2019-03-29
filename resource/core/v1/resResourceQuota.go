package v1

import (
	"errors"
	"gopkg.in/yaml.v2"
	"k8s-client-go/resource"
)

// 需要启动时配置ResourceQuota adminssion control
type IResResourceQuota interface {
	resource.IResource
	SetMetaDataName(name string) error
	GetNamespace() string
	SetNamespace(string) error
}

type ResResourceQuota struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string
		Namespace string
	}
	Spec struct {
		Hard *Hard
	}
}

type Hard struct {
	Configmaps             string
	Persistentvolumeclaims string
	Replicationcontrollers string
	Secrets                string
	Services               string
	Pods                   string
	Requests               *resource.Request
	Limits                 *resource.Limits
}

func NewResResourceQuota() *ResResourceQuota {
	return &ResResourceQuota{
		Kind:       resource.RESOURCE_RESOURCE_QUOTA,
		ApiVersion: "v1",
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Spec: struct{ Hard *Hard }{Hard: &Hard{
			Requests: &resource.Request{
				Cpu:    "",
				Memory: "",
			},
			Limits: &resource.Limits{
				Cpu:    "",
				Memory: "",
			},
		}},
	}
}

func (r *ResResourceQuota) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResResourceQuota) SetMetaDataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResResourceQuota) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResResourceQuota) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
