#!/usr/bin/env bash

trap tearDown SIGINT


function stage() {
    COLOR="\e[96m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$COLOR$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"

    command terraform version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install terraform to run this example" && exit 1

    echo "OK"
}

function run() {
    stage "Running"

    # override plugin under construction install path to local dir
    export TF_CLI_CONFIG_FILE=dev.tfrc

    terraform init
    terraform apply \
        -auto-approve

     echo "Done"
}

function tearDown() {
    stage "Exiting now"
    
    exit 0
}


checkPrerequsites
run
tearDown