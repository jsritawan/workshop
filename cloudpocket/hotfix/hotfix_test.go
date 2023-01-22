package hotfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemicalError(t *testing.T) {
	var num1 float64 = 0.1
	var num2 float64 = 0.2
	var num3 float64 = 0.3

	res := num1 + num2
	assert.Equal(t, num3, res)

}
