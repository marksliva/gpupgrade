package commanders

import (
	"context"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
)

func DoAll() error {
	_, err := p.client.DoAll(context.Background(), &pb.DoAllRequest{})
	if err != nil {
		return err
	}

	gplog.Info("Started process to upgrade cluster with DoAll command, check gpupgrade_agent logs for details")
	return nil
}
