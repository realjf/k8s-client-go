package v1

import (
	"errors"
	"k8s-client-go/resource"
	"gopkg.in/yaml.v2"
)

type IResStatefulSet interface {
	resource.IResource
	SetMetaDataName(string) error
	SetNamespace(string) error
	SetServiceName(string) error
	SetReplicas(int) error
	SetLabels(map[string]string) error
	AddContainer(resource.IContainer) error
	SetAnnotations(map[string]string) error
	SetVolumeClaimName(string) error
	SetAccessMode(string) error
	SetStorage(string) error
	SetSelector(resource.Selector) error
}

type ResStatefulSet struct {
	Kind       string `yaml:"kind"`
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}
	Spec *StatefulSetSpec
}

type StatefulSetSpec struct {
	Selector            resource.Selector
	ServiceName         string `yaml:"serviceName"`
	Replicas            int
	Template            *StatefulSetSpecTemplate
	VolumeClaimTemplate *VolumeClaimTemplate `yaml:"volumeClaimTemplate"`
}

type StatefulSetSpecTemplate struct {
	Metadata struct {
		Labels map[string]string
	}
	Spec struct {
		Containers []resource.IContainer
	}
}

type VolumeClaimTemplate struct {
	Metadata struct {
		Name        string
		Annotations map[string]string
	}
	Spec *VolumeClaimTemplateSpec
}

type VolumeClaimTemplateSpec struct {
	AccessModes []string `yaml:"accessModes"`
	Resources   struct {
		Requests struct {
			Storage string
		}
	}
}

func NewResStatefulSet() *ResStatefulSet {
	return &ResStatefulSet{
		Kind:       resource.RESOURCE_STATEFULE_SET,
		ApiVersion: "apps/v1",
		Spec: &StatefulSetSpec{
			ServiceName: "",
			Replicas:    0,
			Template: &StatefulSetSpecTemplate{
				Metadata: struct{ Labels map[string]string }{Labels: map[string]string{}},
				Spec:     struct{ Containers []resource.IContainer }{Containers: nil},
			},
			VolumeClaimTemplate: &VolumeClaimTemplate{
				Metadata: struct {
					Name        string
					Annotations map[string]string
				}{
					Name:        "",
					Annotations: map[string]string{}},
				Spec: &VolumeClaimTemplateSpec{
					AccessModes: []string{},
					Resources: struct{ Requests struct{ Storage string } }{
						Requests: struct{ Storage string }{
							Storage: ""}},
				},
			},
		},
	}
}

func (r *ResStatefulSet) SetMetaDataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResStatefulSet) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResStatefulSet) SetServiceName(svcName string) error {
	if svcName == "" {
		return errors.New("service name is empty")
	}
	r.Spec.ServiceName = svcName
	return nil
}

func (r *ResStatefulSet) SetReplicas(replica int) error {
	if replica <= 0 {
		return errors.New("replicas are empty")
	}
	r.Spec.Replicas = replica
	return nil
}

func (r *ResStatefulSet) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return errors.New("labels key or value is empty")
		}
		r.Spec.Template.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResStatefulSet) AddContainer(container resource.IContainer) error {
	if container == nil {
		return errors.New("container is nil")
	}
	r.Spec.Template.Spec.Containers = append(r.Spec.Template.Spec.Containers, container)
	return nil
}

func (r *ResStatefulSet) SetAnnotations(annos map[string]string) error {
	if len(annos) <= 0 {
		return errors.New("annotation is empty")
	}
	for k, v := range annos {
		if k == "" || v == "" {
			return errors.New("annotation's key or value is empty")
		}
		r.Spec.VolumeClaimTemplate.Metadata.Annotations[k] = v
	}

	return nil
}

func (r *ResStatefulSet) SetVolumeClaimName(volClaimName string) error {
	if volClaimName == "" {
		return errors.New("volume claim name is empty")
	}
	r.Spec.VolumeClaimTemplate.Metadata.Name = volClaimName
	return nil
}

func (r *ResStatefulSet) SetAccessMode(accMode string) error {
	if accMode == "" {
		return errors.New("access mode is empty")
	}
	r.Spec.VolumeClaimTemplate.Spec.AccessModes = append(r.Spec.VolumeClaimTemplate.Spec.AccessModes, accMode)
	return nil
}

func (r *ResStatefulSet) SetStorage(cap string) error {
	if cap == "" {
		return errors.New("storage is empty")
	}
	r.Spec.VolumeClaimTemplate.Spec.Resources.Requests.Storage = cap
	return nil
}

func (r *ResStatefulSet) SetSelector(selector resource.Selector) error {
	if selector.MatchExpressions == nil && len(selector.MatchLabels) <= 0 {
		return errors.New("selector is nil")
	}
	r.Spec.Selector = selector
	return nil
}

func (r *ResStatefulSet) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
