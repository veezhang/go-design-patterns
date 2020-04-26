package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	sg1 := GetInstance()
	sg2 := GetInstance()

	if sg1 != sg2 {
		t.Errorf("The instance is diff [%p %p]", &sg1, &sg2)
	}
}
