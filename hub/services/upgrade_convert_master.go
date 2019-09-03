package services

import (
	"bytes"
	"io"
	"os/exec"
	"sync"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/greenplum-db/gpupgrade/hub/upgradestatus"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils/log"
)

// Allow exec.Command to be mocked out by exectest.NewCommand.
var execCommand = exec.Command

func (h *Hub) UpgradeConvertMaster(in *idl.UpgradeConvertMasterRequest, stream idl.CliToHub_UpgradeConvertMasterServer) error {
	gplog.Info("starting %s", upgradestatus.CONVERT_MASTER)

	step, err := h.InitializeStep(upgradestatus.CONVERT_MASTER)
	if err != nil {
		gplog.Error(err.Error())
		return err
	}

	go func() {
		defer log.WritePanics()

		if err := ConvertMaster(stream, &bytes.Buffer{}); err != nil {
			gplog.Error(err.Error())
			step.MarkFailed()
		} else {
			step.MarkComplete()
		}
	}()

	return nil
}

// muxedStream provides io.Writers that wrap both gRPC stream and a parallel
// io.Writer (in case the gRPC stream closes) and safely serialize any
// simultaneous writes.
type muxedStream struct {
	stream idl.CliToHub_UpgradeConvertMasterServer
	writer io.Writer
	mutex  sync.Mutex
}

func newMuxedStream(stream idl.CliToHub_UpgradeConvertMasterServer, writer io.Writer) *muxedStream {
	return &muxedStream{
		stream: stream,
		writer: writer,
	}
}

func (m *muxedStream) NewStreamWriter(cType idl.Chunk_Type) io.Writer {
	return &streamWriter{
		muxedStream: m,
		cType:       cType,
	}
}

type streamWriter struct {
	*muxedStream
	cType idl.Chunk_Type
}

func (w *streamWriter) Write(p []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	n, err := w.writer.Write(p)
	if err != nil {
		return n, err
	}

	if w.stream != nil {
		// Attempt to send the chunk to the client. Since the client may close
		// the connection at any point, errors here are logged and otherwise
		// ignored. After the first send error, no more attempts are made.
		err = w.stream.Send(&idl.Chunk{
			Buffer: p,
			Type:   w.cType,
		})

		if err != nil {
			gplog.Info("halting client stream: %v", err)
			w.stream = nil
		}
	}

	return len(p), nil
}

func ConvertMaster(stream idl.CliToHub_UpgradeConvertMasterServer, out io.Writer) error {
	mux := newMuxedStream(stream, out)
	cmd := execCommand("")

	cmd.Stdout = mux.NewStreamWriter(idl.Chunk_STDOUT)
	cmd.Stderr = mux.NewStreamWriter(idl.Chunk_STDERR)

	return cmd.Run()
}

/*
func (h *Hub) ConvertMaster() error {
	pathToUpgradeWD := utils.MasterPGUpgradeDirectory(h.conf.StateDir)
	err := utils.System.MkdirAll(pathToUpgradeWD, 0700)
	if err != nil {
		return errors.Wrapf(err, "mkdir %s failed", pathToUpgradeWD)
	}

	pgUpgradeCmd := fmt.Sprintf("source %s; cd %s; unset PGHOST; unset PGPORT; "+
		"%s --old-bindir=%s --old-datadir=%s --old-port=%d "+
		"--new-bindir=%s --new-datadir=%s --new-port=%d --mode=dispatcher",
		filepath.Join(h.target.BinDir, "..", "greenplum_path.sh"),
		pathToUpgradeWD,
		filepath.Join(h.target.BinDir, "pg_upgrade"),
		h.source.BinDir,
		h.source.MasterDataDir(),
		h.source.MasterPort(),
		h.target.BinDir,
		h.target.MasterDataDir(),
		h.target.MasterPort())

	gplog.Info("Convert Master upgrade command: %#v", pgUpgradeCmd)

	output, err := h.source.Executor.ExecuteLocalCommand(pgUpgradeCmd)
	if err != nil {
		gplog.Error("pg_upgrade failed to start: %s", output)
		return errors.Wrapf(err, "pg_upgrade on master segment failed")
	}

	return nil
}
*/
