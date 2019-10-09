#! /usr/bin/env bats

load helpers

setup() {
    skip_if_no_gpdb

    STATE_DIR=`mktemp -d`
    export GPUPGRADE_HOME="${STATE_DIR}/gpupgrade"

    gpupgrade kill-services

    # XXX We use $PWD here instead of a real binary directory because
    # `make check` is expected to test the locally built binaries, not the
    # installation. This causes problems for tests that need to call GPDB
    # executables...
    gpupgrade initialize \
        --old-bindir="$PWD" \
        --new-bindir="$PWD" \
        --old-port="${PGPORT}" 3>&-
}

teardown() {
    # XXX Beware, BATS_TEST_SKIPPED is not a documented export.
    if [ -z "${BATS_TEST_SKIPPED}" ]; then
        gpupgrade kill-services
        rm -r "$STATE_DIR"
    fi
}

@test "gpupgrade stop-services actually stops hub and agents" {
    # check that hub and agent are up
    run is_process_running "[g]pupgrade_hub"
    [ "$status" -eq 0 ]
    run is_process_running "[g]pupgrade_agent"
    [ "$status" -eq 0 ]

    # stop them
    run gpupgrade kill-services
    [ "$status" -eq 0 ]

    # make sure that they are down
    # check that hub and agent are up
    run is_process_running "[g]pupgrade_hub"
    [ "$status" -eq 1 ]
    run is_process_running "[g]pupgrade_agent"
    [ "$status" -eq 1 ]

    run gpupgrade kill-services
    [ "$status" -eq 0 ]
}

@test "gpupgrade stop-services can be run multiple times without issue " {

    run gpupgrade kill-services
    [ "$status" -eq 0 ]

    run gpupgrade kill-services
    [ "$status" -eq 0 ]
}