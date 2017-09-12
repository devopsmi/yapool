package yapool

import "testing"

func TestDb_Add(t *testing.T) {
	db := GetDB()
	db.Add("localhost:9007", Alive)
}

func TestDb_Delete(t *testing.T) {
	db := GetDB()
	db.Delete("localhost:9007")
}

func TestDb_Read(t *testing.T) {
	db := GetDB()
	stat := db.Read("localhost:9007")
	t.Logf("  localhost:9007  status  is  %s ", stat)
}
