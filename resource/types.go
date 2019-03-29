package resource

type Volume struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	VolumeSource `json:",inline" yaml:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}

type VolumeSource struct {
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" yaml:"hostPath" protobuf:"bytes,1,opt,name=hostPath"`

	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty" yaml:"emptyDir" protobuf:"bytes,2,opt,name=emptyDir"`

	Secret *SecretVolumeSource `json:"secret,omitempty" protobuf:"bytes,6,opt,name=secret"`

	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty" protobuf:"bytes,9,opt,name=glusterfs"`

	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" yaml:"persistentVolumeClaim" protobuf:"bytes,10,opt,name=persistentVolumeClaim"`

	// +optional
	RBD *RBDVolumeSource `json:"rbd,omitempty" yaml:"rbd" protobuf:"bytes,11,opt,name=rbd"`

	CephFS *CephFSVolumeSource `json:"cephfs,omitempty" yaml:"cephfs" protobuf:"bytes,14,opt,name=cephfs"`

	DownwardAPI *DownwardAPIVolumeSource `json:"downwardAPI,omitempty" yaml:"downwardAPI" protobuf:"bytes,16,opt,name=downwardAPI"`

	ConfigMap *ConfigMapVolumeSource `json:"configMap,omitempty" yaml:"configMap" protobuf:"bytes,19,opt,name=configMap"`
}

type SecretVolumeSource struct {
}

type EmptyDirVolumeSource struct {
	Medium StorageMedium `json:"medium,omitempty" protobuf:"bytes,1,opt,name=medium,casttype=StorageMedium"`

	SizeLimit *Quantity `json:"sizeLimit,omitempty" protobuf:"bytes,2,opt,name=sizeLimit"`
}

type GlusterfsVolumeSource struct {
	EndpointsName string `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints"`

	Path string `json:"path" protobuf:"bytes,2,opt,name=path"`

	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
}

type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `json:"claimName" protobuf:"bytes,1,opt,name=claimName"`

	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
}

type RBDVolumeSource struct {
	CephMonitors []string              `json:"monitors" protobuf:"bytes,1,rep,name=monitors"`
	RBDImage     string                `json:"image" protobuf:"bytes,2,opt,name=image"`
	FSType       string                `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`
	RBDPool      string                `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool"`
	RadosUser    string                `json:"user,omitempty" protobuf:"bytes,5,opt,name=user"`
	Keyring      string                `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring"`
	SecretRef    *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef"`
	ReadOnly     bool                  `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly"`
}

type CephFSVolumeSource struct {
	Monitors   []string         `json:"monitors" protobuf:"bytes,1,req,name=monitors"`
	Path       string           `json:"path,omitempty" protobuf:"bytes,2,opt,name=path"`
	User       string           `json:"user,omitempty" protobuf:"bytes,3,opt,name=user"`
	SecretFile string           `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile"`
	SecretRef  *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
	ReadOnly   bool             `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"`
}

type DownwardAPIVolumeSource struct {
	Items []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"`

	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode"`
}

type DownwardAPIVolumeFile struct {
	Path string `json:"path" protobuf:"bytes,1,opt,name=path"`

	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" protobuf:"bytes,2,opt,name=fieldRef"`

	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,3,opt,name=resourceFieldRef"`

	Mode *int32 `json:"mode,omitempty" protobuf:"varint,4,opt,name=mode"`
}

type ConfigMapVolumeSource struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Items                []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	DefaultMode          *int32      `json:"defaultMode,omitempty" protobuf:"varint,3,opt,name=defaultMode"`
	Optional             *bool       `json:"optional,omitempty" `
}

type KeyToPath struct {
	Key  string `json:"key" protobuf:"bytes,1,opt,name=key"`
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"`
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,3,opt,name=mode"`
}

type ObjectFieldSelector struct {
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,1,opt,name=apiVersion"`
	FieldPath  string `json:"fieldPath" protobuf:"bytes,2,opt,name=fieldPath"`
}

type ResourceFieldSelector struct {
	ContainerName string   `json:"containerName,omitempty" protobuf:"bytes,1,opt,name=containerName"`
	Resource      string   `json:"resource" protobuf:"bytes,2,opt,name=resource"`
	Divisor       Quantity `json:"divisor,omitempty" protobuf:"bytes,3,opt,name=divisor"`
}

type SecretReference struct {
	Name      string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
}

type LocalObjectReference struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

type HostPathVolumeSource struct {
	Path string `json:"path" protobuf:"bytes,1,opt,name=path"`

	Type *HostPathType `json:"type,omitempty" protobuf:"bytes,2,opt,name=type"`
}

type HostPathType string

type StorageMedium string

type Limits struct {
	Cpu    string
	Memory string
}

type Request struct {
	Cpu    string
	Memory string
}
