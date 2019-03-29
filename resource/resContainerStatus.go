package resource

type ContainerStatus struct {
	ContainerID  string `yaml:"containerID"`
	Image        string
	ImageID      string         `yaml:"imageID"`
	LastState    ContainerState `yaml:"lastState"`
	Name         string
	Ready        bool
	RestartCount int `yaml:"restartCount"`
	State        ContainerState
}

type ContainerState struct {
	Running    ContainerStateRunning
	Terminated ContainerStateTerminated
	Waiting    ContainerStateWaiting
}

type ContainerStateRunning struct {
	StartedAt Time `yaml:"startedAt"`
}

type Time struct {
}

type ContainerStateTerminated struct {
	ContainerID string `yaml:"containerID"`
	ExitCode    int    `yaml:"exitCode"`
	FinishedAt  Time   `yaml:"finishedAt"`
	Message     string
	Reason      string
	Signal      int
	StartedAt   Time `yaml:"startedAt"`
}

type ContainerStateWaiting struct {
	Message string
	Reason  string
}
