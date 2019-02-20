package v1alpha1

import (
	"gopkg.in/yaml.v2"
	"k8s-client-go/resource"
)

type IResPodPreset interface {
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
type ResPodPreset struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name        string
		Namespace   string
		Labels      map[string]string
		Annotations map[string]string
	}
	Spec *ResPodPresetSpec
}

type ResPodPresetSpec struct {
	Containers       []resource.IContainer
	RestartPolicy    string              `yaml:"restartPolicy"` // [Always | Never | OnFailure]
	NodeSelector     struct{}            `yaml:"nodeSelector"`
	ImagePullSecrets []map[string]string `yaml:"imagePullSecrets"`
	HostNetwork      bool                `yaml:"hostNetwork"`
	Volumes          []*resource.Volume
}

func NewResPodPreset(name string) *ResPodPreset {
	return &ResPodPreset{
		ApiVersion: "settings.k8s.io/v1alpha1",
		Kind:       resource.RESOURCE_POD_PRESET,
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
		Spec: &ResPodPresetSpec{
			Containers:       []resource.IContainer{},
			RestartPolicy:    "",
			NodeSelector:     struct{}{},
			ImagePullSecrets: []map[string]string{},
			HostNetwork:      false,
			Volumes:          []*resource.Volume{},
		},
	}
}

func (r *ResPodPreset) SetMetadataName(string) error {
	return nil
}

func (r *ResPodPreset) SetNamespace(string) error {
	return nil
}

func (r *ResPodPreset) SetRestartPolicy(string) error {
	return nil
}

func (r *ResPodPreset) SetLabels(map[string]string) error {
	return nil
}

func (r *ResPodPreset) AddContainer(resource.IContainer) error {
	return nil
}

func (r *ResPodPreset) AddVolume(*resource.Volume) error {
	return nil
}

func (r *ResPodPreset) SetAnnotations(map[string]string) error {
	return nil
}

func (r *ResPodPreset) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
