#!/bin/bash

set -eux -o pipefail
dirpath=$(dirname "${0}")
source "${dirpath}/../../test/finalize_checks.bash"

# Retrieves the installed GPHOME for a given GPDB RPM.
rpm_gphome() {
    local package_name=$1

    local version=$(ssh -n gpadmin@mdw rpm -q --qf '%{version}' "$package_name")
    echo /usr/local/greenplum-db-$version
}

#
# MAIN
#

# This port is selected by our CI pipeline
MASTER_PORT=5432

# We'll need this to transfer our built binaries over to the cluster hosts.
./ccp_src/scripts/setup_ssh_to_cluster.sh

# Cache our list of hosts to loop over below.
mapfile -t hosts < cluster_env_files/hostfile_all

# Figure out where GPHOMEs are.
export GPHOME_OLD=$(rpm_gphome ${OLD_PACKAGE})
export GPHOME_NEW=$(rpm_gphome ${NEW_PACKAGE})

# Build gpupgrade.
export GOPATH=$PWD/go
export PATH=$GOPATH/bin:$PATH

cd $GOPATH/src/github.com/greenplum-db/gpupgrade
make depend
make

# Install gpupgrade binary onto the cluster machines.
for host in "${hosts[@]}"; do
    scp gpupgrade "gpadmin@$host:/tmp"
    ssh centos@$host "sudo mv /tmp/gpupgrade /usr/local/bin"
done

# Now do the upgrade.
time ssh mdw bash <<EOF
    set -eux -o pipefail

    gpupgrade initialize \
              --target-bindir ${GPHOME_NEW}/bin \
              --source-bindir ${GPHOME_OLD}/bin \
              --source-master-port $MASTER_PORT

    gpupgrade execute
    gpupgrade finalize
EOF

# Test that mirrors and standby actually work
if [[ "${MIRRORS}" = "1" && "${STANDBY}" = "1" ]]; then
    echo 'Doing failover tests of mirrors and standby...'
    validate_mirrors_and_standby "${GPHOME_NEW}" mdw $MASTER_PORT
else
    echo "skipping validate_mirrors_and_standby since the cluster does not have mirrors and a standby"
fi

echo 'Upgrade successful.'
