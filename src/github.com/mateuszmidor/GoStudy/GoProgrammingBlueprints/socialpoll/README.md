# twitter votes

# Design
[design.txt](design.txt)  

# Install GO packages
go get github.com/joeshaw/envdecode
go get github.com/garyburd/go-oauth/oauth
go get github.com/bitly/go-nsq
go get gopkg.in/mgo.v2

# Install NSQ distributed messaging system on localhost
sudo pacman -S yaourt
yaourt nsq

NSQ lookup default ports - important for running nsqd (it connects to nsq lookup):
    TCP: listening on 4160
    HTTP: listening on 4161
Note that our program actually connects to nsq (4150), and nsq connects to nsqlookup for topology info

# Install MongoDB on localhost
sudo pacman -S mongodb
sudo mkdir -p /data/db
sudo chown $USER /data/db

# Run NSQ and MongoDB on localhost
terminal1:
    nsqlookupd
terminal2:
    nsqd --lookupd-tcp-address=localhost:4160
terminal3:
    mongod --dbpath ./db

# Install & run NSQ and MongoDB on remote host using docker in a single step :)
in the folder containing docker-compose.yaml:
    docker-compose up

# Configuration (for localhost, for remote host see: docker-compose.yaml)
The following strings are required as env variables (see "setupenv.sh"):
SP_TWITTER_KEY=
SP_TWITTER_SECRET=
SP_TWITTER_ACCESSTOKEN=
SP_TWITTER_ACCESSSECRET=
SP_MONGODB_ADDR=localhost
SP_NSQD_ADDR=localhost:4150
SP_NSQLOOKUP_ADDR=localhost:4161

# Setup MongoDB initial data on localhost - strings the twittervotes will be looking for
> mongo
> use ballots
> db.polls.insert({"title":"Test poll", "options":["love", "kiss", "hate", "fight"]})

# Track NSQ messages on localhost
> nsq_tail --topic="votes" --lookupd-http-address=localhost:4161

# Run the app
terminal4
    twittervotes
terminal5
    counter