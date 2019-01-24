package shenjian

import "testing"

func TestNewShenjian(t *testing.T) {
	shenjian := newShenjian("11111", "11111")
	t.Log(shenjian)
}
