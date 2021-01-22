package filesystem_test

import (
	"testing"

	"github.com/frozentech/filesystem"
)

func TestExist(t *testing.T) {
	success, err := filesystem.Exist("/tmp")

	if success != true {
		t.Errorf("Expecting %v, Got %v", true, success)
	}

	if err != nil {
		t.Errorf("Expecting %v, Got %v", nil, err)
	}

	success, err = filesystem.Exist("/tmpsasdas")

	if success != false {
		t.Errorf("Expecting %v, Got %v", false, success)
	}

	if err != nil {
		t.Errorf("Expecting %v, Got %v", err, nil)
	}
}

func TestMkdir(t *testing.T) {
	if err := filesystem.Mkdir("mkdir"); err != nil {
		t.Errorf("Expecting %v, Got %v", nil, err)
	}

	if err := filesystem.Mkdir("/tmp"); err != nil {
		t.Errorf("Expecting %v, Got %v", nil, err)
	}
}

func TestDelete(t *testing.T) {
	if err := filesystem.Delete("mkdir"); err != nil {
		t.Errorf("Expecting %v, Got %v", nil, err)
	}
}
