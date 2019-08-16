#!/bin/bash
set -euo pipefail

install_gpdb_rpm() {
    local node_hostname=$1 rpm_dir=$2

    scp "${rpm_dir}"/*.rpm "${node_hostname}:/tmp/gpdb_new.rpm"
    ssh -ttn centos@"$node_hostname" '
        sudo rpm -hi /tmp/gpdb_new.rpm
        sudo chown -R gpadmin:gpadmin /usr/local/greenplum-database*
    '
}

./ccp_src/scripts/setup_ssh_to_cluster.sh

# todo: temporarily use the github rpm
rm $RPM_DIR/*
wget https://github.com/greenplum-db/gpdb/releases/download/6.0.0-beta.6/greenplum-database-6.0.0-beta.6-rhel6-x86_64.rpm -P $RPM_DIR

for segment_host in $(cat cluster_env_files/hostfile_all); do
    install_gpdb_rpm $segment_host "${RPM_DIR}"
done

