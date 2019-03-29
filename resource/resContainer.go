package resource

import (
	"errors"
	"kboard/exception"
)

type IContainer interface {
	SetImage(string) error
	SetName(string) error
	SetImagePullPolicy(string) error
	SetCommands([]string) error
	SetArgs([]string) error
	SetWorkDir(string) error
	SetVolumeMount([]VolumeMount) error
	SetResource(ContainerResources) error
	SetEnv(Env) error
	SetPort(ContainerPort) error
	SetLivenessProbe(Probe) error
	SetReadinessProbe(ReadinessProbe) error
}

// 容器结构体
type Container struct {
	Args                     []string
	Command                  []string
	Env                      []Env
	EnvFrom                  []EnvFromSource `yaml:"envFrom"`
	Name                     string
	Image                    string
	ImagePullPolicy          string `yaml:"imagePullPolicy"` // [Always | Never | IfNotPresent]
	Lifecycle                Lifecycle
	WorkingDir               string        `yaml:"workingDir"`   // 当前工作目录
	VolumeMounts             []VolumeMount `yaml:"volumeMounts"` // 挂载卷
	Resources                ContainerResources
	Ports                    []ContainerPort // 端口号
	LivenessProbe            Probe           `yaml:"livenessProbe"`
	ReadinessProbe           Probe           `yaml:"readinessProbe"`
	Stdin                    bool
	StdinOnce                bool `yaml:"stdinOnce"`
	TerminationMessagePath   string `yaml:"terminationMessagePath"`
	TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
	Tty                      bool
	SecurityContext          SecurityContext `yaml:"securityContext"`
	VolumeDevices            []VolumeDevice  `yaml:"volumeDevices"`
}

func NewContainer(name string, image string) *Container {
	return &Container{
		Name:            name,
		Image:           image,
		Resources:       ContainerResources{},
		LivenessProbe:   Probe{},
		SecurityContext: SecurityContext{},
		Env:             []Env{},
		Ports:           []ContainerPort{},
		Args:            []string{},
		Command:         []string{},
		ImagePullPolicy: "",
		WorkingDir:      "",
		VolumeMounts:    []VolumeMount{},
	}
}

func (r *Container) SetPort(port ContainerPort) error {
	if port == (ContainerPort{}) {
		return errors.New("port is nil")
	}
	r.Ports = append(r.Ports, port)
	return nil
}

func (r *Container) SetVolumeMount(vols []VolumeMount) error {
	if vols == nil {
		return errors.New("volume is nil")
	}
	r.VolumeMounts = vols
	return nil
}

type VolumeMount struct {
	MountPath        string `yaml:"mountPath"`
	MountPropagation string `yaml:"mountPropagation"`
	Name             string
	ReadOnly         bool   `yaml:"readOnly"`
	SubPath          string `yaml:"subPath"`
}

type VolumeDevice struct {
	DevicePath string `yaml:"devicePath"`
	Name       string
}

type SecurityContext struct {
	Privileged               bool   // true-容器运行在特权模式
	AllowPrivilegeEscalation bool   `yaml:"allowPrivilegeEscalation"`
	ProcMount                string `yaml:"procMount"`
	Capabilities             Capabilities
	ReadOnlyRootFilesystem   bool           `yaml:"readOnlyRootFilesystem"`
	RunAsGroup               int            `yaml:"runAsGroup"`
	RunAsNonRoot             bool           `yaml:"runAsNonRoot"`
	RunAsUser                int            `yaml:"runAsUser"`
	SeLinuxOptions           SELinuxOptions `yaml:"seLinuxOptions"`
}

type Capabilities struct {
	Add  []string
	Drop []string
}

type SELinuxOptions struct {
	Level string
	Role  string
	Type  string
	User  string
}

type Lifecycle struct {
	PostStart Handler `yaml:"postStart"`
	PreStop   Handler `yaml:"preStop"`
}

type Handler struct {
	Exec      ExecAction      `yaml:"exec"`
	HttpGet   HttpGetAction   `yaml:"httpGet"`
	TcpSocket TcpSocketAction `yaml:"tcpSocket"`
}

type Env struct {
	Name      string
	ValueFrom *ValueFrom `yaml:"valueFrom"`
}

type ValueFrom struct {
	FieldRef         *FieldRef         `yaml:"fieldRef"`
	ResourceFieldRef *ResourceFieldRef `yaml:"resourceFieldRef"`
}

type FieldRef struct {
	FieldPath string `yaml:"fieldPath"`
}

type ResourceFieldRef struct {
	ContainerName string `yaml:"containerName"`
	Resource      string
}

type ValueFromHandler interface {
}

type EnvFromSource struct {
	ConfigMapRef ConfigMapEnvSource `yaml:"configMapRef"`
	Prefix       string
	SecretRef    SecretEnvSource `yaml:"secretRef"`
}

type ConfigMapEnvSource struct {
	Name     string
	Optional bool
}

type SecretEnvSource struct {
	Name     string
	Optional bool
}

func (r *Container) SetEnv(env Env) error {
	if env == (Env{}) {
		return errors.New("env is nil")
	}
	r.Env = append(r.Env, env)
	return nil
}

