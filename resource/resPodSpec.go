package resource

type PodSpec struct {
	ActiveDeadlineSeconds         int `yaml:"activeDeadlineSeconds"`
	Affinity                      Affinity
	AutomountServiceAccountToken  bool `yaml:"automountServiceAccountToken"`
	Containers                    []Container
	DnsConfig                     PodDNSConfig `yaml:"podDNSConfig"`
	DnsPolicy                     string       `yaml:"dnsPolicy"`
	EnableServiceLinks            bool         `yaml:"enableServiceLinks"`
	HostAliases                   []HostAlias  `yaml:"hostAlias"`
	HostIPC                       bool         `yaml:"hostIPC"`
	HostNetwork                   bool         `yaml:"hostNetwork"`
	HostPID                       bool         `yaml:"hostPID"`
	Hostname                      string
	ImagePullSecrets              []LocalObjectReference `yaml:"imagePullSecrets"`
	InitContainers                []Container            `yaml:"initContainers"`
	NodeName                      string                 `yaml:"nodeName"`
	NodeSelector                  struct{}               `yaml:"nodeSelector"`
	Priority                      int
	PriorityClassName             string             `yaml:"priorityClassName"`
	ReadinessGates                []PodReadinessGate `yaml:"readinessGates"`
	RestartPolicy                 string             `yaml:"restartPolicy"` // [Always | Never | OnFailure]
	RuntimeClassName              string             `yaml:"runtimeClassName"`
	SchedulerName                 string             `yaml:"schedulerName"`
	SecurityContext               PodSecurityContext `yaml:"securityContext"`
	ServiceAccount                string             `yaml:"serviceAccount"`
	ServiceAccountName            string             `yaml:"serviceAccountName"`
	ShareProcessNamespace         bool               `yaml:"shareProcessNamespace"`
	Subdomain                     string
	TerminationGracePeriodSeconds int `yaml:"terminationGracePeriodSeconds"`
	Tolerations                   []Toleration
	Volumes                       []Volume
}

type Toleration struct {
	Effect            string
	Key               string
	Operator          string
	TolerationSeconds int `yaml:"tolerationSeconds"`
	Value             string
}

type PodSecurityContext struct {
	FsGroup            int            `yaml:"fsGroup"`
	RunAsGroup         int            `yaml:"runAsGroup"`
	RunAsNonRoot       bool           `yaml:"runAsNonRoot"`
	RunAsUser          int            `yaml:"runAsUser"`
	SeLinuxOptions     SELinuxOptions `yaml:"seLinuxOptions"`
	SupplementalGroups []int          `yaml:"supplementalGroups"`
	Sysctls            []Sysctl
}

type Sysctl struct {
	Name  string
	Value string
}

type PodReadinessGate struct {
	ConditionType string `yaml:"conditionType"`
}

type HostAlias struct {
	Hostnames []string
	Ip        string
}

type PodDNSConfig struct {
	Nameservers []string
	Options     []PodDNSConfigOption
	Searches    []string
}

type PodDNSConfigOption struct {
	Name  string
	Value string
}

type Affinity struct {
	NodeAffinity    NodeAffinity    `yaml:"nodeAffinity"`
	PodAffinity     PodAffinity     `yaml:"podAffinity"`
	PodAntiAffinity PodAntiAffinity `yaml:"podAntiAffinity"`
}

type NodeAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm `yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution  NodeSelector              `yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

type PreferredSchedulingTerm struct {
	Preference NodeSelectorTerm
	Weight     int
}

type NodeSelector struct {
	NodeSelectorTerms []NodeSelectorTerm `yaml:"nodeSelectorTerms"`
}

type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `yaml:"matchExpressions"`
	MatchFields      []NodeSelectorRequirement `yaml:"matchFields"`
}

type NodeSelectorRequirement struct {
	Key      string
	Operator string
	Values   []string
}

type PodAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

type WeightedPodAffinityTerm struct {
	PodAffinityTerm PodAffinityTerm `yaml:"podAffinityTerm"`
	Weight          int
}

type PodAffinityTerm struct {
	LabelSelector LabelSelector `yaml:"labelSelector"`
	Namespaces    []string
	TopologyKey   string `yaml:"topologyKey"`
}

type LabelSelector struct {
	MatchExpressions []LabelSelectorRequirement `yaml:"matchExpressions"`
	MatchLabels      struct{}                   `yaml:"matchLabels"`
}

type LabelSelectorRequirement struct {
	Key      string
	Operator string
	Values   []string
}

type PodAntiAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}
