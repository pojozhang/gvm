package conversion

import (
	"github.com/magiconair/properties/assert"
	"gvm/machine/runtime"
	"testing"
)

func TestD2I(t *testing.T) {
	frame := runtime.NewFrame(nil, 1, 2)
	stack := frame.GetOperandStack()
	stack.PushDouble(1.1)
	i := &D2I{}

	i.Execute(frame)

	assert.Equal(t, stack.PopInt(), int32(1))
}
