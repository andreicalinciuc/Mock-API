package repository

import (
	"github.com/andreicalinciuc/mock-api/model"
	"testing"
)

func TestDataRepo_Update(t *testing.T) {
	r := NewData("useru")
	err := r.Update(nil, model.Data{
		Id:      12,
		Payload: map[string]interface{}{"test": "test"},
	})
	if err != nil {
		t.Errorf("update %s", err)
	}
}

func TestDataRepo_Create(t *testing.T) {
	r := NewData("useru")
	err := r.Create(nil, []model.Data{
		{Id: 132, Payload: map[string]interface{}{"test": "test"}},
		{Id: 185, Payload: map[string]interface{}{"da": "d43243a"}},
	})

	if err != nil {
		t.Errorf("create %s", err)
	}
}

func TestDataRepo_Delete(t *testing.T) {
	r := NewData("user")
	err := r.Delete(nil, 185)
	if err != nil {
		t.Errorf("delete %s", err)
	}
}

func TestDataRepo_GetById(t *testing.T) {
	r := NewData("user")
	data, err := r.GetById(nil, 185)
	if err != nil {
		t.Errorf("get by id %s", err.Error())
	}
	t.Logf("%v", data)
}

func TestDataRepo_GetFile(t *testing.T) {
	r := NewData("user")
	data, err := r.FindAll(nil)
	if err != nil {
		t.Errorf("get file %s", err.Error())
	}
	t.Logf("%v", data)
}
