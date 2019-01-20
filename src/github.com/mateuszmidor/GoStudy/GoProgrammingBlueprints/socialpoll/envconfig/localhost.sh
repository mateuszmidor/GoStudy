# export Twitter API strings necessary for interacting with Twitter
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "$DIR"/twitter.sh

# export MongoDB & NSQ addresses
export SP_MONGODB_ADDR=localhost
export SP_NSQD_ADDR=localhost:4150
export SP_NSQLOOKUP_ADDR=localhost:4161