func (r *Container) SetArgs(args []string) error {
	if len(args) <= 0 {
		return errors.New("args is empty")
	}
	for _, v := range args {
		if v == "" {
			continue
		}
		r.Args = append(r.Args, v)
	}
	return nil
}

func (r *Container) SetImage(img string) error {
	if img == "" {
		return errors.New("image is empty")
	}
	r.Image = img
	return nil
}

func (r *Container) SetName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Name = name
	return nil
}

func (r *Container) SetImagePullPolicy(imgplc string) error {
	if imgplc == "" {
		return exception.NewError("image pull policy is empty")
	}
	r.ImagePullPolicy = imgplc
	return nil
}

func (r *Container) SetCommands(cmds []string) error {
	if len(cmds) <= 0 {
		return exception.NewError("commands is empty")
	}
	for _, cmd := range cmds {
		if cmd == "" {
			continue
		}
		r.Command = append(r.Command, cmd)
	}
	return nil
}

func (r *Container) SetWorkDir(wkdir string) error {
	if wkdir == "" {
		return exception.NewError("work directory is empty")
	}
	r.WorkingDir = wkdir
	return nil
}

func (r *Container) SetLivenessProbe(liveness Probe) error {
	if liveness.HttpGet != nil {
		// http get 方式检查 port、path
		if liveness.HttpGet.Path == "" || liveness.HttpGet.Port == "" {
			return errors.New("liveness probe use httpget way, need path and port")
		}
	} else if liveness.TcpSocket.Port == 0 {
		return errors.New("tcp socket port way need port")
	}

	if len(liveness.Exec.Command) <= 0 {
		return errors.New("command is empty")
	}
	if liveness.FailureThreshold <= 0 {
		return errors.New("failure threshold is invalid")
	}
	if liveness.InitialDelaySeconds <= 0 {
		return errors.New("initial delay second is invalid")
	}
	if liveness.PeriodSeconds <= 0 {
		return errors.New("period second is invalid")
	}
	if liveness.SuccessThreshold <= 0 {
		return errors.New("success second is invalid")
	}
	if liveness.TimeoutSeconds <= 0 {
		return errors.New("timeout second is invalid")
	}
	r.LivenessProbe = liveness
	return nil
}

func (r *Container) SetReadinessProbe(readiness ReadinessProbe) error {
	if readiness.HttpGet != nil {
		// http get 方式检查 port、path
		if readiness.Path == "" || readiness.Port == 0 {
			return errors.New("liveness probe use httpget way, need path and port")
		}
	} else if readiness.TcpSocket.Port == 0 {
		return errors.New("tcp socket port way need port")
	}

	if readiness.FailureThreshold <= 0 {
		return errors.New("failure threshold is invalid")
	}
	if readiness.InitialDelaySeconds <= 0 {
		return errors.New("initial delay second is invalid")
	}
	if readiness.PeriodSeconds <= 0 {
		return errors.New("period second is invalid")
	}
	if readiness.SuccessThreshold <= 0 {
		return errors.New("success second is invalid")
	}
	if readiness.TimeoutSeconds <= 0 {
		return errors.New("timeout second is invalid")
	}
	return nil
}

func (r *Container) SetResource(res ContainerResources) error {
	if res.Requests.Cpu == "" || res.Requests.Memory == "" {
		return errors.New("request cpu or memory is empty")
	}
	r.Resources = res
	return nil
}

type ProbeAction struct {
	Exec      *ExecAction      `yaml:"exec"`
	HttpGet   *HttpGetAction   `yaml:"httpGet"`
	TcpSocket *TcpSocketAction `yaml:"tcpSocket"`
}

type TcpSocketAction struct {
	Port int
}

type ExecAction struct {
	Command []string
}

type HttpGetAction struct {
	Path        string
	Port        string
	Host        string
	Scheme      string
	HttpHeaders []map[string]string `yaml:"httpHeaders"`
}

type Probe struct {
	ProbeAction         `yaml:",inline"`
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int `yaml:"timeoutSeconds"`
	PeriodSeconds       int `yaml:"periodSeconds"`
	SuccessThreshold    int `yaml:"successThreshold"`
	FailureThreshold    int `yaml:"failureThreshold"`
}

type ReadinessProbe struct {
	ProbeAction
	Path                string
	Port                int
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int `yaml:"timeoutSeconds"`
	PeriodSeconds       int `yaml:"periodSeconds"`
	SuccessThreshold    int `yaml:"successThreshold"`
	FailureThreshold    int `yaml:"failureThreshold"`
}

type Secret struct {
	SecretName string              `yaml:"secretName"`
	Items      []map[string]string // [key:string, path:string]
}

func NewResource() *ContainerResources {
	return &ContainerResources{
		Limits:   Limits{Cpu: "", Memory: ""},
		Requests: Request{Cpu: "", Memory: ""},
	}
}

type ContainerResources struct {
	Limits   Limits
	Requests Request
}

type Port struct {
	Name          string
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	Protocol      string // 仅支持 TCP UDP
}

type ContainerPort struct {
	Name          string
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	Protocol      string // 仅支持 TCP UDP
	HostIP        string `yaml:"hostIP"`
}

func NewPort(name string) *Port {
	return &Port{
		Name:          name,
		ContainerPort: 0,
		HostPort:      0,
		Protocol:      "",
	}
}
