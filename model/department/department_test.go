package department

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	log.Println(TypeDepartmentMin)
	log.Println(TypeCity)
}

func TestDepartment_Add(t *testing.T) {
	depart := &Department{
		Name:         "网络",
		Address:      "桂林",
		Type:         1,
		Owner:        "富强",
		OwnerContact: "13800138000",
		Admin:        "民主",
		AdminContact: "13800138010",
	}

	if err := depart.Add(); err != nil {
		t.Error(err)
	}
}

func TestDepartment_List(t *testing.T) {
	depart := New()

	departs, count, err := depart.List(0, 20)
	if err != nil {
		t.Error(err)
	}

	log.Printf("count: %d, value: %#v", count, departs)
}
