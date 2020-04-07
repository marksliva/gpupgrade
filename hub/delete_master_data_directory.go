package hub

import (
	"os"
	"path/filepath"

	"github.com/greenplum-db/gpupgrade/agent"
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

func DeleteMasterDataDirectory(masterDataDir string) error {
	var mErr *multierror.Error

	_, err := utils.System.Stat(masterDataDir)
	if err != nil {
		if xerrors.Is(err, os.ErrNotExist) {
			return nil
		}
		mErr = multierror.Append(mErr, err)
		return mErr
	}

	for _, fileName := range agent.PostgresFiles {
		filePath := filepath.Join(masterDataDir, fileName)
		_, err = utils.System.Stat(filePath)
		if err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}

	if mErr != nil {
		return mErr
	}

	err = utils.System.RemoveAll(masterDataDir)
	if err != nil {
		mErr = multierror.Append(err)
		return mErr
	}

	return nil
}
