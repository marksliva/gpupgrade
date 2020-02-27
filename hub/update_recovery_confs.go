package hub

import (
	"context"
	"sync"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils"
	"github.com/hashicorp/go-multierror"
)

func UpdateRecoveryConfs(ctx context.Context, agentConns []*Connection, sourceCluster *utils.Cluster, targetCluster *utils.Cluster, initializeConfig InitializeConfig) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(agentConns))

	for _, conn := range agentConns {
		wg.Add(1)

		go func(conn *Connection) {
			defer wg.Done()

			mirrors := utils.FilterSegmentsOnHost(initializeConfig.Mirrors, conn.Hostname)
			if len(mirrors) == 0 {
				return
			}

			var confInfos []*idl.RecoveryConfInfo
			for _, mirror := range mirrors {
				confInfos = append(confInfos, &idl.RecoveryConfInfo{
					TemporaryPort: int32(targetCluster.Primaries[mirror.ContentID].Port),
					SourcePort:    int32(sourceCluster.Primaries[mirror.ContentID].Port),
					DataDir:       mirror.DataDir,
				})
			}

			_, err := conn.AgentClient.UpdateRecoveryConfs(ctx, &idl.UpdateRecoveryConfsRequest{RecoveryConfInfos: confInfos})
			if err != nil {
				errChan <- err
			}
		}(conn)
	}
	wg.Wait()
	close(errChan)

	var mErr multierror.Error
	for err := range errChan {
		if err != nil {
			mErr = *multierror.Append(&mErr, err)
		}
	}

	return mErr.ErrorOrNil()
}
