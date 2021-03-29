package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
	cm.CreateConcurrentMap(9)
}