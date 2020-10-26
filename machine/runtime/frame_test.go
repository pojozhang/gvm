package runtime

import (
	"github.com/magiconair/properties/assert"
	"math"
	"testing"
)

func TestOperandStack_Long(t *testing.T) {
	stack := newOperandStack(2)
	stack.PushLong(math.MaxInt64)

	value := stack.PopLong()

	assert.Equal(t, value, int64(math.MaxInt64))
}

func TestOperandStack_Double(t *testing.T) {
	stack := newOperandStack(2)
	stack.PushDouble(1.1)

	value := stack.PopDouble()

	assert.Equal(t, value, 1.1)
}
