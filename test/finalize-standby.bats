#!/usr/bin/env bats

load helpers

setup_state_dir() {
    STATE_DIR=$(mktemp -d /tmp/gpupgrade.XXXXXX)
    export GPUPGRADE_HOME="${STATE_DIR}/gpupgrade"
}

teardown_new_cluster() {
    if [ -n "$NEW_CLUSTER" ]; then
        delete_cluster $NEW_CLUSTER
    fi
}

setup() {
    skip_if_no_gpdb

    setup_state_dir

    gpupgrade kill-services
}

teardown() {
    teardown_new_cluster
}

#
# This test should probably become a journey test for
# more finalize steps that actually go through the process of
# finalizing a cluster (mirrors, etc)
#
@test "finalize brings up the standby for the new cluster" {
    gpupgrade initialize \
        --old-bindir="$GPHOME/bin" \
        --new-bindir="$GPHOME/bin" \
        --old-port="${PGPORT}" \
        --disk-free-ratio 0 \
        --verbose

    gpupgrade execute --verbose

    gpupgrade finalize

    local standby_status=$(get_standby_status)

    [[ $standby_status == *"Standby host passive"* ]] || fail "expected standby to be up and in passive mode, got ${standby_status}"
}

get_standby_status() {
    local return_value=$(
        gpstate > /tmp/.gpstate-output &&
            cat /tmp/.gpstate-output |
            grep 'Master standby' |
            cut -d '-' -f 3
    )

    echo $return_value
}

