package tests

import (
	"test/manager"
	"testing"
)

func TestPetData(t *testing.T) {

	petInfo := manager.GetAllPets()

	if len(petInfo) <= 0 {
		t.Error("not able to got data")
	}
	t.Logf("got data %v", petInfo)
}
func TestPetDataQuery(t *testing.T) {

	res := manager.GetPetById("2")
	t.Logf("got data %v", res)
}
