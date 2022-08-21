package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiOnce_Do(t *testing.T) {
	m := MultiOnce{}

	count := 0
	f := func() {
		count++
	}

	m.Append(f)
	m.Append(f)

	m.Do()
	assert.Equal(t, 2, count)

	m.Do()
	assert.Equal(t, 2, count)
}
