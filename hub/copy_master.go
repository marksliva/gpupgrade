package hub

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"sync"

	"github.com/greenplum-db/gpupgrade/step"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

type Result struct {
	stdout bytes.Buffer
	stderr bytes.Buffer
	err    error
}

func (h *Hub) CopyMasterDataDir(streams step.OutStreams) error {
	// Make sure sourceDir ends with a trailing slash so that rsync will
	// transfer the directory contents and not the directory itself.
	sourceDir := filepath.Clean(h.Target.MasterDataDir()) + string(filepath.Separator)
	destinationDirName := filepath.Join(h.StateDir, "master.bak")

	/*
	 * Copy the directory once per host.
	 *
	 * We don't need to copy the master directory on the master host
	 * If there are primaries on the same host, the hostname will be
	 * added for the corresponding primaries.
	 */
	var wg sync.WaitGroup

	hosts := h.Target.PrimaryHostnames()
	results := make(chan *Result, len(hosts))

	for _, hostname := range hosts {
		hostname := hostname // capture range variable

		wg.Add(1)
		go func() {
			defer wg.Done()

			dest := fmt.Sprintf("%s:%s", hostname, destinationDirName)
			cmd := execCommand("rsync",
				"--archive", "--compress", "--delete", "--stats",
				sourceDir, dest)

			result := Result{}
			cmd.Stdout = &result.stdout
			cmd.Stderr = &result.stderr

			err := cmd.Run()
			if err != nil {
				err = xerrors.Errorf("copying master data directory to host %s: %w", hostname, err)
				result.err = err
			}
			results <- &result
		}()
	}

	wg.Wait()
	close(results)

	var multierr *multierror.Error

	for result := range results {
		if _, err := io.Copy(streams.Stdout(), &result.stdout); err != nil {
			multierr = multierror.Append(multierr, err)
		}

		if _, err := io.Copy(streams.Stderr(), &result.stderr); err != nil {
			multierr = multierror.Append(multierr, err)
		}

		if result.err != nil {
			multierr = multierror.Append(multierr, result.err)
		}
	}

	return multierr.ErrorOrNil()
}
