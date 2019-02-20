package v1beta1

import (
	"k8s-client-go/resource"
	"gopkg.in/yaml.v2"
	"errors"
)

type ICronJob interface {
	resource.IResource
	SetMetadataName(name string) error
	GetMetaDataName() string
	SetNamespace(string) error
}

type ResCronJob struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string
		Namespace string
	}
}

func NewResCronJob() *ResCronJob {
	return &ResCronJob{
		Kind:       resource.RESOURCE_CRON_JOB,
		ApiVersion: "batch/v1beta1",
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
	}
}

func (r *ResCronJob) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func (r *ResCronJob) SetMetadataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResCronJob) GetMetaDataName() string {
	return r.Metadata.Name
}

func (r *ResCronJob) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}
