package v1

import (
	"errors"
	"gopkg.in/yaml.v2"
	"k8s-client-go/resource"
)

type IResPod interface {
	resource.IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetRestartPolicy(string) error
	SetLabels(map[string]string) error
	AddContainer(resource.IContainer) error
	AddVolume(*resource.Volume) error
	SetAnnotations(map[string]string) error
}

// pod结构体
type ResPod struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name        string
		Namespace   string
		Labels      map[string]string
		Annotations map[string]string
	}
	Spec resource.PodSpec
}

func NewResPod(name string) *ResPod {
	return &ResPod{
		ApiVersion: "v1",
		Kind:       resource.RESOURCE_POD,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{
			Name:        name,
			Namespace:   "",
			Labels:      map[string]string{},
			Annotations: map[string]string{}},
		Spec: resource.PodSpec{
			Containers:       []resource.Container{},
			RestartPolicy:    "",
			NodeSelector:     struct{}{},
			ImagePullSecrets: []resource.LocalObjectReference{},
			HostNetwork:      false,
			Volumes:          []resource.Volume{}},
	}
}

func (r *ResPod) SetMetadataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResPod) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResPod) SetRestartPolicy(policy string) error {
	if policy == "" {
		return errors.New("policy is empty")
	}
	r.Spec.RestartPolicy = policy
	return nil
}

func (r *ResPod) AddContainer(container resource.Container) error {
	r.Spec.Containers = append(r.Spec.Containers, container)
	return nil
}

func (r *ResPod) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return errors.New("labels key or val is empty")
		}

		r.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResPod) AddVolume(vol resource.Volume) error {
	r.Spec.Volumes = append(r.Spec.Volumes, vol)
	return nil
}

func (r *ResPod) SetAnnotations(annos map[string]string) error {
	if len(annos) <= 0 {
		return errors.New("annotations is empty")
	}
	for k, v := range annos {
		if k == "" || v == "" {
			return errors.New("annotation key or val is empty")
		}
		r.Metadata.Annotations[k] = v
	}
	return nil
}

func (r *ResPod) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
