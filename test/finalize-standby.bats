#!/usr/bin/env bats

load helpers

setup() {
    skip_if_no_gpdb

    gpupgrade kill-services
}

@test "finalize brings up the standby for the new cluster" {
    gpupgrade initialize \
        --old-bindir="$GPHOME/bin" \
        --new-bindir="$GPHOME/bin" \
        --old-port="${PGPORT}" \
        --disk-free-ratio 0 \
        --verbose

    gpupgrade execute --verbose

    gpupgrade finalize --verbose


}
