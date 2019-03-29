package resource

type ResPodStatus struct {
	Conditions            PodCondition
	ContainerStatuses     []ContainerStatus `yaml:"containerStatuses"`
	HostIP                string            `yaml:"hostIP"`
	InitContainerStatuses []ContainerStatus `yaml:"initContainerStatuses"`
	Message               string
	NominatedNodeName     string `yaml:"nominatedNodeName"`
	Phase                 string
	PodIP                 string `yaml:"podIP"`
	QosClass              string `yaml:"qosClass"`
	Reason                string
	StartTime             Time `yaml:"startTime"`
}

type PodCondition struct {
	LastProbeTime      Time `yaml:"lastProbeTime"`
	LastTransitionTime Time `yaml:"lastTransitionTime"`
	Message            string
	Reason             string
	Status             string
	Type               string
}
