package agent

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func getTempDir(t *testing.T) string {
	sourceDir, err := ioutil.TempDir("", "rsync-source")
	if err != nil {
		t.Fatalf("creating temporary directory: %+v", err)
	}

	return sourceDir
}

func writeToFile(filepath string, contents []byte, t *testing.T) {
	err := ioutil.WriteFile(filepath, contents, 0644)

	if err != nil {
		t.Fatalf("error writing file '%v'", filepath)
	}
}

func TestMasterDataDirBackupTask(t *testing.T) {
	t.Run("it copies data from a source directory to a target directory", func(t *testing.T) {
		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		writeToFile(sourceDir+"/hi", []byte("hi"), t)

		client := NewMasterDataDirBackupTask()
		client.Restore(sourceDir, targetDir)

		targetContents, _ := ioutil.ReadFile(targetDir + "/hi")

		if bytes.Compare(targetContents, []byte("hi")) != 0 {
			t.Errorf("target directory file 'hi' contained %v, wanted %v",
				targetContents,
				"hi")
		}
	})

	t.Run("it removes files that existed in the target directory before the sync", func(t *testing.T) {
		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		writeToFile(targetDir+"/file-that-should-get-removed", []byte("goodbye"), t)

		client := NewMasterDataDirBackupTask()
		client.Restore(sourceDir, targetDir)

		targetContents, _ := ioutil.ReadFile(targetDir + "/file-that-should-get-removed")

		// XXX this checks that the file is either empty or does not exist; we
		// should just check for existence
		if bytes.Compare(targetContents, []byte("")) != 0 {
			t.Errorf("target directory file 'file-that-should-get-removed' should not exist, but contains %v",
				string(targetContents))
		}
	})

	// TODO: port functionality checks from agent/copy_master_test.go

	t.Run("returns underlying copy errors", func(t *testing.T) {
		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		client := NewMasterDataDirBackupTask()
		err := client.Restore("/does/not/exist", targetDir)

		// XXX currently the error that comes back is heavily
		// implementation-dependent; I'm choosing not to implement a sentinel at
		// the moment.
		if err == nil {
			t.Errorf("returned nil error")
		}
	})
}
