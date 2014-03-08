package trash

import (
	"path/filepath"
	"io/ioutil"
	"os"
	"testing"
)

func testDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(pwd, "testdir")
}

func setup(t *testing.T) {
	testDir := testDir()
	err := os.MkdirAll(testDir, 0700)
	if err != nil {
		t.Fatal(err)
	}
}

func teardown(t *testing.T) {
	testDir := testDir()
	os.RemoveAll(testDir)
}

func touch(filePath string) {
	ioutil.WriteFile(filePath, []byte(""), 0700)
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func Test_MoveToTrash(t *testing.T) {
	setup(t)
	defer teardown(t)
	
	filePath := testDir() + "/test"
	touch(filePath)
	
	err := MoveToTrash(filePath)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	
	if fileExists(filePath) {
		t.Error("File was not deleted")
	}
}