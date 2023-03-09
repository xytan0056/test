package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_f1(t *testing.T) {
	assert.True(t, false, "f1")
}

func Test_f2(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 1, 1)
		assert.Equal(t, 1, 2)

		// panic
		var p *int
		assert.Equal(t, 2, *p)


		assert.Equal(t, 2, 2)
	})
}

func Test_f3(t *testing.T) {
	assert.True(t, true, "f1")
}
