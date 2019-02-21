package v1

import (
	"encoding/base64"
	"k8s-client-go/resource"
	"gopkg.in/yaml.v2"
	"errors"
)

type IResSecret interface {
	resource.IResource
	SetMetaDataName(string) error
	GetMetaDataName() string
	SetType(string) error
	SetData([]map[string]string) error
	GetData() (map[string]string, error)
	SetNamespace(string) error
}

type ResSecret struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	Type       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Data map[string]string
}

func NewSecret() *ResSecret {
	return &ResSecret{
		Kind:       resource.RESOURCE_SECRET,
		ApiVersion: "v1",
		Type:       "Opaque",
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Data: map[string]string{},
	}
}

func (r *ResSecret) SetNamespace(ns string) error {
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResSecret) SetMetaDataName(name string) error {
	r.Metadata.Name = name
	return nil
}

func (r *ResSecret) GetMetaDataName() string {
	return r.Metadata.Name
}

func (r *ResSecret) SetType(typeName string) error {
	r.Type = typeName
	return nil
}

func (r *ResSecret) SetData(data []map[string]string) error {
	if len(data) > 0 {
		for _, v := range data {
			if v["key"] == "" || v["val"] == "" {
				return errors.New("key or val is empty")
			}
			// base64编码存储
			r.Data[v["key"]] = base64.StdEncoding.EncodeToString([]byte(v["val"]))
		}
		return nil
	} else {
		return errors.New("no data will be set")
	}
}

func (r *ResSecret) GetData() (map[string]string, error) {
	data := map[string]string{}
	if len(r.Data) > 0 {
		for k, v := range r.Data {
			// base64解码
			val, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				return map[string]string{}, err
			}
			data[k] = string(val)
		}
	}
	return data, nil
}

func (r *ResSecret) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
