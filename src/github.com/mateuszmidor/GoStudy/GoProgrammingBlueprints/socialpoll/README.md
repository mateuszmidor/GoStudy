# twitter votes

# Design
+----------+        +----------------+          +--------------+            +-------------+
|          |        |    Twitter     |  read    | twittervotes |  publish   |     NSQ     |
| Twitter  +-------->   Streaming    +---------->     (go)     +------------> Distributed |
|          |        |      API       |  http    |              |  messages  |  Messaging  |
+----------+        +----------------+          +-------^------+            +------+------+
                                                        |                          |
                                          read keywords |                          | consume messages
                                                        |                          |
                                                +-------+------+            +------v------+
                                                |              |   store    |   counter   |
                                                |   MongoDB    <------------+    (go)     |
                                                |              |   hits     |             |
                                                +--------------+            +-------------+

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
terminal3 (only if using localhost mongodb):
    mongod --dbpath ./db

# Configuration
The following strings are required as env variables (see "setupenv.sh"):
SP_TWITTER_KEY=
SP_TWITTER_SECRET=
SP_TWITTER_ACCESSTOKEN=
SP_TWITTER_ACCESSSECRET=
SP_MONGODB_ADDR=

# Setup MongoDB initial data - strings the twittervotes will be looking for
> mongo
> use ballots
> db.polls.insert({"title":"Test poll", "options":["love", "kiss", "hate", "fight"]})

# Track NSQ
> nsq_tail --topic="votes" --lookupd-http-address=localhost:4161

# Run the app
terminal4
    twittervotes
terminal5
    counter