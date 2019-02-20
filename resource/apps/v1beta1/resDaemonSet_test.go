package v1beta1

import "testing"

func TestNewResDaemonSet(t *testing.T) {
	var daemonSet IResDaemonSet

	daemonSet = NewResDaemonSet()
	daemonSet.SetNamespace("my")
	daemonSet.SetMetaDataName("")
}
