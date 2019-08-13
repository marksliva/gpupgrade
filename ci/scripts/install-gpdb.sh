#!/bin/bash
set -euo pipefail

extract_gpdb_tarball() {
    local node_hostname=$1 tarball_dir=$2

    scp "${tarball_dir}"/*.tar.gz "${node_hostname}:/tmp/gpdb_binary_new.tar.gz"
    ssh -ttn centos@"$node_hostname" '
        sudo mkdir -p /usr/local/greenplum-db-devel-new
        sudo tar -xf /tmp/gpdb_binary_new.tar.gz -C /usr/local/greenplum-db-devel-new
        sudo chown -R gpadmin:gpadmin /usr/local/greenplum-db-devel-new
    '
}

./ccp_src/scripts/setup_ssh_to_cluster.sh

for segment_host in $(cat cluster_env_files/hostfile_all); do
  extract_gpdb_tarball $segment_host "${TARBALL_DIR}"
done

