package v1

import (
	"errors"
	"gopkg.in/yaml.v2"
	"k8s-client-go/resource"
)

type IResConfigMap interface {
	resource.IResource
	SetMetadataName(string) error
	SetData([]map[string]string) error
	SetNamespace(string) error
	GetNamespace() string
}

type ResConfigMap struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Data map[string]string
}

func NewConfigMap() *ResConfigMap {
	return &ResConfigMap{
		ApiVersion: "v1",
		Kind:       resource.RESOURCE_CONFIG_MAP,
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Data: map[string]string{},
	}
}

func (r *ResConfigMap) SetMetadataName(name string) error {
	r.Metadata.Name = name
	return nil
}

func (r *ResConfigMap) SetNamespace(ns string) error {
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResConfigMap) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResConfigMap) SetData(data []map[string]string) error {
	if len(data) > 0 {
		for _, v := range data {
			if v["key"] == "" || v["val"] == "" {
				return errors.New("key or val is empty")
			}
			r.Data[v["key"]] = v["val"]
		}
		return nil
	} else {
		return errors.New("no data will be set")
	}
}

func (r *ResConfigMap) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
