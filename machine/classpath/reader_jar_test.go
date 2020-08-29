package classpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJarFileReader(t *testing.T) {
	r, err := NewJarFileReader("../../hello-world.jar")
	if err != nil {
		assert.Error(t, err)
	}
	r.readClass("")
}
