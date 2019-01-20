#!/bin/bash


# array of background-running precess_ids to be killed on script exit
pids=()
function on_exit {
    echo ">>> killing socialpoll component processes: ${pids[*]}"
    
    for process_id in  ${pids[*]}  ; do
        kill $process_id
        [[ -z $? ]] && echo "killed $process_id"
    done

    exit 0
}

trap on_exit SIGINT SIGTERM

# setup env variables for localhost run
. envconfig/localhost.sh

# here we run all the necessary socialpoll components
nsqlookupd &    
#    pids+=" $!" # nsqlookupd intercepts ^C and closes gracefuly
nsqd --lookupd-tcp-address=localhost:4160 &
#    pids+=" $!" # nsqd intercepts ^C and closes gracefuly
mongod &
#    pids+=" $!" # mongod intercepts ^C and closes gracefuly
twittervotes/twittervotes &
    pids+=" $!"
counter/counter &
    pids+=" $!"
api/api &
    pids+=" $!"
cd web && ./web &
    pids+=" $!"

echo "All running. Check http://localhost:8081/ "

# keep the script from dying
while true; do sleep 1; done
