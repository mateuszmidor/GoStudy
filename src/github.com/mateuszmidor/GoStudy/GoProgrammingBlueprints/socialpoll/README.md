# twitter votes

# Install NSQ distributed messaging system and NSQ GO drivers
sudo pacman -S yaourt
yaourt nsq
go get github.com/bitly/go-nsq

NSQ default ports - important for running nsqd:
    TCP: listening on 4160
    HTTP: listening on 4161

# Install MongoDB software and MongoDB GO drivers
sudo pacman -S mongodb
sudo mkdir -p /data/db
sudo chown $USER /data/db
go get gopkg.in/mgo.v2

# Install GO packages
go get github.com/joeshaw/envdecode
go get github.com/garyburd/go-oauth/oauth

# Run the env
terminal1:
    nsqlookupd
terminal2:
    nsqd --lookupd-tcp-address=localhost:4160
terminal3:
    mongod --dbpath ./db

# Twitter API
The following API strings are required as evn variables:
SP_TWITTER_KEY=
SP_TWITTER_SECRET=
SP_TWITTER_ACCESSTOKEN=
SP_TWITTER_ACCESSSECRET=

# Setup MongoDB initial data
> mongo
> use ballots
> db.polls.insert({"title":"Test poll", "options":["love", "kiss", "hate", "fight"]})

# Track NSQ
> nsq_tail --topic="votes" --lookupd-http-address=localhost:4161