#!/bin/bash
set -euo pipefail

install_gpdb_rpm() {
    local node_hostname=$1 rpm_dir=$2

    scp "${rpm_dir}"/*.rpm "${node_hostname}:/tmp/gpdb_new.rpm"
    ssh -ttn centos@"$node_hostname" '
        sudo rpm -hi /tmp/gpdb_new.rpm
        sudo chown -R gpadmin:gpadmin /usr/local/greenplum-db*
    '
}

./ccp_src/scripts/setup_ssh_to_cluster.sh

for segment_host in $(cat cluster_env_files/hostfile_all); do
    install_gpdb_rpm $segment_host "${RPM_DIR}"
done

