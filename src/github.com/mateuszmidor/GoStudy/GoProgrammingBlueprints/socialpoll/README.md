# twitter votes
Filter twitter tweets against keywords and show results in form of a web poll.

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
- wittervotes  

terminal5  
- counter

# Remote on AWS instance
0. Create network security key-pair, save the key-pair as ~/.ssh/mm_keypair.pem  
1. Create Amazon Linux instance with docker package, set security key-pair  
2. Setup instance->security groups->inbound:  
    allow SSH 22 
    allow TCP 27017 - default mongodb
    allow TCP 4100-4200 - default nsqd and nsqlookup
3. SSH to the instance:  
    ssh -v -i ~/.ssh/mm_keypair.pem ec2-user@ec2-54-93-96-97.eu-central-1.compute.amazonaws.com  
4. Install docker, add yourself to docker group so you dont need be root:  
    sudo yum install docker  
    sudo usermod -aG docker $(whoami)  
5. Enable docker service autostart and run the service:  
    sudo systemctl enable docker.service  
    sudo systemctl start docker.service  
6. Install docker-compose:  
    sudo yum install -y python-pip  
    sudo pip install docker-compose  
7. Copy docker-compose.yaml and run_docker-compose.sh to AWS instance  
8. Run ./run_docker-compose.sh  