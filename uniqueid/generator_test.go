package twlib_uniqueid

import (
	"testing"
)

func TestNextID(t *testing.T) {
	err := InitGenerator(1)
	if err != nil {
		t.Fatalf("TestNextID err:%s\n", err.Error())
	}
	t.Logf("generate id:%d\n", NextID())
}

func TestNextUUID(t *testing.T) {
	err := InitGenerator(1)
	if err != nil {
		t.Fatalf("TestNextUUID err:%s\n", err.Error())
	}
	uuid, _ := NextUUID()
	t.Logf("generate uuid:%s\n", uuid)
}
