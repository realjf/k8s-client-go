package resource

import "testing"

func TestNewContainer(t *testing.T) {
	var container IContainer

	container = NewContainer("nginx", "nginx:latest")

	container.SetArgs([]string{"/bin/sh -c"})
	env := new(Env)
	env.Name = "SYS_ENV"
	env.ValueFrom.ResourceFieldRef.Resource = "1"
	env.ValueFrom.ResourceFieldRef.ContainerName = "container"
	container.SetEnv(*env)
	container.SetCommands([]string{"hello"})
	port := ContainerPort{}
	port.Name = "port"
	container.SetPort(port)
	container.SetVolumeMount(map[string]interface{}{
		"volume1": "1313",
	})

	t.Fatalf("%v", container)
}
