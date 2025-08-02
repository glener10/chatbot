package matematica

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Errorf("A soma estava incorreta, esperava %d, mas obteve %d", 10, total)
	}
	assert.Equal(t, 10, total, "A soma deve ser 10")
}
