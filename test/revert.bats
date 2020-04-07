#! /usr/bin/env bats

load helpers

setup() {
    skip_if_no_gpdb

    STATE_DIR=`mktemp -d /tmp/gpupgrade.XXXXXX`
    export GPUPGRADE_HOME="${STATE_DIR}/gpupgrade"
    echo $GPUPGRADE_HOME
}

@test "revert stops the gpupgrade processes" {
    gpupgrade initialize \
        --source-bindir="$GPHOME/bin" \
        --target-bindir="$GPHOME/bin" \
        --source-master-port="${PGPORT}" \
        --temp-port-range 6020-6040 \
        --disk-free-ratio 0 \
        --stop-before-cluster-creation \
        --verbose 3>&-

    if [[ $(process_is_running "[g]pupgrade hub") -eq 0 ]]; then
        # todo: this makes a single host assumption
        echo 'expected hub to be running'
        exit 1
    fi
    if [[ $(process_is_running "[g]pupgrade agent") -eq 0 ]]; then
        echo 'expected agent to be running'
        exit 1
    fi

    gpupgrade revert

    if [[ $(process_is_running "[g]pupgrade hub") -ne 0 ]]; then
        echo 'expected hub to have been stopped'
        exit 1
    fi
    if [[ $(process_is_running "[g]pupgrade agent") -ne 0 ]]; then
        echo 'expected agent to have been stopped'
        exit 1
    fi
}

@test "the target cluster gets deleted" {
    gpupgrade initialize \
        --source-bindir="$GPHOME/bin" \
        --target-bindir="$GPHOME/bin" \
        --source-master-port="${PGPORT}" \
        --temp-port-range 6020-6040 \
        --disk-free-ratio 0 \
        --verbose 3>&-

    # parse config.json for the datadirs
    local target_datadirs=$(jq ".Target.Primaries[].DataDir" "${GPUPGRADE_HOME}/config.json")

    gpupgrade revert

    while read -r datadir; do
        if [ -d $(echo "${datadir}" | tr -d '"') ]; then
            echo "expected datadir ${datadir} to have been deleted"
            exit 1
        fi
    done <<< "${target_datadirs}"
}

@test "the state directory gets deleted" {
    gpupgrade initialize \
        --source-bindir="$GPHOME/bin" \
        --target-bindir="$GPHOME/bin" \
        --source-master-port="${PGPORT}" \
        --temp-port-range 6020-6040 \
        --disk-free-ratio 0 \
        --stop-before-cluster-creation \
        --verbose 3>&-

    gpupgrade revert

    if [ -d $(echo "${GPUPGRADE_HOME}") ]; then
        echo "expected GPUPGRADE_HOME directory ${GPUPGRADE_HOME} to have been deleted"
        exit 1
    fi
}
