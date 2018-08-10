package services

import (
	"context"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	pb "github.com/greenplum-db/gpupgrade/idl"
)

func (h *Hub) DoAll(ctx context.Context, in *pb.DoAllRequest) (*pb.DoAllReply, error) {
	gplog.Info("starting gpupgrade with Do All request")

	go h.ExecuteAllSteps()

	return &pb.DoAllReply{}, nil
}

func (h *Hub) ExecuteAllSteps() {
	// check config

	// check version

	// check seginstall

	// prepare start-agents

	// prepare init-cluster

	// prepare shutdown-clusters
	h.ShutdownClusters()

	// upgrade convert-master
	err := h.ConvertMaster()
	if err != nil {
		gplog.Error(err.Error())
		return
	}

	// upgrade share-oids
	h.shareOidFiles()

	// upgrade convert-primaries
	err := h.ConvertPrimaries()
	if err != nil {
		gplog.Error(err.Error())
		return
	}

	// upgrade validate-start-cluster
	h.startNewCluster()

	// upgrade reconfigure-ports
	err := ReconfigurePorts()
	if err != nil {
		gplog.Error(err.Error())
		return
	}
}
