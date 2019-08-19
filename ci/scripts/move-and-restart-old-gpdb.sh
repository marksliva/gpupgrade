#!/bin/bash
set -euo pipefail

export MASTER_DATA_DIRECTORY=/data/gpdata/master/gpseg-1
export PGPORT=5432

move_and_update_path() {
	local node_hostname=$1 gphome=$2
    ssh -ttn centos@"$node_hostname" GPHOME="${gphome}" '
		sudo mv /usr/local/greenplum-db-devel ${GPHOME}
		sudo sed -e "s|GPHOME=.*$|GPHOME=$GPHOME|" -i ${GPHOME}/greenplum_path.sh
	'
}

stop_old_cluster() {
    ssh gpadmin@mdw \
			MASTER_DATA_DIRECTORY=$MASTER_DATA_DIRECTORY \
			PGPORT=$PGPORT bash <<"EOF"
		source /usr/local/greenplum-db-devel/greenplum_path.sh
		gpstop -a
EOF
}

start_old_cluster() {
	local gphome=$1
    ssh gpadmin@mdw \
			GPHOME="${GPHOME}" \
			MASTER_DATA_DIRECTORY=$MASTER_DATA_DIRECTORY \
			PGPORT=$PGPORT bash <<"EOF"
		source ${GPHOME}/greenplum_path.sh
		gpstart -a
EOF
}

./ccp_src/scripts/setup_ssh_to_cluster.sh

stop_old_cluster "${GPHOME}"

for segment_host in $(cat cluster_env_files/hostfile_all); do
  move_and_update_path $segment_host "${GPHOME}"
done

start_old_cluster "${GPHOME}"
