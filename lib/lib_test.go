package lib

import (
	"encoding/csv"
	"fmt"
	"github.com/flipfit/constant"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenFile(t *testing.T) {
	wd := GetParentDirectory()
	parent := filepath.Dir(wd)
	fileName := fmt.Sprintf("%s/%s%s", parent, constant.FILE_PATH, constant.USER_FILE_NAME)
	file, err := OpenFile(fileName)
	if err != nil || file == nil {
		t.Errorf("Failed in testing OpenFile Error: %+v", err)
	}
	defer file.Close()
	t.Log("open file test success")
}

func TestWriteFile(t *testing.T) {
	wd := GetParentDirectory()
	parent := filepath.Dir(wd)

	fileName := fmt.Sprintf("%s/%s%s", parent, constant.FILE_PATH, "test.csv")
	file, err := OpenFile(fileName)
	if err != nil || file == nil {
		t.Errorf("Failed in testing OpenFile Error: %+v", err)
	}
	csvWriter := csv.NewWriter(file)
	if err := WriteFile(csvWriter, []string{"test", "test1", "test2"}); err != nil {
		t.Errorf("Failed in testing writing file")
	}

	if err = os.Remove(fileName); err != nil {
		t.Logf("Failed in testing OpenFile Error: %+v", err)
	}

	t.Log("write file test success")
}

func TestGetParentDirectory(t *testing.T) {
	parent := GetParentDirectory()
	if parent == "" {
		t.Errorf("Failed in GetParentDirectory")
	}
	t.Log("GetParentDirectory test success")
}
