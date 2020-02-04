package agent

import (
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func TestNewMasterDataDirBackupTask(t *testing.T) {

	t.Run("it asks the copying utility to preserves some files", func(t *testing.T) {
		copyUtility := &mockCopyUtility{}

		task := NewMasterDataDirBackupTask(
			copyUtility,
			[]string{"foo.txt", "bar.txt"},
		)

		err := task.Restore("/some/source/dir", "/some/target/dir")
		if err != nil {
			t.Errorf("Restore() returned error %+v", err)
		}

		request := copyUtility.requests[0]

		expectedSourceDir := "/some/source/dir"
		if request.sourceDir != expectedSourceDir {
			t.Errorf("wanted copy utility to recieve source dir as %v, got %v",
				request.sourceDir,
				expectedSourceDir)
		}

		expectedTargetDir := "/some/target/dir"
		if request.targetDir != expectedTargetDir {
			t.Errorf("wanted copy utility to recieve target dir as %v, got %v",
				request.targetDir,
				expectedTargetDir)
		}

		if request.excludedFiles[0] != "foo.txt" {
			t.Errorf("wanted copy utility to recieve excluded file %v in %v",
				"foo.txt",
				request.excludedFiles)
		}

		if request.excludedFiles[1] != "bar.txt" {
			t.Errorf("wanted copy utility to recieve excluded file %v in %v",
				"bar.txt",
				request.excludedFiles)
		}
	})

	t.Run("bubbles up errors from utility", func(t *testing.T) {
		copyUtility := &mockCopyUtility{Err: errors.New("ahhhh")}
		task := NewMasterDataDirBackupTask(copyUtility, []string{})

		err := task.Restore("doesn't", "matter")
		if !xerrors.Is(err, copyUtility.Err) {
			t.Errorf("returned %#v, want %#v", err, copyUtility.Err)
		}
	})
}

type requestData struct {
	sourceDir     string
	targetDir     string
	excludedFiles []string
}

type mockCopyUtility struct {
	Err      error
	requests []requestData
}

func (m *mockCopyUtility) Copy(sourceDir, targetDir string, excludedFiles []string) error {
	m.requests = append(m.requests, requestData{
		sourceDir:     sourceDir,
		targetDir:     targetDir,
		excludedFiles: excludedFiles,
	})

	return m.Err
}
