# export Twitter API strings necessary for interacting with Twitter
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "$DIR"/twitter.sh

# export MongoDB & NSQ addresses
export SP_MONGODB_ADDR=mongodb # this is hostname for mongodb docker container set by docker-compose
export SP_NSQD_ADDR=nsqd:4150
export SP_NSQLOOKUP_ADDR=nsqlookupd:4161
