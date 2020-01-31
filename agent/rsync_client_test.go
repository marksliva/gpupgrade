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

func TestRsyncClient(t *testing.T) {
	t.Run("it copies data from a source directory to a target directory", func(t *testing.T) {
		sourceDir := getTempDir(t)
		targetDir := getTempDir(t)

		// Cleanup
		defer os.RemoveAll(sourceDir)
		defer os.RemoveAll(targetDir)

		writeToFile(sourceDir+"/hi", []byte("hi"), t)

		client := NewRsyncClient()
		client.Copy(sourceDir, targetDir)

		targetContents, _ := ioutil.ReadFile(targetDir + "/hi")

		if bytes.Compare(targetContents, []byte("hi")) != 0 {
			t.Errorf("target directory file 'hi' contained %v, wanted %v",
				targetContents,
				"hi")
		}
	})

	t.Run("it removes files that existed in the target directory before the sync", func(t *testing.T) {
		sourceDir := getTempDir(t)
		targetDir := getTempDir(t)

		// Cleanup
		defer os.RemoveAll(sourceDir)
		defer os.RemoveAll(targetDir)

		writeToFile(targetDir+"/file-that-should-get-removed", []byte("goodbye"), t)

		client := NewRsyncClient()
		client.Copy(sourceDir, targetDir)

		targetContents, _ := ioutil.ReadFile(targetDir + "/file-that-should-get-removed")

		if bytes.Compare(targetContents, []byte("")) != 0 {
			t.Errorf("target directory file 'file-that-should-get-removed' should not exist, but contains %v",
				string(targetContents))
		}
	})

}
