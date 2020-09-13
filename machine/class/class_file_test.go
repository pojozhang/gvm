package class

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseClassFile(t *testing.T) {
	r, _ := NewJarFileReader("../../hello-world.jar")
	bytes, _ := r.readClass("gvm/example/Main.class")
	file := &ClassFile{}
	err := file.read(&ClassFileReader{data: bytes})
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, uint32(0xCAFEBABE), file.magic)
	assert.Equal(t, uint16(52), file.majorVersion)
	assert.Equal(t, uint16(0), file.minorVersion)
	assert.Equal(t, uint16(5), file.thisClass)
}
